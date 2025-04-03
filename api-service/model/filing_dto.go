package model

import (
	"strings"
	"time"
)

type FilingsDto struct {
	FilingId     string            `json:"filing_id"`
	FilingDate   time.Time         `json:"-"`
	CompanyId    string            `json:"-"`
	CompanyCode  int64             `json:"company_code"` // Dei: SEC: cik, NSE: scrip code, UK: CompaniesHouseRegisteredNumber
	Symbol       []string          `json:"-"`
	CompanyName  string            `json:"company_name"`
	IsAmended    bool              `json:"-"`
	FilingLink   string            `json:"-"`
	FiscalYear   int               `json:"fiscal_year"`
	FiscalPeriod string            `json:"fiscal_period"`
	Props        map[string]string `json:"props"`
}

type FilingDocumentsDto struct {
	FilingId     string  `json:"filing_id"`
	DocId        int     `json:"doc_id"`
	Type         string  `json:"type"`
	Name         string  `json:"name"`
	Tag          string  `json:"tag"`
	Title        string  `json:"title"`
	IsCompressed bool    `json:"-"`
	Data         *[]byte `json:"-"`
}

func (d *FilingDocumentsDto) IsStatement() bool {
	return strings.ToUpper(d.Type) == "STATEMENT"
}
