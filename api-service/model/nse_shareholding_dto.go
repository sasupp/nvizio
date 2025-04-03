package model

import "time"

type NseShareholdingDto struct {
	FilingDate     string         `json:"filing_date"`
	PublishTime    NsePublishTime `json:"publish_time"`
	Id             string         `json:"id"`
	Ticker         string         `json:"ticker"`
	CompanyName    string         `json:"company_name"`
	AsOnDate       string         `json:"as_on_date"`
	XbrlLink       string         `json:"xbrl_link"`
	PromoterGroup  float32        `json:"promoter_group"`
	Public         float32        `json:"public"`
	EmployeeTrusts float32        `json:"employee_trusts"`
}

// ByPublishTime implements sort.Interface based on the PublishTime field.
type SortNseShareholdingByPublishTime []NseShareholdingDto

func (a SortNseShareholdingByPublishTime) Len() int {
	return len(a)
}
func (a SortNseShareholdingByPublishTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortNseShareholdingByPublishTime) Less(i, j int) bool {
	return time.Time(a[j].PublishTime).Before(time.Time(a[i].PublishTime))
}
