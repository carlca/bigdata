package sqlserver

import (
	"database/sql"
	"flag"
	"fmt"

	e "github.com/carlca/utils/essentials"
	// MSSQL driver
	_ "github.com/denisenkom/go-mssqldb"
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

// Connect establishes contact with an SQL Server
func Connect() (*DB, bool) {
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
	e.CheckError("Open DB", err, *debug)
	err = dbx.Ping()
	e.CheckError("db.Ping", err, *debug)
	return &DB{dbx}, *debug
}
