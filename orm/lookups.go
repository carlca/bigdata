package orm

// Row is one lookup row
type Row struct {
	ID    int
	Descr string
}

// Table is one lookup table
type Table struct {
	Name string
	Rows []*Row
}

// Tables is a map of pointers to Table
var (
	Tables map[string]*Table
)

func init() {
	Tables = make(map[string]*Table)
}

// GetTable either returns an existing Table or creates a new Table
func GetTable(name string) *Table {
	tbl, ok := Tables[name]
	if !ok {
		tbl = new(Table)
		tbl.Name = name
		Tables[name] = tbl
	}
	return tbl
}
