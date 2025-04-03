package repository

import (
	"context"

	"xtrinio.com/db"
	"xtrinio.com/model"
)

type NseInsiderRepository struct {
	db *db.Db
}

func NewNseInsiderRepository(db *db.Db) NseInsiderRepository {
	return NseInsiderRepository{
		db: db,
	}
}

func (a NseInsiderRepository) GetNseCompanyInsider(ticker string) ([]model.NseInsiderDto, error) {
	return a.db.QueryNseInsider(context.Background(), ticker)
}
