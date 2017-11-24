package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"

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

func main() {
	// establish MongoDB session
	session, err := mgo.Dial("127.0.0.1")
	checkError(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// get collection count
	coll := session.DB("Companies").C("Companies")
	count, err := coll.Count()
	checkError(err)
	p := message.NewPrinter(language.English)
	p.Printf("Record count for Companies: %d\n", count)
	// read subset of collection
	var companies []c.Company
	coll.Find(bson.M{}).Limit(20).All(&companies)
	//fmt.Println(companies)
	for _, company := range companies {
		if company.CompanyName != "" {
			fmt.Println(company.CompanyName)
		}
		if company.AddressLine1 != "" {
			fmt.Println(company.AddressLine1)
		}
		if company.AddressLine2 != "" {
			fmt.Println(company.AddressLine2)
		}
		if company.PostCode != "" {
			fmt.Println(company.PostCode)
		}
		fmt.Println("")
	}
}
