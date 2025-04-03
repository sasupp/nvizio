package repository

import (
	"context"

	"xtrinio.com/db"
)

type ApikeyRepository struct {
	db *db.Db
}

func NewApikeyRepository(db *db.Db) ApikeyRepository {
	return ApikeyRepository{
		db: db,
	}
}

func (r ApikeyRepository) IsApikeyHashAvailable(apiKeyHash string) (bool, error) {
	apikey, err := r.db.QueryApiKey(context.Background(), apiKeyHash)
	if err != nil {
		return false, err
	}

	return apikey != nil, nil
}
