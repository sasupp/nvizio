package repository

import (
	"context"

	"xtrinio.com/db"
	"xtrinio.com/model"
)

type SymbolsRepository struct {
	db *db.Db
}

func NewSymbolsRepository(db *db.Db) SymbolsRepository {
	return SymbolsRepository{
		db: db,
	}
}

func (f SymbolsRepository) GetCompanyId(symbol string, exchange string) ([]model.SymbolsDto, error) {
	return f.db.QueryCompanyId(context.Background(), symbol, exchange)
}
