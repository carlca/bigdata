package main

import (
	"testing"

	o "github.com/carlca/bigdata/orm"
	"github.com/franela/goblin"
)

func TestLookups(t *testing.T) {
	var expected string
	var actual string
	g := goblin.Goblin(t)

	tbl1 := o.GetTable("PostTown")
	tbl2 := o.GetTable("County")

	actual = tbl1.Name
	expected = "PostTown"
	g.Describe("GetTable", func() {
		g.It("Should return a Table struct with the specified name", func() {
			g.AssertEqual(expected, actual)
		})
	})

	actual = tbl2.Name
	expected = "County"
	g.Describe("GetTable", func() {
		g.It("Should return a Table struct with the specified name", func() {
			g.AssertEqual(expected, actual)
		})
	})
}
