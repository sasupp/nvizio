package db

import (
	"context"
	"time"

	"xtrinio.com/model"
)

func (db *Db) QueryNseShareholding(ctx context.Context, ticker string) ([]model.NseShareholdingDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT filing_date, publish_time, id, company_name, as_on_date, promoter_group, public, employee_trusts FROM nse.shareholding Where ticker = ?",
		ticker).WithContext(requestCtx).Iter().Scanner()

	var result []model.NseShareholdingDto = []model.NseShareholdingDto{}

	for scanner.Next() {
		var publish_time time.Time
		var filing_date string
		var as_on_date time.Time
		var id, company_name string
		var promoter_group, public, employee_trusts float32

		err := scanner.Scan(&filing_date, &publish_time, &id, &company_name, &as_on_date, &promoter_group, &public, &employee_trusts)
		if err != nil {
			return nil, err
		}

		result = append(result, model.NseShareholdingDto{
			FilingDate:     filing_date,
			PublishTime:    model.NsePublishTime(publish_time),
			Id:             id,
			Ticker:         ticker,
			CompanyName:    company_name,
			AsOnDate:       as_on_date.Format("02 Jan, 2006"),
			PromoterGroup:  promoter_group,
			Public:         public,
			EmployeeTrusts: employee_trusts,
		})
	}

	return result, nil
}
