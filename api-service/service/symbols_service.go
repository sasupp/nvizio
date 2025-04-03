package service

import (
	"fmt"

	"xtrinio.com/repository"
)

type ISymbolsService interface {
	GetCompanyId(symbol string, exchange string) (string, error)
}

type symbolsServiceImpl struct {
	repo repository.SymbolsRepository
}

func (a symbolsServiceImpl) GetCompanyId(symbol string, exchange string) (string, error) {
	symbolsDto, err := a.repo.GetCompanyId(symbol, exchange)
	if err != nil {
		return "", err
	}
	if len(symbolsDto) == 0 {
		return "", fmt.Errorf("CompanyNotFound %v:%v", exchange, symbol)
	}
	return symbolsDto[0].CompanyId, nil
}

func NewSymbolsService(repo repository.SymbolsRepository) ISymbolsService {
	return symbolsServiceImpl{
		repo: repo,
	}
}
