package repository

import (
	"context"

	"xtrinio.com/db"
	"xtrinio.com/model"
)

type NseShareholdingRepository struct {
	db *db.Db
}

func NewNseShareholdingRepository(db *db.Db) NseShareholdingRepository {
	return NseShareholdingRepository{
		db: db,
	}
}

func (a NseShareholdingRepository) GetNseCompanyShareholding(ticker string) ([]model.NseShareholdingDto, error) {
	return a.db.QueryNseShareholding(context.Background(), ticker)
}
