package repository

import (
	"context"

	"xtrinio.com/db"
	"xtrinio.com/model"
)

type NseCompanySearchRepository struct {
	db *db.Db
}

func NewNseCompanySearchRepository(db *db.Db) NseCompanySearchRepository {
	return NseCompanySearchRepository{
		db: db,
	}
}

func (a NseCompanySearchRepository) GetNseCompanyByWord(word string) ([]model.NseCompanySearchDto, error) {
	return a.db.QueryNseCompanySearchByWord(context.Background(), word)
}

func (a NseCompanySearchRepository) TickerExists(ticker string) (bool, error) {
	return a.db.QueryTickerExists(context.Background(), ticker)
}
