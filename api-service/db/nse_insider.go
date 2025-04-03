package db

import (
	"context"
	"time"

	"xtrinio.com/model"
)

func (db *Db) QueryNseInsider(ctx context.Context, ticker string) ([]model.NseInsiderDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT filing_date, publish_time, id, company_name, name, person_category, transaction_type, transaction_value, derivative_contract_type FROM nse.insider Where ticker = ?",
		ticker).WithContext(requestCtx).Iter().Scanner()

	var result []model.NseInsiderDto = []model.NseInsiderDto{}

	for scanner.Next() {
		var publish_time time.Time
		var filing_date time.Time
		var id, company_name, name, person_category, transaction_type, derivative_contract_type string
		var transaction_value float32

		err := scanner.Scan(&filing_date, &publish_time, &id, &company_name, &name, &person_category, &transaction_type, &transaction_value, &derivative_contract_type)
		if err != nil {
			return nil, err
		}

		if derivative_contract_type == "Options" {
			continue
		}

		result = append(result, model.NseInsiderDto{
			FilingDate:       filing_date.Format("02 Jan, 2006"),
			PublishTime:      model.NsePublishTime(publish_time),
			Id:               id,
			Ticker:           ticker,
			CompanyName:      company_name,
			Name:             name,
			PersonCategory:   person_category,
			TransactionType:  transaction_type,
			TransactionValue: transaction_value,
		})
	}

	return result, nil
}
