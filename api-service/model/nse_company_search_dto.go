package model

type NseCompanySearchDto struct {
	Ticker      string `json:"ticker"`
	Word        string `json:"-"`
	CompanyName string `json:"company_name"`
}
