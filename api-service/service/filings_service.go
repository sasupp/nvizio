package service

import (
	"sort"

	"xtrinio.com/model"
	"xtrinio.com/repository"
)

type IFilingsService interface {
	GetFilings(company_id string) ([]model.FilingsDto, error)
	GetFilingDocuments(filingId string) ([]model.FilingDocumentsDto, error)
	GetFilingDocumentsWithType(filingId string, docType string) ([]model.FilingDocumentsDto, error)
	GetFilingDocument(filingId string, docId int) ([]model.FilingDocuments, error)
}

type filingsServiceImpl struct {
	repo repository.FilingsRepository
}

func NewFilingsService(repo repository.FilingsRepository) IFilingsService {
	return filingsServiceImpl{
		repo: repo,
	}
}

func (f filingsServiceImpl) GetFilings(company_id string) ([]model.FilingsDto, error) {
	filings, err := f.repo.GetFilings(company_id)
	if err != nil {
		return []model.FilingsDto{}, err
	}

	sort.Slice(filings, func(i, j int) bool {
		return filings[i].FilingDate.After(filings[j].FilingDate)
	})
	return filings, nil
}

func (f filingsServiceImpl) GetFilingDocuments(filingId string) ([]model.FilingDocumentsDto, error) {
	return f.repo.GetFilingDocuments(filingId)
}

func (f filingsServiceImpl) GetFilingDocumentsWithType(filingId string, docType string) ([]model.FilingDocumentsDto, error) {
	return f.repo.GetFilingDocumentsWithType(filingId, docType)
}

func (f filingsServiceImpl) GetFilingDocument(filingId string, docId int) ([]model.FilingDocuments, error) {
	return f.repo.GetFilingDocument(filingId, docId)
}
