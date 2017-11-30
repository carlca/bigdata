package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	c "github.com/carlca/bigdata/company"
	e "github.com/carlca/utils/essentials"
)

// SQLColumn represents one column in an SQLSchema
type SQLColumn struct {
	Name  string
	T     string
	Size  int
	Mask  string
	Table string
}

func (col SQLColumn) String() string {
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

func (col SQLColumn) lookupMaskTable() (string, string) {
	// original text field is replaced by a join to the lookup table
	fieldDDL := fmt.Sprintf("\t%v nvarchar(%v) NOT NULL,\n", col.Name+"_ID", len(col.Mask))
	// lookup table is created with a text index based on the Mask
	var lookupDDL bytes.Buffer
	lookupDDL.WriteString(fmt.Sprintf("CREATE TABLE dbo.%v\n", col.Table))
	lookupDDL.WriteString(fmt.Sprintf("(\n"))
	lookupDDL.WriteString(fmt.Sprintf("\tID nvarchar(%v) NOT NULL PRIMARY KEY,\n", len(col.Mask)))
	lookupDDL.WriteString(fmt.Sprintf("\tDesc nvarchar(%v) NULL\n", col.Size))
	lookupDDL.WriteString(fmt.Sprintf(")\n"))
	return fieldDDL, lookupDDL.String()
}

func (col SQLColumn) lookupMask() (string, string) {
	// original text field is replaced by a join to the lookup table
	fieldDDL := fmt.Sprintf("\t%v nvarchar(%v) NOT NULL,\n", col.Name+"_ID", len(col.Mask))
	// lookup table is created with a text index based on the Mask
	var lookupDDL bytes.Buffer
	lookupDDL.WriteString(fmt.Sprintf("CREATE TABLE dbo.%v\n", col.Table))
	lookupDDL.WriteString(fmt.Sprintf("(\n"))
	lookupDDL.WriteString(fmt.Sprintf("\tID nvarchar(%v) NOT NULL PRIMARY KEY,\n", len(col.Mask)))
	lookupDDL.WriteString(fmt.Sprintf("\tDesc nvarchar(%v) NULL\n", col.Size))
	lookupDDL.WriteString(fmt.Sprintf(")\n"))
	return fieldDDL, lookupDDL.String()
}

func (col SQLColumn) lookup() (string, string) {
	// the original text field is replaced by a join to the lookup table
	fieldDDL := fmt.Sprintf("\t%v %v NULL,\n", col.Name+"_ID", "int")
	// the lookup table is created with a generated int index
	var lookupDDL bytes.Buffer
	lookupDDL.WriteString(fmt.Sprintf("CREATE TABLE dbo.%v\n", col.Name))
	lookupDDL.WriteString(fmt.Sprintf("(\n"))
	lookupDDL.WriteString(fmt.Sprintf("\tID int IDENTITY (1,1),\n"))
	lookupDDL.WriteString(fmt.Sprintf("\tDesc nvarchar(%v) NULL\n", col.Size))
	lookupDDL.WriteString(fmt.Sprintf(")\n"))
	return fieldDDL, lookupDDL.String()
}

// SQLSchema represents the metadata for an SQLServer table
type SQLSchema struct {
	Columns []SQLColumn
}

// AddColumn adds an SQLColumn struct to the SQLSchema
func (s *SQLSchema) AddColumn(name string, t string, size int, mask string, table string) {
	s.Columns = append(s.Columns, SQLColumn{name, t, size, mask, table})
}

// CreateTable return a line in TSQL to create a table
func (s *SQLSchema) CreateTable(tableName string) []string {
	// create lookups slice
	var (
		lookups   []string
		tableDDL  string
		lookupDDL string
	)
	// initialize TSQL text
	tableDDL = fmt.Sprintf("CREATE TABLE dbo.%v\n", tableName)
	tableDDL = tableDDL + fmt.Sprintf("(\n")
	for _, col := range s.Columns {
		switch col.T {
		case "lookup":
			{
				var fieldDDL string
				switch {
				case col.Mask != "" && col.Table != "":
					fieldDDL, lookupDDL = col.lookupMaskTable()
				case col.Mask != "" && col.Table == "":
					fieldDDL, lookupDDL = col.lookupMask()
				case col.Mask == "" && col.Table != "":
					{
					}
				case col.Mask == "" && col.Table == "":
					fieldDDL, lookupDDL = col.lookup()
				}
				tableDDL = tableDDL + fmt.Sprintf("%v", fieldDDL)
				lookups = append(lookups, lookupDDL)
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
	schema := &SQLSchema{}

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
		e.CheckError(err)
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
