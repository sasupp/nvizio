package service

import (
	"sort"

	"xtrinio.com/lib/table"
	"xtrinio.com/model"
	"xtrinio.com/repository"
)

type INseInsiderService interface {
	GetNseInsiderByCompany(ticker string) ([]model.NseInsiderDto, error)
	ToHtmlTable(entries []model.NseInsiderDto) *table.HtmlTable
}

type nseInsiderServiceImpl struct {
	repo repository.NseInsiderRepository
}

func (a nseInsiderServiceImpl) GetNseInsiderByCompany(ticker string) ([]model.NseInsiderDto, error) {
	insiderInfo, err := a.repo.GetNseCompanyInsider(ticker)
	if err != nil {
		return nil, err
	}
	sort.Sort(model.SortNseInsiderByPublishTime(insiderInfo))
	return insiderInfo, nil
}

func NewNseInsiderService(repo repository.NseInsiderRepository) INseInsiderService {
	return nseInsiderServiceImpl{
		repo: repo,
	}
}

func (a nseInsiderServiceImpl) ToHtmlTable(entries []model.NseInsiderDto) *table.HtmlTable {
	builder := table.NewTableBuilder()
	builder.Add(table.InsiderHeader())
	builder.Add(table.InsiderRows(entries))
	return builder.Build()
}
