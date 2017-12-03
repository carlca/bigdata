package main

import (
	q "github.com/carlca/bigdata/sqlserver"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db := q.Connect()
	defer db.Close()
	tx := db.NewTx()
	defer tx.CommitTx()
	tx.Exec("DROP TABLE IF EXISTS Company")
	tx.Exec(`CREATE TABLE Company
		(ProductID int PRIMARY KEY NOT NULL,
		ProductName varchar(25) NOT NULL,
		Price money NULL,
		Price2 money NULL,		
		ProductDescription text NULL)`)
	//tx.CommitTx()
}
