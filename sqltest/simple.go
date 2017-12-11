package main

import (
	q "github.com/carlca/bigdata/sqlserver"
	e "github.com/carlca/utils/essentials"
	//_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db := q.Connect()
	defer db.Close()
	tx, err := db.Begin()
	e.CheckError("db.Begin() failed", err)
	defer tx.Commit()
	tx.Exec("DROP TABLE IF EXISTS Company")
	//time.Sleep(200 * time.Millisecond)
	tx.Exec(`CREATE TABLE Company
		(ProductID int PRIMARY KEY NOT NULL,
		ProductName varchar(25) NOT NULL,
		Price money NULL,
		Price2 money NULL,		
		ProductDescription text NULL)`)
	//time.Sleep(200 * time.Millisecond)
	//tx.CommitTx()
}
