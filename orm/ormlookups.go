package orm

// LookupRow is one row withing a lookup table
type LookupRow struct {
	ID    int
	Descr string
}

// LookupTable is an inmemory lookup table
type LookupTable struct {
	Name string
	Rows []LookupRow
}

// LookupTables is a slice of LookupTable structs
var LookupTables []LookupTable
