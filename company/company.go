package company

// Company struct
type Company struct {
	// Details
	CompanyName   string `bson:"companyname,omitempty"`   // 160
	CompanyNumber string `bson:"companynumber,omitempty"` // 8
	Careof        string `bson:"careof,omitempty"`        // 100
	POBox         string `bson:"pobox,omitempty"`         // 10
	AddressLine1  string `bson:"addressline1,omitempty"`  // (HouseNumber and Street) 300
	AddressLine2  string `bson:"addressline2,omitempty"`  // (area) 300
	PostTown      string `bson:"posttown,omitempty"`      // 50
	County        string `bson:"county,omitempty"`        // (region) 50
	Country       string `bson:"country,omitempty"`       // 50
	PostCode      string `bson:"postcode,omitempty"`      // 10
	// Corporate
	CompanyCategory   string `bson:"companycategory,omitempty"`   // (corporate_body_type_desc) 100
	CompanyStatus     string `bson:"companystatus,omitempty"`     // (action_code_desc) 70
	CountryofOrigin   string `bson:"countryoforigin,omitempty"`   // 50
	DissolutionDate   string `bson:"dissolutiondate,omitempty"`   // 10
	IncorporationDate string `bson:"incorporationdate,omitempty"` // 10
	// Accounts
	AccountingRefDay       string `bson:"accountingrefday,omitempty"`       // 2
	AccountingRefMonth     string `bson:"accountingrefmonth,omitempty"`     // 2
	AccountsNextDueDate    string `bson:"accountsnextduedate,omitempty"`    // 10
	AccountsNextMadeUpDate string `bson:"accountsnextmadeupdate,omitempty"` // 10
	AccountsCategory       string `bson:"accountscategory,omitempty"`       // (accounts_type_desc) 30
	// Returns
	ReturnsNextDueDate    string `bson:"returnsnextduedate,omitempty"`    // 10
	ReturnsNextMadeUpDate string `bson:"returnsnextmadeupdate,omitempty"` // 10
	// Mortgages
	NumMortChanges       string `bson:"nummortchanges,omitempty"`       // 6
	NumMortOutstanding   string `bson:"nummortoutstanding,omitempty"`   // 6
	NumMortPartSatisfied string `bson:"nummortpartsatisfied,omitempty"` // 6
	NumMortSatisfied     string `bson:"nummortsatisfied,omitempty"`     // 6
	// SICCodes
	SICCode1 string `bson:"siccode1,omitempty"` // 170
	SICCode2 string `bson:"siccode2,omitempty"` // 170
	SICCode3 string `bson:"siccode3,omitempty"` // 170
	SICCode4 string `bson:"siccode4,omitempty"` // 170
	// LtdPartners
	NumGenPartners string `bson:"numgenpartners,omitempty"` // 6
	NumLimPartners string `bson:"numlimpartners,omitempty"` // 6
	// Web
	URI string `bson:"uri,omitempty"` // 47
	// Old
	ChangeDate1   string `bson:"changedate1,omitempty"`   // 10
	CompanyName1  string `bson:"companyname1,omitempty"`  // 160
	ChangeDate2   string `bson:"changedate2,omitempty"`   // 10
	CompanyName2  string `bson:"companyname2,omitempty"`  // 160
	ChangeDate3   string `bson:"changedate3,omitempty"`   // 10
	CompanyName3  string `bson:"companyname3,omitempty"`  // 160
	ChangeDate4   string `bson:"changedate4,omitempty"`   // 10
	CompanyName4  string `bson:"companyname4,omitempty"`  // 160
	ChangeDate5   string `bson:"changedate5,omitempty"`   // 10
	CompanyName5  string `bson:"companyname5,omitempty"`  // 160
	ChangeDate6   string `bson:"changedate6,omitempty"`   // 10
	CompanyName6  string `bson:"companyname6,omitempty"`  // 160
	ChangeDate7   string `bson:"changedate7,omitempty"`   // 10
	CompanyName7  string `bson:"companyname7,omitempty"`  // 160
	ChangeDate8   string `bson:"changedate8,omitempty"`   // 10
	CompanyName8  string `bson:"companyname8,omitempty"`  // 160
	ChangeDate9   string `bson:"changedate9,omitempty"`   // 10
	CompanyName9  string `bson:"companyname9,omitempty"`  // 160
	ChangeDate10  string `bson:"changedate10,omitempty"`  // 10
	CompanyName10 string `bson:"companyname10,omitempty"` // 160
	// Confirmation
	ConfNextDueData    string `bson:"confnextduedata,omitempty"`      // 10
	ConfLastMadeUpDate string `bson:"conflastmadeupcddate,omitempty"` // 10
}
