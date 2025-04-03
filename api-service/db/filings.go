package db

import (
	"context"
	"fmt"
	"time"

	"xtrinio.com/model"
)

func (db *Db) QueryCompanyId(ctx context.Context, symbol string, exchange string) ([]model.SymbolsDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT company_name, company_code, company_id FROM finance.symbols Where symbol = ? AND exchange = ?",
		symbol, exchange).WithContext(requestCtx).Iter().Scanner()

	var result []model.SymbolsDto = []model.SymbolsDto{}

	for scanner.Next() {
		var company_name, company_id string
		var company_code int64

		err := scanner.Scan(&company_name, &company_code, &company_id)
		if err != nil {
			return nil, err
		}

		result = append(result, model.SymbolsDto{
			Symbol:      symbol,
			Exchange:    exchange,
			CompanyName: company_name,
			CompanyCode: company_code,
			CompanyId:   company_id,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("CompanyNotFound %v:%v", exchange, symbol)
	}

	return result, nil

}

func (db *Db) QueryFilings(ctx context.Context, company_id string) ([]model.FilingsDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT filing_id, filing_date, company_code, symbol, company_name, is_amended, filing_link, fiscal_year, fiscal_period, props FROM finance.filings Where company_id = ?",
		company_id).WithContext(requestCtx).Iter().Scanner()

	var result []model.FilingsDto = []model.FilingsDto{}

	for scanner.Next() {
		var filing_date time.Time
		var filing_id, fiscal_period, company_name, filing_link string
		var symbol []string
		var fiscal_year int
		var is_amended bool
		var company_code int64
		var props map[string]string

		err := scanner.Scan(&filing_id, &filing_date, &company_code, &symbol, &company_name, &is_amended, &filing_link, &fiscal_year, &fiscal_period, &props)
		if err != nil {
			return nil, err
		}

		result = append(result, model.FilingsDto{
			FilingId:     filing_id,
			FilingDate:   filing_date,
			CompanyId:    company_id,
			CompanyCode:  company_code,
			FiscalYear:   fiscal_year,
			FiscalPeriod: fiscal_period,
			IsAmended:    is_amended,
			Symbol:       symbol,
			CompanyName:  company_name,
			FilingLink:   filing_link,
			Props:        props,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("FilingsNotFound")
	}

	return result, nil
}

func (db *Db) QueryFiling(ctx context.Context, filing_id string) ([]model.FilingsDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT filing_date, company_id, company_code, symbol, company_name, is_amended, filing_link, fiscal_year, fiscal_period, props FROM finance.filings Where filing_id = ?",
		filing_id).WithContext(requestCtx).Iter().Scanner()

	var result []model.FilingsDto = []model.FilingsDto{}

	for scanner.Next() {
		var filing_date time.Time
		var company_id, fiscal_period, company_name, filing_link string
		var company_code int64
		var symbol []string
		var fiscal_year int
		var is_amended bool
		var props map[string]string

		err := scanner.Scan(&filing_date, &company_id, &company_code, &symbol, &company_name, &is_amended, &filing_link, &fiscal_year, &fiscal_period, &props)
		if err != nil {
			return nil, err
		}

		result = append(result, model.FilingsDto{
			FilingId:     filing_id,
			FilingDate:   filing_date,
			CompanyId:    company_id,
			CompanyCode:  company_code,
			FiscalYear:   fiscal_year,
			FiscalPeriod: fiscal_period,
			IsAmended:    is_amended,
			Symbol:       symbol,
			CompanyName:  company_name,
			FilingLink:   filing_link,
			Props:        props,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("FilingNotFound")
	}

	return result, nil
}

func (db *Db) QueryFilingDocuments(ctx context.Context, filingId string) ([]model.FilingDocumentsDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT doc_id, type, name, tag, title, is_compressed FROM finance.docs Where filing_id = ?",
		filingId).WithContext(requestCtx).Iter().Scanner()

	var result []model.FilingDocumentsDto = []model.FilingDocumentsDto{}

	for scanner.Next() {
		var name, t, tag, title string
		var docId int
		var isCompressed bool

		err := scanner.Scan(&docId, &t, &name, &tag, &title, &isCompressed)
		if err != nil {
			return nil, err
		}

		result = append(result, model.FilingDocumentsDto{
			FilingId:     filingId,
			DocId:        docId,
			Name:         name,
			Type:         t,
			Tag:          tag,
			Title:        title,
			IsCompressed: isCompressed,
			Data:         nil,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("DocumentsNotFound")
	}

	return result, nil
}

func (db *Db) QueryFilingDocumentsByType(ctx context.Context, filingId string, t string) ([]model.FilingDocumentsDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT doc_id, name, tag, title, is_compressed FROM finance.docs Where filing_id = ? and type = ?",
		filingId, t).WithContext(requestCtx).Iter().Scanner()

	var result []model.FilingDocumentsDto = []model.FilingDocumentsDto{}

	for scanner.Next() {
		var name, tag, title string
		var docId int
		var isCompressed bool

		err := scanner.Scan(&docId, &name, &tag, &title, &isCompressed)
		if err != nil {
			return nil, err
		}

		result = append(result, model.FilingDocumentsDto{
			FilingId:     filingId,
			DocId:        docId,
			Name:         name,
			Type:         t,
			Tag:          tag,
			Title:        title,
			IsCompressed: isCompressed,
			Data:         nil,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("DocumentsNotFound")
	}

	return result, nil
}

func (db *Db) QueryFilingDocument(ctx context.Context, filingId string, docId int) ([]model.FilingDocumentsDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT type, name, tag, title, is_compressed, data FROM finance.docs Where filing_id = ? and doc_id = ?",
		filingId, docId).WithContext(requestCtx).Iter().Scanner()

	var result []model.FilingDocumentsDto

	for scanner.Next() {
		var name, t, tag, title string
		var data *[]byte
		var isCompressed bool

		err := scanner.Scan(&t, &name, &tag, &title, &isCompressed, &data)
		if err != nil {
			return nil, err
		}

		result = append(result, model.FilingDocumentsDto{
			FilingId:     filingId,
			DocId:        docId,
			Type:         t,
			Name:         name,
			Tag:          tag,
			Title:        title,
			IsCompressed: isCompressed,
			Data:         data,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("DocumentNotFound")
	}

	return result, nil
}

func (db *Db) QueryApiKey(ctx context.Context, apiKeyHash string) (*model.ApiKey, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT email_id Where apikey_hash = ?", apiKeyHash).WithContext(requestCtx).Iter().Scanner()

	for scanner.Next() {

		var emailId string
		err := scanner.Scan(&emailId)
		if err != nil {
			return nil, err
		}

		return &model.ApiKey{
			ApiKeyHash: apiKeyHash,
			EmailId:    emailId,
		}, nil
	}

	return nil, fmt.Errorf("ApiKeyNotFound")
}
