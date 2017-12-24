package orm

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	e "github.com/carlca/utils/essentials"
)

// Dbg is a global, ooh!
var Dbg string

// Column represents one column in an Schema
type Column struct {
	schema Schema
	Name   string
	T      string
	Index  int
	Size   int
	Mask   string
	Table  string
}

// Implements the String() func of the Stringer interface
func (col Column) String() string {
	s := fmt.Sprintf(" [name: %v] [type: %v] [index: %v]", col.Name, col.T, col.Index)
	if col.Size > 0 {
		s += fmt.Sprintf(" [size: %v]", col.Size)
	}
	if col.Mask != "" {
		s += fmt.Sprintf(" [mask: %v]", col.Mask)
	}
	if col.Table != "" {
		s += fmt.Sprintf(" [lookup table: %v]", col.Table)
	}
	return s
}

// used internally by Schema.CreateDDL
func (col Column) createDDL() (string, string) {
	var fieldDDL string
	var idDDL string
	// choose between how to handle the lookups
	nprefix := e.Choice(col.schema.IsMSSQL, "n", "")
	dboprefix := e.Choice(col.schema.IsMSSQL, "dbo.", "")
	if col.Mask != "" {
		fieldDDL = fmt.Sprintf("\t%v %vvarchar(%v) NULL,\n", col.Name+"_ID", nprefix, len(col.Mask))
		idDDL = fmt.Sprintf("\tID %vvarchar(%v) NOT NULL PRIMARY KEY,\n", nprefix, len(col.Mask))
	} else {
		fieldDDL = fmt.Sprintf("\t%v %v NULL,\n", col.Name+"_ID", "int")
		idDDL = fmt.Sprintf("\tID int NOT NULL PRIMARY KEY,\n")
		// NOT NULL PRIMARY KEY was IDENTITY (1,1)
	}
	// choose the table name for the lookup
	tableName := e.Choice(col.Table != "", col.Table, col.Name)
	// create the table DDL and the lookup DDL
	var lookupDDL bytes.Buffer
	lookupDDL.WriteString(fmt.Sprintf("CREATE TABLE %v%v\n", dboprefix, tableName))
	lookupDDL.WriteString(fmt.Sprintf("(\n"))
	lookupDDL.WriteString(idDDL)
	lookupDDL.WriteString(fmt.Sprintf("\tDescr %vvarchar(%v) NULL\n", nprefix, col.Size))
	lookupDDL.WriteString(fmt.Sprintf(")\n"))
	return fieldDDL, lookupDDL.String()
}

// Schema represents the metadata for a table
type Schema struct {
	Name    string
	IsMSSQL bool
	Columns []Column
}

// AddColumn adds an Column struct to the SQLSchema
func (s *Schema) AddColumn(name string, t string, size int, mask string, table string) {
	index := len(s.Columns)
	s.Columns = append(s.Columns, Column{*s, name, t, index, size, mask, table})
}

// ImportCSVDef takes any CSV struct and builds Columns from it
func (s *Schema) ImportCSVDef(csvDef interface{}) {
	val := reflect.Indirect(reflect.ValueOf(csvDef))
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
		s.AddColumn(name, t, int(n), m, l)
	}
}

// DropDDLs return a line in TSQL to create a table
func (s *Schema) DropDDLs() []string {
	var drops []string
	dboprefix := e.Choice(s.IsMSSQL, "dbo.", "")
	drops = append(drops, fmt.Sprintf("DROP TABLE IF EXISTS %v%v\n", dboprefix, s.Name))
	for _, col := range s.Columns {
		if col.T == "lookup" {
			// choose the table name for the lookup
			tableName := e.Choice(col.Table != "", col.Table, col.Name)

			drops = append(drops, fmt.Sprintf("DROP TABLE IF EXISTS %v%v\n", dboprefix, tableName))
		}
	}
	return drops
}

// CreateDDLs returns a slice of TSQL statements to create the tables
func (s *Schema) CreateDDLs(tableName string) []string {
	// create lookups slice
	var (
		lookups  []string
		tableDDL string
	)
	// initialize TSQL text
	nprefix := e.Choice(s.IsMSSQL, "n", "")
	dboprefix := e.Choice(s.IsMSSQL, "dbo.", "")
	tableDDL = fmt.Sprintf("CREATE TABLE %v%v\n(\n", dboprefix, tableName)
	for _, col := range s.Columns {
		switch col.T {
		case "lookup":
			{
				fieldDDL, lookupDDL := col.createDDL()
				tableDDL += fmt.Sprintf("%v", fieldDDL)
				if !e.Contains(lookups, lookupDDL) {
					lookups = append(lookups, lookupDDL)
				}
			}
		case "varchar":
			{
				tableDDL += fmt.Sprintf("\t%v %v%v(%v) NULL,\n", col.Name, nprefix, col.T, col.Size)
			}
		default:
			{
				tableDDL += fmt.Sprintf("\t%v %v NULL,\n", col.Name, col.T)
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

// InsertData constructs a DDL statement to insert the values for the main table
func (s *Schema) InsertData(source []string) string {
	// start DDL statement
	ins := fmt.Sprintf("insert into %v (\n", s.Name)
	// run through columns
	for _, col := range s.Columns {
		idsuffux := ""
		if col.T == "lookup" {
			idsuffux = "_ID"
		}
		ins += fmt.Sprintf("%v%v,\n", col.Name, idsuffux)
	}
	// remove final comma
	if strings.HasSuffix(ins, ",\n") {
		ins = ins[:len(ins)-2] + "\n"
	}
	// insert closing )
	ins += fmt.Sprintf(") values (\n")
	// run through columns again
	for _, col := range s.Columns {
		// escape ' to \'
		datum := strings.Replace(source[col.Index], `'`, `''`, -1)
		switch col.T {
		case "varchar":
			vcline := fmt.Sprintf("'%s',\n", datum)
			if len(datum) > col.Size {
				vcline = fmt.Sprintf("'%s',\n", datum[0:col.Size])
				blame := fmt.Sprintf("%s (%d) %s", col.Name, col.Size, datum)
				AddOverflow(source, blame)
			}
			ins += vcline
		case "int":
			ins += fmt.Sprintf("'%s',\n", datum)
		case "smallint":
			ins += fmt.Sprintf("'%s',\n", datum)
		case "date":
			if datum != "" {
				// covert to US date format YYYY-MM-DD
				usDate := ukToUSDate(datum)
				ins += fmt.Sprintf("'%s',\n", usDate)
			} else {
				// blank dates are acceptable
				ins += fmt.Sprintf("'',\n")
			}
		case "lookup":
			ins += fmt.Sprintf("'',\n")
			name := ""
			switch {
			case col.Table == "":
				name = col.Name
			case col.Table != "":
				name = col.Table
			}
			lookupTable := GetLookupTable(name)
			lookupTable.AddRow(0, "Descr")
			dmp := fmt.Sprintf("col.Name: %v  col.Table: %v  col.Mask: %v  Datum: %v\n", col.Name, col.Table, col.Mask, datum)
			Dbg += dmp
		}
	}
	// remove final comma
	if strings.HasSuffix(ins, ",\n") {
		ins = ins[:len(ins)-2] + "\n"
	}
	// insert closing )
	ins += fmt.Sprintf(")\n")
	return ins
}

func ukToUSDate(ukDate string) string {
	parts := strings.Split(ukDate, "/")
	usDate := fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])
	return usDate
}

// DumpColumns returns a string which is a dump of the schema columns
func (s *Schema) DumpColumns() string {
	result := ""
	for _, column := range s.Columns {
		result += fmt.Sprintf("%v\n", column)
	}
	return result
}
