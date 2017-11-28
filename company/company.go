package company

// Company struct - unfinished
type Company struct {
	// Details
	CompanyName   string `bson:"companyname,omitempty" sql:"nvarchar 160"`
	CompanyNumber string `bson:"companynumber,omitempty" sql:"nvarchar 8"`
	Careof        string `bson:"careof,omitempty" sql:"nvarchar 100"`
	POBox         string `bson:"pobox,omitempty" sql:"nvarchar 10"`
	AddressLine1  string `bson:"addressline1,omitempty" sql:"nvarchar 300"`
	AddressLine2  string `bson:"addressline2,omitempty" sql:"nvarchar 300"`
	PostTown      string `bson:"posttown,omitempty" sql:"lookup 50"`
	County        string `bson:"county,omitempty" sql:"lookup 50"`
	Country       string `bson:"country,omitempty" sql:"lookup 50"`
	PostCode      string `bson:"postcode,omitempty" sql:"nvarchar 10"`
	// Corporate
	CompanyCategory   string `bson:"companycategory,omitempty" sql:"lookup 100"`
	CompanyStatus     string `bson:"companystatus,omitempty" sql:"lookup 70"`
	CountryofOrigin   string `bson:"countryoforigin,omitempty" sql:"lookup 50"`
	DissolutionDate   string `bson:"dissolutiondate,omitempty" sql:"date"`
	IncorporationDate string `bson:"incorporationdate,omitempty" sql:"date"`
	// Accounts
	AccountingRefDay       string `bson:"accountingrefday,omitempty" sql:"tinyint"`
	AccountingRefMonth     string `bson:"accountingrefmonth,omitempty" sql:"tinyint"`
	AccountsNextDueDate    string `bson:"accountsnextduedate,omitempty" sql:"date"`
	AccountsNextMadeUpDate string `bson:"accountsnextmadeupdate,omitempty" sql:"date"`
	AccountsCategory       string `bson:"accountscategory,omitempty" sql:"lookup 30"`
	// Returns
	ReturnsNextDueDate    string `bson:"returnsnextduedate,omitempty" sql:"date"`
	ReturnsNextMadeUpDate string `bson:"returnsnextmadeupdate,omitempty" sql:"date"`
	// Mortgages
	NumMortChanges       string `bson:"nummortchanges,omitempty" sql:"int"`
	NumMortOutstanding   string `bson:"nummortoutstanding,omitempty" sql:"int"`
	NumMortPartSatisfied string `bson:"nummortpartsatisfied,omitempty" sql:"int"`
	NumMortSatisfied     string `bson:"nummortsatisfied,omitempty" sql:"int"`
	// SICCodes
	SICCode1 string `bson:"siccode1,omitempty" sql:"lookup 170 00000XXX"`
	SICCode2 string `bson:"siccode2,omitempty" sql:"lookup 170 00000XXX"`
	SICCode3 string `bson:"siccode3,omitempty" sql:"lookup 170 00000XXX"`
	SICCode4 string `bson:"siccode4,omitempty" sql:"lookup 170 00000XXX"`
	// LtdPartners
	NumGenPartners string `bson:"numgenpartners,omitempty" sql:"int"`
	NumLimPartners string `bson:"numlimpartners,omitempty" sql:"int"`
	// Web
	URI string `bson:"uri,omitempty" sql:"nvarchar 47"`
	// Old
	ChangeDate1   string `bson:"changedate1,omitempty" sql:"date"`
	CompanyName1  string `bson:"companyname1,omitempty" sql:"nvarchar 160"`
	ChangeDate2   string `bson:"changedate2,omitempty" sql:"date"`
	CompanyName2  string `bson:"companyname2,omitempty" sql:"nvarchar 160"`
	ChangeDate3   string `bson:"changedate3,omitempty" sql:"date"`
	CompanyName3  string `bson:"companyname3,omitempty" sql:"nvarchar 160"`
	ChangeDate4   string `bson:"changedate4,omitempty" sql:"date"`
	CompanyName4  string `bson:"companyname4,omitempty" sql:"nvarchar 160"`
	ChangeDate5   string `bson:"changedate5,omitempty" sql:"date"`
	CompanyName5  string `bson:"companyname5,omitempty" sql:"nvarchar 160"`
	ChangeDate6   string `bson:"changedate6,omitempty" sql:"date"`
	CompanyName6  string `bson:"companyname6,omitempty" sql:"nvarchar 160"`
	ChangeDate7   string `bson:"changedate7,omitempty" sql:"date"`
	CompanyName7  string `bson:"companyname7,omitempty" sql:"nvarchar 160"`
	ChangeDate8   string `bson:"changedate8,omitempty" sql:"date"`
	CompanyName8  string `bson:"companyname8,omitempty" sql:"nvarchar 160"`
	ChangeDate9   string `bson:"changedate9,omitempty" sql:"date"`
	CompanyName9  string `bson:"companyname9,omitempty" sql:"nvarchar 160"`
	ChangeDate10  string `bson:"changedate10,omitempty" sql:"date"`
	CompanyName10 string `bson:"companyname10,omitempty" sql:"nvarchar 160"`
	// Confirmation
	ConfNextDueData    string `bson:"confnextduedata,omitempty" sql:"date"`
	ConfLastMadeUpDate string `bson:"conflastmadeupcddate,omitempty" sql:"date"`
}
