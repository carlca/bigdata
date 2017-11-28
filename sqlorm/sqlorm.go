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
		switch len(sqls) {
		case 1:
			t = sqls[0]
			n = 0
			m = ""
		case 2:
			t = sqls[0]
			n, err = strconv.ParseInt(sqls[1], 10, 64)
			m = ""
		case 3:
			t = sqls[0]
			n, err = strconv.ParseInt(sqls[1], 10, 64)
			m = sqls[2]
		}
		e.CheckError(err)
		schema.AddColumn(name, t, int(n), m)
	}
	for _, column := range schema.Columns {
		fmt.Println(column)
	}
	fmt.Println()
}
