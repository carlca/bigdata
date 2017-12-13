package main

import (
	"fmt"

	c "github.com/carlca/bigdata/company"
	o "github.com/carlca/bigdata/orm"
)

func main() {
	schema := &o.Schema{}
	schema.ImportCSVDef(&c.Company{})
	fmt.Printf("%v\n", schema.DumpColumns())
	sqls := schema.CreateDDL("Company")
	for _, sql := range sqls {
		fmt.Println(sql)
	}
}
