package postgresql

import (
	"database/sql"
	"flag"
	"fmt"

	e "github.com/carlca/utils/essentials"
	// PostgreSQL driver
	_ "github.com/lib/pq"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	user     = flag.String("user", "", "the database user")
	dbname   = flag.String("dbname", "", "the database name")
	password = flag.String("password", "", "the database password")
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
		fmt.Printf("user: %s\n", *user)
		fmt.Printf("dbname: %s\n", *dbname)
		fmt.Printf("password: %s\n", *password)
	}
	// build connection string
	connString := fmt.Sprintf("user=%s dbname=%s password=%s", *user, *dbname, *password)
	connString = connString + " sslmode=verify-full"
	// if debug dump connection string
	if *debug {
		fmt.Printf("connString: %s\n", connString)
	}
	// create an SQL Server connection
	dbx, err := sql.Open("postgres", connString)
	e.CheckError("Open DB", err, *debug)
	err = dbx.Ping()
	e.CheckError("db.Ping", err, *debug)
	return &DB{dbx}, *debug
}
