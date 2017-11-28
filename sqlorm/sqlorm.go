package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	c "github.com/carlca/bigdata/company"
	e "github.com/carlca/utils/essentials"
)

// SQLColumn represents one column in an SQLSchema
type SQLColumn struct {
	Name string
	T    string
	Size int
	Mask string
}

func (col SQLColumn) String() string {
	s := fmt.Sprintf(" [name: %v] [type: %v]", col.Name, col.T)
	if col.Size > 0 {
		s = s + fmt.Sprintf(" [size: %v]", col.Size)
	}
	if col.Mask != "" {
		s = s + fmt.Sprintf(" [mask: %v]", col.Mask)
	}
	return s
}

// SQLSchema represents the metadata for an SQLServer table
type SQLSchema struct {
	Columns []SQLColumn
}

// AddColumn adds an SQLColumn struct to the SQLSchema
func (s *SQLSchema) AddColumn(name string, t string, size int, mask string) {
	s.Columns = append(s.Columns, SQLColumn{name, t, size, mask})
}

// CreateTable return a line in TSQL to create a table
func (s *SQLSchema) CreateTable(tableName string) []string {
	// create lookups slice
	var lookups []string
	// initialize TSQL text
	l := fmt.Sprintf("CREATE TABLE dbo.%v\n", tableName)
	l = l + fmt.Sprintf("(\n")
	for _, col := range s.Columns {
		switch col.T {
		case "lookup":
			{
				l = l + fmt.Sprintf("	%v %v NULL,\n", col.Name+"_ID", "int")
				ll := fmt.Sprintf("CREATE TABLE dbo.%v\n", col.Name)
				ll = ll + fmt.Sprintf("(\n")
				ll = ll + fmt.Sprintf("\tID int IDENTITY (1,1),\n")
				ll = ll + fmt.Sprintf("\tDesc nvarchar %v\n", col.Size)
				ll = ll + fmt.Sprintf(")\n")
				lookups = append(lookups, ll)
			}
		case "nvarchar":
			{
				l = l + fmt.Sprintf("\t%v %v %v NULL,\n", col.Name, col.T, col.Size)
			}
		default:
			{
				l = l + fmt.Sprintf("\t%v %v NULL,\n", col.Name, col.T)
			}
		}
	}
	// remove final comma
	if strings.HasSuffix(l, ",\n") {
		l = l[:len(l)-2] + "\n"
	}
	// insert closing )
	l = l + fmt.Sprintf(")\n")
	var r []string
	r = append(r, l)
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
		)
		t = sqls[0]
		if len(sqls) > 1 {
			n, err = strconv.ParseInt(sqls[1], 10, 64)
			if len(sqls) > 2 {
				m = sqls[2]
			}
		}
		e.CheckError(err)
		schema.AddColumn(name, t, int(n), m)
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
