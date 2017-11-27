package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	c "github.com/carlca/bigdata/company"
	e "github.com/carlca/utils/essentials"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/mgo.v2"
)

func main() {
	// establish MongoDB session
	session, err := mgo.Dial("127.0.0.1")
	e.CheckError(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// get collection count
	coll := session.DB("Companies").C("Companies")
	count, err := coll.Count()
	e.CheckError(err)
	p := message.NewPrinter(language.English)
	p.Printf("Record count for Companies: %d\n", count)
	// read subset of collection
	var companies []c.Company
	coll.Find(bson.M{}).Limit(200).All(&companies)
	//fmt.Println(companies)
	for _, company := range companies {
		if company.CompanyName != "" {
			fmt.Println(company.CompanyName)
		}
		// if company.AddressLine1 != "" {
		// 	fmt.Println(company.AddressLine1)
		// }
		// if company.AddressLine2 != "" {
		// 	fmt.Println(company.AddressLine2)
		// }
		// if company.PostCode != "" {
		// 	fmt.Println(company.PostCode)
		// }
		// if company.CompanyCategory != "" {
		// 	fmt.Println(company.CompanyCategory)
		// }
		// if company.CompanyStatus != "" {
		// 	fmt.Println(company.CompanyStatus)
		// }
		// if company.CountryofOrigin != "" {
		// 	fmt.Println(company.CountryofOrigin)
		// }
		if company.SICCode1 != "" {
			fmt.Println(company.SICCode1)
		}
		if company.SICCode2 != "" {
			fmt.Println(company.SICCode2)
		}
		if company.SICCode3 != "" {
			fmt.Println(company.SICCode3)
		}
		if company.SICCode4 != "" {
			fmt.Println(company.SICCode4)
		}
		fmt.Println("")
	}
}
