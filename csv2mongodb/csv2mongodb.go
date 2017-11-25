package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"

	c "github.com/carlca/bigdata/company"
	e "github.com/carlca/utils/essentials"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/mgo.v2"
)

const (
	fileName = "../Data/CompanyData.csv"
)

func main() {
	// open CSV file
	csvFile, err := os.Open(fileName)
	e.CheckError(err)
	defer csvFile.Close()
	// record size of CSV file
	fileInfo, err := csvFile.Stat()
	e.CheckError(err)
	fileSize := fileInfo.Size()
	// create counting reader
	cr := &e.CountingReader{Reader: csvFile}
	reader := csv.NewReader(cr)
	recordCount := 0
	// establish MongoDB session
	session, err := mgo.Dial("127.0.0.1")
	e.CheckError(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// skip first row
	_, err = reader.Read()
	// create MongoDB collection
	collection := session.DB("Companies").C("Companies")
	// empty
	collection.RemoveAll(nil)
	// read loop
	for {
		recordCount++
		// read an entire record of CSV values
		record, err := reader.Read()
		// enable printing of thousands characters
		p := message.NewPrinter(language.English)
		if err == io.EOF {
			p.Printf("\nEOF reached")
			p.Printf("\n%d records read", recordCount)
			break
		} else {
			e.CheckError(err)
		}
		// print bytesRead / fileSize
		p.Printf("\r%d / %d", cr.BytesRead, fileSize)
		// insert data into MongoDB
		doc := &c.Company{}
		elem := reflect.ValueOf(doc).Elem()
		for index := 0; index < elem.NumField(); index++ {
			elem.Field(index).SetString(record[index])
		}
		err = collection.Insert(doc)
		e.CheckError(err)
	}
	fmt.Println()
}
