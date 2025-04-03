package db

import (
	"context"
	"time"

	"xtrinio.com/model"
)

func (db *Db) QueryNseCompanySearchByWord(ctx context.Context, word string) ([]model.NseCompanySearchDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	scanner := db.session.Query(
		"SELECT ticker, company_name FROM nse.company_search Where word = ?",
		word).WithContext(requestCtx).Iter().Scanner()

	var result []model.NseCompanySearchDto = []model.NseCompanySearchDto{}

	for scanner.Next() {
		var ticker, company_name string

		err := scanner.Scan(&ticker, &company_name)
		if err != nil {
			return nil, err
		}

		result = append(result, model.NseCompanySearchDto{
			Ticker:      ticker,
			Word:        word,
			CompanyName: company_name,
		})
	}

	return result, nil
}

func (db *Db) QueryTickerExists(ctx context.Context, ticker string) (bool, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var count int
	if err := db.session.Query("SELECT count(*) FROM nse.company_search Where ticker = ?", ticker).WithContext(requestCtx).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}
