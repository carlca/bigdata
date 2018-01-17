package orm

import (
	"fmt"
)

// LookupRow is one row within a lookup table
type LookupRow struct {
	ID    int
	Descr string
}

// String implements Stringer interface for LookupRow
func (r LookupRow) String() string {
	return fmt.Sprintf("%d: %s", r.ID, r.Descr)
}

// ByID implements sort.Interface for []LookupRow based on the ID field
type ByID []LookupRow

func (r ByID) Len() int           { return len(r) }
func (r ByID) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByID) Less(i, j int) bool { return r[i].ID < r[j].ID }

// LookupTable is an in memory lookup table
type LookupTable struct {
	Name string
	Rows []LookupRow
}

// String implements Stringer interface for LookupTable
func (l LookupTable) String() string {
	result := fmt.Sprintf("%s:\n\n", l.Name)
	for _, row := range l.Rows {
		result += fmt.Sprintf("%s\n", row)
	}
	return result
}

// FindDescr returns the id of the Row which contains descr
func (l LookupTable) FindDescr(descr string) int {
	for _, row := range l.Rows {
		if row.Descr == descr {
			return row.ID
		}
	}
	return 0
}

// NextID returns the next unused vakue for ID
func (l LookupTable) NextID() int {
	maxID := 0
	for _, row := range l.Rows {
		if row.ID > maxID {
			maxID = row.ID
		}
	}
	return maxID + 1
}

// AddRow adds a LookupRow struct to a LookupTable
func (l LookupTable) AddRow(id int, descr string) {
	row := &LookupRow{id, descr}
	l.Rows = append(l.Rows, *row)
}

// LookupTables is a struct containing a slice of LookupTable structs
type LookupTables []LookupTable

// CreateLookupTables is a factory method for LookupTables
func CreateLookupTables() {
	Lookups = make(LookupTables, 0)
}

//  Lookups is a global reference with the orm package
var (
	Lookups LookupTables
)

// GetTable returns a new or existing instance of *LookupTable
func (ls LookupTables) GetTable(name string) LookupTable {
	var result LookupTable
	result = ls.findTable(name)
	if result.Name == "" {
		result = ls.AddTable(name)
	}
	return result
}

// AddTable adds a new LookupTable to ls
func (ls LookupTables) AddTable(name string) LookupTable {
	newTable := LookupTable{name, nil}
	ls = append(ls, newTable)
	return newTable
}

// findTable is used internally by GetTable to find a lookupTable by Name
func (ls LookupTables) findTable(name string) LookupTable {
	var result LookupTable
	for _, table := range ls {
		if table.Name == name {
			return table
		}
	}
	return result
}

// String implements Stringer interface for LookupTables
func (ls LookupTables) String() string {
	result := ""
	for _, table := range ls {
		result += fmt.Sprintf("%s\n", table)
	}
	return result
}
