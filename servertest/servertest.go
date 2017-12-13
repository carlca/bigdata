package main

import (
	s "github.com/carlca/bigdata/server"
	e "github.com/carlca/utils/essentials"
)

// go run servertest.go -debug=true -driver=mssql -user=sa -password=23Skidoo -server=192.168.99.100
//
// go run servertest.go -debug=true -driver=postgres -user=postgres -password=23Skidoo

func main() {
	dbase, debug := s.Connect()
	defer dbase.Close()
	tx, err := dbase.Begin()
	e.CheckError("dbase.Begin()", err, debug)
	defer tx.Commit()
	_, err = tx.Exec("DROP TABLE IF EXISTS Company")
	e.CheckError("DROP TABLE", err, debug)
	_, err = tx.Exec(`CREATE TABLE Company
		(ProductID int PRIMARY KEY NOT NULL,
		ProductName nvarchar(25) NOT NULL,
		Price money NULL,
		Price2 money NULL,
		ProductDescription text NULL)`)
	e.CheckError("CREATE TABLE", err, debug)
}
