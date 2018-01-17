package main

import (
	"fmt"

	o "github.com/carlca/bigdata/orm"
)

func main() {
	tbl1 := o.GetTable("PostTown")
	tbl2 := o.GetTable("County")
	fmt.Println(tbl1)
	fmt.Println(tbl2)
	tbl1 = o.GetTable("PostTown")
	fmt.Println(tbl1)
	fmt.Println(o.Tables)
}
