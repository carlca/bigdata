package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"

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

// Company parent struct
type Company struct {
	Details      *CompanyData      `bson:"details,omitempty"`
	Corporate    *CorporateData    `bson:"corporate,omitempty"`
	Accounts     *AccountsData     `bson:"accounts,omitempty"`
	Returns      *ReturnsData      `bson:"returns,omitempty"`
	Mortgages    *MortgagesData    `bson:"mortgages,omitempty"`
	SICCodes     *SICCodesData     `bson:"sic_codes,omitempty"`
	LtdPartners  *LtdPartnersData  `bson:"ltd_partners,omitempty"`
	Web          *WebData          `bson:"web,omitempty"`
	Old          *OldData          `bson:"old,omitempty"`
	Confirmation *ConfirmationData `bson:"confirmation,omitempty"`
}

// CompanyData child struct
type CompanyData struct {
	CompanyName   string `bson:"companyname,omitempty"`    // 160
	CompanyNumber string `bson:"company_number,omitempty"` // 8
	Careof        string `bson:"careof,omitempty"`         // 100
	POBox         string `bson:"po_box,omitempty"`         // 10
	AddressLine1  string `bson:"address_line_1,omitempty"` // (HouseNumber and Street) 300
	AddressLine2  string `bson:"address_line_2,omitempty"` // (area) 300
	PostTown      string `bson:"post_town,omitempty"`      // 50
	County        string `bson:"county,omitempty"`         // (region) 50
	Country       string `bson:"country,omitempty"`        // 50
	PostCode      string `bson:"post_code,omitempty"`      // 10
}

// CorporateData child struct
type CorporateData struct {
	CompanyCategory   string `bson:"company_category,omitempty"`   // (corporate_body_type_desc) 100
	CompanyStatus     string `bson:"company_status,omitempty"`     // (action_code_desc) 70
	CountryofOrigin   string `bson:"countryof_origin,omitempty"`   // 50
	DissolutionDate   string `bson:"dissolution_date,omitempty"`   // 10
	IncorporationDate string `bson:"incorporation_date,omitempty"` // 10
}

// AccountsData child struct
type AccountsData struct {
	AccountingRefDay   string `bson:"accounting_ref_day,omitempty"`   // 2
	AccountingRefMonth string `bson:"accounting_ref_month,omitempty"` // 2
	NextDueDate        string `bson:"next_due_date,omitempty"`        // 10
	NextMadeUpDate     string `bson:"next_made_up_date,omitempty"`    // 10
	AccountsCategory   string `bson:"accounts_category,omitempty"`    // (accounts_type_desc) 30
}

// ReturnsData child struct
type ReturnsData struct {
	NextDueDate    string `bson:"next_due_date,omitempty"`     // 10
	NextMadeUpDate string `bson:"next_made_up_date,omitempty"` // 10
}

// MortgagesData child struct
type MortgagesData struct {
	NumMortChanges       string `bson:"num_mort_changes,omitempty"`        // 6
	NumMortOutstanding   string `bson:"num_mort_outstanding,omitempty"`    // 6
	NumMortPartSatisfied string `bson:"num_mort_part_satisfied,omitempty"` // 6
	NumMortSatisfied     string `bson:"num_mort_satisfied,omitempty"`      // 6
}

// SICCodesData child structure
type SICCodesData struct {
	SICCode1 string `bson:"sic_code_1,omitempty"` // 170
	SICCode2 string `bson:"sic_code_2,omitempty"` // 170
	SICCode3 string `bson:"sic_code_3,omitempty"` // 170
	SICCode4 string `bson:"sic_code_4,omitempty"` // 170
}

// LtdPartnersData child structure
type LtdPartnersData struct {
	NumGenPartners string `bson:"num_gen_partners,omitempty"` // 6
	NumLimPartners string `bson:"num_lim_partners,omitempty"` // 6
}

// WebData child structure
type WebData struct {
	URI string `bson:"uri,omitempty"` // 47
}

// OldData child structure
type OldData struct {
	ChangeDate1   string `bson:"change_date_1,omitempty"`   // 10
	CompanyName1  string `bson:"company_name_1,omitempty"`  // 160
	ChangeDate2   string `bson:"change_date_2,omitempty"`   // 10
	CompanyName2  string `bson:"company_name_2,omitempty"`  // 160
	ChangeDate3   string `bson:"change_date_3,omitempty"`   // 10
	CompanyName3  string `bson:"company_name_3,omitempty"`  // 160
	ChangeDate4   string `bson:"change_date_4,omitempty"`   // 10
	CompanyName4  string `bson:"company_name_4,omitempty"`  // 160
	ChangeDate5   string `bson:"change_date_5,omitempty"`   // 10
	CompanyName5  string `bson:"company_name_5,omitempty"`  // 160
	ChangeDate6   string `bson:"change_date_6,omitempty"`   // 10
	CompanyName6  string `bson:"company_name_6,omitempty"`  // 160
	ChangeDate7   string `bson:"change_date_7,omitempty"`   // 10
	CompanyName7  string `bson:"company_name_7,omitempty"`  // 160
	ChangeDate8   string `bson:"change_date_8,omitempty"`   // 10
	CompanyName8  string `bson:"company_name_8,omitempty"`  // 160
	ChangeDate9   string `bson:"change_date_9,omitempty"`   // 10
	CompanyName9  string `bson:"company_name_9,omitempty"`  // 160
	ChangeDate10  string `bson:"change_date_10,omitempty"`  // 10
	CompanyName10 string `bson:"company_name_10,omitempty"` // 160
}

// ConfirmationData child structure
type ConfirmationData struct {
	ConfNextDueData    string `bson:"conf_next_due_data,omitempty"`     // 10
	ConfLastMadeUpDate string `bson:"conf_last_made_up_date,omitempty"` // 10
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
	var companies []Company
	coll.Find(bson.M{}).Limit(20).All(&companies)
	//fmt.Println(companies)
	for _, company := range companies {
		if company.Details.CompanyName != "" {
			fmt.Println(company.Details.CompanyName)
		}
		if company.Details.AddressLine1 != "" {
			fmt.Println(company.Details.AddressLine1)
		}
		if company.Details.AddressLine2 != "" {
			fmt.Println(company.Details.AddressLine2)
		}
		if company.Details.PostCode != "" {
			fmt.Println(company.Details.PostCode)
		}
		fmt.Println("")
	}
}
