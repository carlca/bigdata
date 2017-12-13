package server

import (
	"database/sql"
	"flag"
	"fmt"

	e "github.com/carlca/utils/essentials"
	// MSSQL driver
	_ "github.com/denisenkom/go-mssqldb"
	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// DB inherits from sql.DB ...
type DB struct {
	*sql.DB
}

// ConnectPostgreSQL establishes contact with an PostgreSQL server
func ConnectPostgreSQL() (*DB, bool) {
	// flags
	debug := flag.Bool("debug", false, "enable debugging")
	user := flag.String("user", "", "the database user")
	password := flag.String("password", "", "the database password")
	// parse command line flags
	flag.Parse()
	// dump flags if debug
	if *debug {
		fmt.Printf("user: %s\n", *user)
		// fmt.Printf("dbname: %s\n", *dbname)
		fmt.Printf("password: %s\n", *password)
	}
	// build connection string
	connString := fmt.Sprintf("user=%s password=%s", *user, *password)
	connString += " sslmode=disable"
	// return DB object
	return CreateDB(connString, "postgres", debug), *debug
}

// ConnectSQLServer establishes contact with an SQL Server
func ConnectSQLServer() (*DB, bool) {
	// flags
	debug := flag.Bool("debug", false, "enable debugging")
	password := flag.String("password", "", "the database password")
	port := flag.Int("port", 1433, "the database port")
	server := flag.String("server", "", "the database server")
	user := flag.String("user", "", "the database user")
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
	// return DB object
	return CreateDB(connString, "mssql", debug), *debug
}

// CreateDB returns a DB object based on connecion string and driver name
func CreateDB(connString, driverName string, debug *bool) *DB {
	// if debug dump connection string
	if *debug {
		fmt.Printf("connString: %s\n", connString)
	}
	// create an SQL Server connection
	dbx, err := sql.Open("mssql", connString)
	e.CheckError("Open DB", err, *debug)
	err = dbx.Ping()
	e.CheckError("db.Ping", err, *debug)
	return &DB{dbx}
}
