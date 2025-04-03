package db

import (
	"context"
	"time"

	"xtrinio.com/model"
)

func (db *Db) QueryNseFeedByDate(ctx context.Context, filingDates []string) ([]model.NseUpdateDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := "SELECT filing_date, publish_time, id, ticker, subject, company_name, details, attachment FROM nse.feed where filing_date in ?"
	scanner := db.session.Query(query, filingDates).WithContext(requestCtx).Iter().Scanner()

	var result []model.NseUpdateDto = []model.NseUpdateDto{}

	for scanner.Next() {
		var ticker, filing_date string
		var id, company_name, details, attachment, subject string
		var publish_time time.Time

		err := scanner.Scan(&filing_date, &publish_time, &id, &ticker, &subject, &company_name, &details, &attachment)
		if err != nil {
			return nil, err
		}

		result = append(result, model.NseUpdateDto{
			FilingDate:  filing_date,
			PublishTime: model.NsePublishTime(publish_time),
			Id:          id,
			Ticker:      ticker,
			Subject:     subject,
			CompanyName: company_name,
			Details:     details,
			Attachment:  attachment,
		})
	}

	return result, nil
}

func (db *Db) QueryNseCompanyFeed(ctx context.Context, ticker string) ([]model.NseUpdateDto, error) {
	requestCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	scanner := db.session.Query(
		"SELECT filing_date, publish_time, id, subject, company_name, details, attachment FROM nse.feed Where ticker = ?",
		ticker).WithContext(requestCtx).Iter().Scanner()

	var result []model.NseUpdateDto = []model.NseUpdateDto{}

	for scanner.Next() {
		var publish_time time.Time
		var filing_date string
		var id, company_name, attachment, details, subject string

		err := scanner.Scan(&filing_date, &publish_time, &id, &subject, &company_name, &details, &attachment)
		if err != nil {
			return nil, err
		}

		result = append(result, model.NseUpdateDto{
			FilingDate:  filing_date,
			PublishTime: model.NsePublishTime(publish_time),
			Id:          id,
			Ticker:      ticker,
			Subject:     subject,
			CompanyName: company_name,
			Details:     details,
			Attachment:  attachment,
		})
	}

	return result, nil
}
