package service

import (
	"sort"
	"time"

	"xtrinio.com/model"
	"xtrinio.com/repository"
)

const FilingDateFormat = "2006-01-02"

type INseFeedService interface {
	GetNseFeedByDate(date time.Time) ([]model.NseUpdateDto, error)
	GetNseFeedAll() ([]model.NseUpdateDto, error)
	GetNseFeedByCompany(ticker string) ([]model.NseUpdateDto, error)
}

type nseFeedServiceImpl struct {
	repo repository.NseFeedRepository
}

func getLastNDates(startDate time.Time, n int) []string {
	var dates []string
	dates = append(dates, startDate.Format(FilingDateFormat))
	if n > 1 {
		for i := 1; i <= n; i++ {
			date := startDate.AddDate(0, 0, -i)
			dates = append(dates, date.Format(FilingDateFormat))
		}
	}
	return dates
}

func (a nseFeedServiceImpl) GetNseFeedByDate(date time.Time) ([]model.NseUpdateDto, error) {
	filings, err := a.repo.GetNseFeedByDate([]string{date.Format(FilingDateFormat)})
	if err != nil {
		return nil, err
	}
	sort.Sort(model.ByPublishTime(filings))
	return filings, nil
}

func (a nseFeedServiceImpl) GetNseFeedAll() ([]model.NseUpdateDto, error) {
	filings, err := a.repo.GetNseFeedByDate(getLastNDates(time.Now().UTC(), 3))
	if err != nil {
		return nil, err
	}
	sort.Sort(model.ByPublishTime(filings))
	return filings, nil
}

func (a nseFeedServiceImpl) GetNseFeedByCompany(ticker string) ([]model.NseUpdateDto, error) {
	filings, err := a.repo.GetNseCompanyFeed(ticker)
	if err != nil {
		return nil, err
	}
	sort.Sort(model.ByPublishTime(filings))
	return filings, nil
}

func NewNseFeedService(repo repository.NseFeedRepository) INseFeedService {
	return nseFeedServiceImpl{
		repo: repo,
	}
}
