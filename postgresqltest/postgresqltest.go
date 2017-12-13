package main

import (
	s "github.com/carlca/bigdata/server"
	e "github.com/carlca/utils/essentials"
)

func main() {
	dbase, debug := s.ConnectPostgreSQL()
	defer dbase.Close()
	tx, err := dbase.Begin()
	e.CheckError("dbase.Begin()", err, debug)
	defer tx.Commit()
	_, err = tx.Exec("DROP TABLE IF EXISTS Company")
	e.CheckError("DROP TABLE", err, debug)
	_, err = tx.Exec(`CREATE TABLE Company
		(ProductID int PRIMARY KEY NOT NULL,
		ProductName varchar(25) NOT NULL,
		Price money NULL,
		Price2 money NULL,
		ProductDescription text NULL)`)
	e.CheckError("CREATE TABLE", err, debug)
}
