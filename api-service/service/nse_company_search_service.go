package service

import (
	"xtrinio.com/model"
	"xtrinio.com/repository"
)

type INseCompanySearchService interface {
	GetNseCompanyByWord(word string) ([]model.NseCompanySearchDto, error)
	TickerExists(ticker string) (bool, error)
}

func NewNseCompanySearchService(repo repository.NseCompanySearchRepository) INseCompanySearchService {
	return nseCompanySearchServiceImpl{
		repo: repo,
	}
}

type nseCompanySearchServiceImpl struct {
	repo repository.NseCompanySearchRepository
}

func (a nseCompanySearchServiceImpl) GetNseCompanyByWord(word string) ([]model.NseCompanySearchDto, error) {
	return a.repo.GetNseCompanyByWord(word)
}

func (a nseCompanySearchServiceImpl) TickerExists(ticker string) (bool, error) {
	return a.repo.TickerExists(ticker)
}
