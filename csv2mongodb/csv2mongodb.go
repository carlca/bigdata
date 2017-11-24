package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	c "github.com/carlca/bigdata/company"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/mgo.v2"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
		os.Exit(1) // or anything else ...
	}
}

// contributed by Jakob Borg
type countingReader struct {
	reader    io.Reader
	bytesRead int64 // bytes
}

// contributed by Jakob Borg
func (c *countingReader) Read(bs []byte) (int, error) {
	n, err := c.reader.Read(bs)
	c.bytesRead += int64(n)
	return n, err
}

const (
	fileName = "../Data/CompanyData.csv"
)

func main() {
	// open CSV file
	csvFile, err := os.Open(fileName)
	checkError(err)
	defer csvFile.Close()
	// record size of CSV file
	fileInfo, err := csvFile.Stat()
	checkError(err)
	fileSize := fileInfo.Size()
	// create counting reader
	cr := &countingReader{reader: csvFile}
	reader := csv.NewReader(cr)
	recordCount := 0
	// establish MongoDB session
	session, err := mgo.Dial("127.0.0.1")
	checkError(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
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
			checkError(err)
		}
		// print bytesRead / fileSize
		p.Printf("\r%d / %d", cr.bytesRead, fileSize)
		// create MongoDB collection
		collection := session.DB("Companies").C("Companies")
		// insert data into MongoDB
		err = collection.Insert(&c.Company{
			// Details
			CompanyName:            record[0],
			CompanyNumber:          record[1],
			Careof:                 record[2],
			POBox:                  record[3],
			AddressLine1:           record[4],
			AddressLine2:           record[5],
			PostTown:               record[6],
			County:                 record[7],
			Country:                record[8],
			PostCode:               record[9],
			CompanyCategory:        record[10],
			CompanyStatus:          record[11],
			CountryofOrigin:        record[12],
			DissolutionDate:        record[13],
			IncorporationDate:      record[14],
			AccountingRefDay:       record[15],
			AccountingRefMonth:     record[16],
			AccountsNextDueDate:    record[17],
			AccountsNextMadeUpDate: record[18],
			AccountsCategory:       record[19],
			ReturnsNextDueDate:     record[20],
			ReturnsNextMadeUpDate:  record[21],
			NumMortChanges:         record[22],
			NumMortOutstanding:     record[23],
			NumMortPartSatisfied:   record[24],
			NumMortSatisfied:       record[25],
			SICCode1:               record[26],
			SICCode2:               record[27],
			SICCode3:               record[28],
			SICCode4:               record[29],
			NumGenPartners:         record[30],
			NumLimPartners:         record[31],
			URI:                    record[32],
			ChangeDate1:            record[33],
			CompanyName1:           record[34],
			ChangeDate2:            record[35],
			CompanyName2:           record[36],
			ChangeDate3:            record[37],
			CompanyName3:           record[38],
			ChangeDate4:            record[39],
			CompanyName4:           record[40],
			ChangeDate5:            record[41],
			CompanyName5:           record[42],
			ChangeDate6:            record[43],
			CompanyName6:           record[44],
			ChangeDate7:            record[45],
			CompanyName7:           record[46],
			ChangeDate8:            record[47],
			CompanyName8:           record[48],
			ChangeDate9:            record[49],
			CompanyName9:           record[50],
			ChangeDate10:           record[51],
			CompanyName10:          record[52],
			ConfNextDueData:        record[53],
			ConfLastMadeUpDate:     record[54],
		})
		checkError(err)
		// if recordCount == 100 {
		// 	break
	}
	fmt.Println()
}
