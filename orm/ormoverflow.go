package orm

// Overflow is a struct to capture any VarChar errors in size
type Overflow struct {
	Source []string
	Blame  string
}

// Overflows is a slice of Overflow structs
var Overflows []Overflow
