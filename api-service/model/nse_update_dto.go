package model

import (
	"encoding/json"
	"time"
)

type NsePublishTime time.Time

const publishTimeLayout = "02 Jan, 2006 3:04 PM"

func (t NsePublishTime) String() string {
	return time.Time(t).Format(publishTimeLayout)
}

func (t NsePublishTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

type NseUpdateDto struct {
	FilingDate  string         `json:"filing_date"`
	PublishTime NsePublishTime `json:"publish_time"`
	Id          string         `json:"id"`
	Ticker      string         `json:"ticker"`
	CompanyName string         `json:"company_name"`
	Subject     string         `json:"subject"`
	Details     string         `json:"details"`
	Attachment  string         `json:"attachment"`
}

// ByPublishTime implements sort.Interface based on the PublishTime field.
type ByPublishTime []NseUpdateDto

func (a ByPublishTime) Len() int {
	return len(a)
}
func (a ByPublishTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByPublishTime) Less(i, j int) bool {
	return time.Time(a[j].PublishTime).Before(time.Time(a[i].PublishTime))
}
