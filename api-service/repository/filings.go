package repository

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"io"

	"xtrinio.com/db"
	"xtrinio.com/model"
)

type FilingsRepository struct {
	db *db.Db
}

func NewFilingsRepository(db *db.Db) FilingsRepository {
	return FilingsRepository{
		db: db,
	}
}

func UnmarshalData(data *[]byte, isCompressed bool) (*model.DocumentData, error) {
	if data == nil {
		return nil, nil
	}

	if isCompressed {
		r, err := gzip.NewReader(bytes.NewReader(*data))
		if err != nil {
			return nil, err
		}
		defer r.Close()
		out, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		var docData model.DocumentData
		err = json.Unmarshal(out, &docData)
		if err != nil {
			return nil, err
		}
		return &docData, nil
	} else {
		var docData model.DocumentData
		err := json.Unmarshal(*data, &docData)
		if err != nil {
			return nil, err
		}
		return &docData, nil
	}
}

func (f FilingsRepository) GetFilings(company_id string) ([]model.FilingsDto, error) {
	return f.db.QueryFilings(context.Background(), company_id)
}

func (f FilingsRepository) GetFilingDocuments(filingId string) ([]model.FilingDocumentsDto, error) {
	return f.db.QueryFilingDocuments(context.Background(), filingId)
}

func (f FilingsRepository) GetFilingDocumentsWithType(filingId string, docType string) ([]model.FilingDocumentsDto, error) {
	return f.db.QueryFilingDocumentsByType(context.Background(), filingId, docType)
}

func BuildFilingDocument(documents []model.FilingDocumentsDto) ([]model.FilingDocuments, error) {
	var result []model.FilingDocuments
	for _, document := range documents {

		data, err := UnmarshalData(document.Data, document.IsCompressed)
		if err != nil {
			return nil, err
		}

		result = append(result, model.FilingDocuments{
			FilingId:     document.FilingId,
			DocId:        document.DocId,
			Name:         document.Name,
			Type:         document.Type,
			Tag:          document.Tag,
			Title:        document.Title,
			IsCompressed: document.IsCompressed,
			Data:         data,
		})
	}
	return result, nil
}

func (f FilingsRepository) GetFilingDocument(filingId string, docId int) ([]model.FilingDocuments, error) {
	documents, err := f.db.QueryFilingDocument(context.Background(), filingId, docId)
	if err != nil {
		return nil, err
	}
	return BuildFilingDocument(documents)
}
