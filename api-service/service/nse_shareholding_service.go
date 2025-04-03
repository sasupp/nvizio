package service

import (
	"sort"

	"xtrinio.com/lib/table"
	"xtrinio.com/model"
	"xtrinio.com/repository"
)

type INseShareholdingService interface {
	GetNseShareholdingByCompany(ticker string) ([]model.NseShareholdingDto, error)
	ToHtmlTable(entries []model.NseShareholdingDto) *table.HtmlTable
}

type nseShareholdingServiceImpl struct {
	repo repository.NseShareholdingRepository
}

func (a nseShareholdingServiceImpl) GetNseShareholdingByCompany(ticker string) ([]model.NseShareholdingDto, error) {
	companyShareholdings, err := a.repo.GetNseCompanyShareholding(ticker)
	if err != nil {
		return nil, err
	}
	sort.Sort(model.SortNseShareholdingByPublishTime(companyShareholdings))
	return companyShareholdings, nil
}

func NewNseShareholdingService(repo repository.NseShareholdingRepository) INseShareholdingService {
	return nseShareholdingServiceImpl{
		repo: repo,
	}
}

func (a nseShareholdingServiceImpl) ToHtmlTable(entries []model.NseShareholdingDto) *table.HtmlTable {
	builder := table.NewTableBuilder()
	builder.Add(table.ShpHeader())
	builder.Add(table.ShpRows(entries))
	return builder.Build()
}
