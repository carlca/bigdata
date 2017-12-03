package sqlserver

import (
	"database/sql"
	"flag"
	"fmt"

	e "github.com/carlca/utils/essentials"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "", "the database password")
	port     = flag.Int("port", 1433, "the database port")
	server   = flag.String("server", "", "the database server")
	user     = flag.String("user", "", "the database user")
)

// DB inherits from sql.DB
type DB struct {
	*sql.DB
}

// Exec combines Prepare and Exec methods
// func (db *DB) Exec(cmd string) {
// 	stmt, err := db.Prepare(cmd)
// 	e.CheckError("prepare: "+cmd+" failed", err)
// 	_, err = stmt.Exec()
// 	e.CheckError("exec: "+cmd+" failed", err)
// 	if *debug {
// 		fmt.Printf("exec: %v succeeded\n", cmd)
// 	}
// }

// NewTx wraps the *DB.Begin func
func (db *DB) NewTx() *Tx {
	tnx, err := db.Begin()
	e.CheckError("BeginTx failed: ", err)
	if *debug {
		fmt.Printf("NewTx: succeeded\n")
	}
	return &Tx{*tnx}
}

// Tx inherits from sql.Tx
type Tx struct {
	sql.Tx
}

// Exec wraps *Tx.Exec
func (tx *Tx) Exec(cmd string) {
	stmt, err := tx.Prepare(cmd)
	e.CheckError("prepare: "+cmd+" failed", err)
	_, err = stmt.Exec()
	e.CheckError("exec: "+cmd+" failed", err)
	if *debug {
		fmt.Printf("exec: %v succeeded\n", cmd)
	}
}

// CommitTx wraps the *Tx.Commit func
func (tx *Tx) CommitTx() {
	err := tx.Commit()
	e.CheckError("tx.CommitTx failed: ", err)
	if *debug {
		fmt.Printf("tx.CommitTx succeeded\n")
	}
}

// Connect establishes contact with an SQL Server
func Connect() *DB {
	// parse command line flags
	flag.Parse()
	// dump flags if debug
	if *debug {
		fmt.Printf("password: %s\n", *password)
		fmt.Printf("port: %d\n", *port)
		fmt.Printf("server: %s\n", *server)
		fmt.Printf("user: %s\n", *user)
	}
	// build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	// if debug dump connection string
	if *debug {
		fmt.Printf("connString: %s\n", connString)
	}
	// create an SQL Server connection
	dbx, err := sql.Open("mssql", connString)
	e.CheckError("Open DB failed: ", err)
	if *debug {
		fmt.Printf("open mssql: succeeded\n")
	}
	err = dbx.Ping()
	e.CheckError("db.Ping failed", err)
	return &DB{dbx}
}
