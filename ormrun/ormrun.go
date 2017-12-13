package main

import (
	c "github.com/carlca/bigdata/company"
	o "github.com/carlca/bigdata/orm"
)

func main() {
	schema := &o.Schema{}
	schema.ImportCSVDef(&c.Company{})
}
