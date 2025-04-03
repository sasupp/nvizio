package table

type HtmlCell struct {
	Text string
	Href string
}

type HtmlRow struct {
	Cells []HtmlCell
	Style string
}

type HtmlTable struct {
	Header HtmlRow
	Body   []HtmlRow
}

func Row(ops ...RowOperation) HtmlRow {
	row := HtmlRow{Cells: []HtmlCell{}}
	for _, op := range ops {
		op(&row)
	}
	return row
}

func StyledRow(style string, ops ...RowOperation) HtmlRow {
	row := HtmlRow{Style: style, Cells: []HtmlCell{}}
	for _, op := range ops {
		op(&row)
	}
	return row
}

type TableBuilder struct {
	table HtmlTable
}

type TableOperation func(*TableBuilder)

func Header(row HtmlRow) TableOperation {
	return func(tb *TableBuilder) {
		tb.table.Header = row
	}
}

func TableRow(row HtmlRow) TableOperation {
	return func(tb *TableBuilder) {
		if len(row.Cells) > 0 {
			tb.table.Body = append(tb.table.Body, row)
		}
	}
}

func NewTableBuilder() *TableBuilder {
	return &TableBuilder{
		table: HtmlTable{
			Body: []HtmlRow{},
		},
	}
}

func (tb *TableBuilder) Add(ops ...TableOperation) {
	for _, op := range ops {
		op(tb)
	}
}

func (tb *TableBuilder) Build() *HtmlTable {
	return &tb.table
}
