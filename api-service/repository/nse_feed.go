package repository

import (
	"context"

	"xtrinio.com/db"
	"xtrinio.com/model"
)

type NseFeedRepository struct {
	db *db.Db
}

func NewNseFeedRepository(db *db.Db) NseFeedRepository {
	return NseFeedRepository{
		db: db,
	}
}

func (a NseFeedRepository) GetNseFeedByDate(filingDates []string) ([]model.NseUpdateDto, error) {
	return a.db.QueryNseFeedByDate(context.Background(), filingDates)
}

func (a NseFeedRepository) GetNseCompanyFeed(ticker string) ([]model.NseUpdateDto, error) {
	return a.db.QueryNseCompanyFeed(context.Background(), ticker)
}
