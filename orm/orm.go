package orm

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	c "github.com/carlca/bigdata/company"
	e "github.com/carlca/utils/essentials"
)

// Column represents one column in an Schema
type Column struct {
	Name  string
	T     string
	Size  int
	Mask  string
	Table string
}

func (col Column) String() string {
	s := fmt.Sprintf(" [name: %v] [type: %v]", col.Name, col.T)
	if col.Size > 0 {
		s = s + fmt.Sprintf(" [size: %v]", col.Size)
	}
	if col.Mask != "" {
		s = s + fmt.Sprintf(" [mask: %v]", col.Mask)
	}
	if col.Table != "" {
		s = s + fmt.Sprintf(" [lookup table: %v]", col.Table)
	}
	return s
}

func (col Column) createDDL() (string, string) {
	var fieldDDL string
	var idDDL string
	// choose between how to handle the lookups
	if col.Mask != "" {
		fieldDDL = fmt.Sprintf("\t%v nvarchar(%v) NULL,\n", col.Name+"_ID", len(col.Mask))
		idDDL = fmt.Sprintf("\tID nvarchar(%v) NOT NULL PRIMARY KEY,\n", len(col.Mask))
	} else {
		fieldDDL = fmt.Sprintf("\t%v %v NULL,\n", col.Name+"_ID", "int")
		idDDL = fmt.Sprintf("\tID int IDENTITY (1,1),\n")
	}
	// choose the table name for the lookup
	tableName := e.Choice(col.Table != "", col.Table, col.Name)
	// create the table DDL and the lookup DDL
	var lookupDDL bytes.Buffer
	lookupDDL.WriteString(fmt.Sprintf("CREATE TABLE dbo.%v\n", tableName))
	lookupDDL.WriteString(fmt.Sprintf("(\n"))
	lookupDDL.WriteString(idDDL)
	lookupDDL.WriteString(fmt.Sprintf("\tDesc nvarchar(%v) NULL\n", col.Size))
	lookupDDL.WriteString(fmt.Sprintf(")\n"))
	return fieldDDL, lookupDDL.String()
}

// Schema represents the metadata for an SQLServer table
type Schema struct {
	Columns []Column
}

// AddColumn adds an Column struct to the SQLSchema
func (s *Schema) AddColumn(name string, t string, size int, mask string, table string) {
	s.Columns = append(s.Columns, Column{name, t, size, mask, table})
}

// CreateTable return a line in TSQL to create a table
func (s *Schema) CreateTable(tableName string) []string {
	// create lookups slice
	var (
		lookups  []string
		tableDDL string
	)
	// initialize TSQL text
	tableDDL = fmt.Sprintf("CREATE TABLE dbo.%v\n", tableName)
	tableDDL = tableDDL + fmt.Sprintf("(\n")
	for _, col := range s.Columns {
		switch col.T {
		case "lookup":
			{
				fieldDDL, lookupDDL := col.createDDL()
				tableDDL = tableDDL + fmt.Sprintf("%v", fieldDDL)
				if !e.Contains(lookups, lookupDDL) {
					lookups = append(lookups, lookupDDL)
				}
			}
		case "nvarchar":
			{
				tableDDL = tableDDL + fmt.Sprintf("\t%v %v(%v) NULL,\n", col.Name, col.T, col.Size)
			}
		default:
			{
				tableDDL = tableDDL + fmt.Sprintf("\t%v %v NULL,\n", col.Name, col.T)
			}
		}
	}
	// remove final comma
	if strings.HasSuffix(tableDDL, ",\n") {
		tableDDL = tableDDL[:len(tableDDL)-2] + "\n"
	}
	// insert closing )
	tableDDL = tableDDL + fmt.Sprintf(")\n")

	var r []string
	r = append(r, tableDDL)
	r = append(r, lookups...)
	return r
}

func main() {
	doc := &c.Company{}
	schema := &Schema{}

	val := reflect.Indirect(reflect.ValueOf(doc))
	for index := 0; index < val.NumField(); index++ {
		name := val.Type().Field(index).Name
		sql := val.Type().Field(index).Tag.Get("sql")
		sqls := strings.Split(sql, " ")
		var (
			err error
			t   string
			n   int64
			m   string
			l   string
		)
		t = sqls[0]
		if len(sqls) > 1 {
			n, err = strconv.ParseInt(sqls[1], 10, 64)
			if len(sqls) > 2 {
				m = sqls[2]
				if len(sqls) > 2 {
					l = sqls[3]
				}
			}
		}
		e.CheckError("", err, false)
		schema.AddColumn(name, t, int(n), m, l)
	}
	for _, column := range schema.Columns {
		fmt.Println(column)
	}
	fmt.Println()
	sqls := schema.CreateTable("Company")
	for _, sql := range sqls {
		fmt.Println(sql)
	}
}
