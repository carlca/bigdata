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

// AddRow adds a LookupRow struct to a LookupTable
func (l *LookupTable) AddRow(id int, descr string) {
	row := &LookupRow{id, descr}
	l.Rows = append(l.Rows, *row)
}

// LookupTables is a slice of LookupTable structs
var LookupTables []*LookupTable

func findTable(name string) *LookupTable {
	for _, lookupTable := range LookupTables {
		if lookupTable.Name == name {
			return lookupTable
		}
	}
	return nil
}

// GetLookupTable returns a new or existing instance iof *LookupTable
func GetLookupTable(name string) *LookupTable {
	var result *LookupTable
	result = findTable(name)
	if result == nil {
		result = &LookupTable{name, nil}
		LookupTables = append(LookupTables, result)
	}
	return result
}
