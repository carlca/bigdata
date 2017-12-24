package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	c "github.com/carlca/bigdata/company"
	o "github.com/carlca/bigdata/orm"
	s "github.com/carlca/bigdata/server"
	e "github.com/carlca/utils/essentials"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// go run ormrun.go -debug=true -driver=mssql -user=sa -password=23Skidoo -server=192.168.99.100

// go run ormrun.go -debug=true -driver=postgres -user=postgres -password=23Skidoo

const (
	fileName = "../Data/CompanyData.csv"
)

func main() {
	// get implementation independant DB and Tx references
	dbase, debug := s.Connect()
	defer dbase.Close()
	tx, err := dbase.Begin()
	e.CheckError("dbase.Begin()", err, debug)
	defer tx.Commit()
	// create Company schema
	schema := &o.Schema{Name: "Company", IsMSSQL: dbase.IsMSSQL}
	// read Company struct data
	schema.ImportCSVDef(&c.Company{})
	if debug {
		fmt.Printf("%v\n", schema.DumpColumns())
	}
	// drop any existing tables
	for _, drop := range schema.DropDDLs() {
		_, err = tx.Exec(drop)
		e.CheckError(drop, err, true)
	}
	// create new tables
	for _, sql := range schema.CreateDDLs("Company") {
		_, err = tx.Exec(sql)
		e.CheckError(sql, err, true)
	}
	// open CSV file
	csvFile, err := os.Open(fileName)
	e.CheckError("os.Open()", err, false)
	defer csvFile.Close()
	// record size of CSV file
	fileInfo, err := csvFile.Stat()
	e.CheckError("csvFile.Stat()", err, false)
	fileSize := fileInfo.Size()
	// create counting reader
	cr := &e.CountingReader{Reader: csvFile}
	reader := csv.NewReader(cr)
	recordCount := 0
	// skip first row
	_, err = reader.Read()
	// read loop
	var record []string
	for {
		recordCount++
		// read an entire record of CSV values
		record, err = reader.Read()
		// enable printing of thousands characters
		p := message.NewPrinter(language.English)
		if err == io.EOF {
			p.Printf("\nEOF reached")
			p.Printf("\n%d records read", recordCount)
			break
		} else {
			e.CheckError("reader.Read()", err, false)
		}
		// print bytesRead / fileSize
		p.Printf("\r%d / %d", cr.BytesRead, fileSize)
		// insert data into database
		ins := schema.InsertData(record)
		_, err = tx.Exec(ins)
		if err != nil {
			bytes := []byte(ins)
			ioutil.WriteFile("/users/carlca/debug.txt", bytes, 0755)
		}
		e.CheckError(ins, err, false)
		if recordCount > 3 {
			break
		}
	}
	if len(o.Overflows) > 0 {
		for _, overflow := range o.Overflows {
			fmt.Println(overflow)
		}
	}
	bytes := []byte(o.Dbg)
	ioutil.WriteFile("/users/carlca/debug3.txt", bytes, 0755)
}
