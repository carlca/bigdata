package main

import (
	"fmt"
	"reflect"

	c "github.com/carlca/bigdata/company"
)

// SQLColumn represents one column in an SQLSchema
type SQLColumn struct {
	Name string
	T    string
	Size int
}

// SQLSchema represents the metadata for an SQLServer table
type SQLSchema struct {
	Columns []SQLColumn
}

func (s *SQLSchema) addColumn(name string, t string, size int) {
	s.Columns = append(s.Columns, SQLColumn{name, t, size})
}

func main() {
	doc := &c.Company{}
	schema := &SQLSchema{}

	val := reflect.Indirect(reflect.ValueOf(doc))
	for index := 0; index < val.NumField(); index++ {
		name := val.Type().Field(index).Name
		sqlInfo := val.Type().Field(index).Tag.Get("sql")
		schema.addColumn(name, sqlInfo, 0)
	}
	for _, column := range schema.Columns {
		fmt.Println(column)
	}
	fmt.Println()
}
