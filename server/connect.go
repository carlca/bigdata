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

func buildConnString(driver, user, password, server string) string {
	var connString string
	switch driver {
	case "mssql":
		{
			port := 1433
			connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)
		}
	case "postgres":
		{
			sslmode := "disable"
			connString = fmt.Sprintf("user=%s password=%s sslmode=%s", user, password, sslmode)
		}
	}
	return connString
}

// Connect establishes contact with an SQL Server
func Connect() (*DB, bool) {
	// flags
	driver := flag.String("driver", "", "db driver name")
	debug := flag.Bool("debug", false, "enable debugging")
	user := flag.String("user", "", "the database user")
	password := flag.String("password", "", "the database password")
	server := flag.String("server", "", "the database server")
	// parse command line flags
	flag.Parse()
	// connection string
	connString := buildConnString(*driver, *user, *password, *server)
	if *debug {
		fmt.Printf("connString: %s\n", connString)
	}
	// create an SQL Server connection
	dbx, err := sql.Open(*driver, connString)
	e.CheckError("Open DB", err, *debug)
	err = dbx.Ping()
	e.CheckError("db.Ping", err, *debug)
	return &DB{dbx}, *debug
}
