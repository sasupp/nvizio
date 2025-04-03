package model

import "time"

type NseInsiderDto struct {
	FilingDate       string         `json:"filing_date"`
	PublishTime      NsePublishTime `json:"publish_time"`
	Id               string         `json:"id"`
	Ticker           string         `json:"ticker"`
	CompanyName      string         `json:"company_name"`
	Name             string         `json:"name"`
	PersonCategory   string         `json:"person_category"`
	TransactionValue float32        `json:"transaction_value"`
	TransactionType  string         `json:"transaction_type"`
}

// ByPublishTime implements sort.Interface based on the PublishTime field.
type SortNseInsiderByPublishTime []NseInsiderDto

func (a SortNseInsiderByPublishTime) Len() int {
	return len(a)
}
func (a SortNseInsiderByPublishTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortNseInsiderByPublishTime) Less(i, j int) bool {
	return time.Time(a[j].PublishTime).Before(time.Time(a[i].PublishTime))
}
