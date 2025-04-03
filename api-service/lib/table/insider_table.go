package table

import (
	"fmt"

	"xtrinio.com/model"
)

func InsiderHeader() TableOperation {
	row := Row(FromCells([]string{"Date", "Type", "Value (Lakhs)", "Person", "Person Category"}))
	return func(tb *TableBuilder) {
		tb.table.Header = row
	}
}

func InsiderRows(entries []model.NseInsiderDto) TableOperation {
	rows := make([]HtmlRow, 0)
	for _, e := range entries {
		rowOp := FromCells([]string{
			e.FilingDate,
			e.TransactionType,
			fmt.Sprintf("%0.1f", e.TransactionValue/(100000.0)),
			e.Name,
			e.PersonCategory,
		})
		rows = append(rows, Row(rowOp))
	}
	return func(tb *TableBuilder) {
		tb.table.Body = append(tb.table.Body, rows...)
	}
}
