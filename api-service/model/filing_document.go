package model

type Context struct {
	Type      string  `json:"type"`
	EndDate   string  `json:"end_date"`
	StartDate *string `json:"start_date,omitempty"`
	Duration  int     `json:"duration"`
}

// Equals compares two Context structs for equality
func (c Context) Equals(other Context) bool {
	if c.Type != other.Type {
		return false
	}
	if c.EndDate != other.EndDate {
		return false
	}
	if c.Duration != other.Duration {
		return false
	}
	return true
}

type Dimension struct {
	Axis        string `json:"axis"`
	AxisLabel   string `json:"axis_label"`
	Member      string `json:"member"`
	MemberLabel string `json:"member_label"`
}

type Concept struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Balance    string      `json:"balance"`
	Label      string      `json:"label"`
	LabelType  string      `json:"label_type"`
	IsAbstract bool        `json:"is_abstract"`
	IsTotal    bool        `json:"is_total"`
	IsNegated  bool        `json:"is_negated"`
	Dimension  []Dimension `json:"dimensions"`
}

type Fact struct {
	Value     float64 `json:"value"`
	Decimals  int     `json:"decimals"`
	Unit      string  `json:"unit"`
	UnitShort string  `json:"unit_short"`
	IsText    bool    `json:"is_text"`
	Text      string  `json:"text"`
	IsNil     bool    `json:"is_nil"`
}

type Item struct {
	Parent  string  `json:"parent_concept_id"`
	Seq     int     `json:"seq"`
	Level   int     `json:"level"`
	Concept Concept `json:"concept"`
	Facts   []Fact  `json:"facts"`
}

type DocumentData struct {
	Contexts []Context `json:"contexts"`
	Items    []Item    `json:"items"`
}

type FilingDocuments struct {
	FilingId     string        `json:"filing_id"`
	DocId        int           `json:"doc_id"`
	Type         string        `json:"type"`
	Name         string        `json:"name"`
	Tag          string        `json:"tag"`
	Title        string        `json:"title"`
	IsCompressed bool          `json:"-"`
	Data         *DocumentData `json:"data,omitempty"`
}
