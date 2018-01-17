package main

import (
	"fmt"
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
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = tbl2.Name
	expected = "County"
	g.Describe("GetTable", func() {
		g.It("Should return a Table struct with the specified name", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

}

func ReportAndAssert(g *goblin.G, expected, actual string) {
	if expected == actual {
		PrintCheck()
	} else {
		PrintFail()
	}
	fmt.Print(MakeGray("Expected: " + expected))
	fmt.Println(MakeGray("   Actual: " + actual))
	g.Assert(actual).Equal(expected)
}

func PrintCheck() {
	fmt.Print("    \033[32m\u2713\033[0m ")
}

func PrintFail() {
	fmt.Print("    \033[31m" + "x" + "\033[0m ")
}

func MakeGray(s string) string {
	return "\033[90m" + s + "\033[0m"
}
