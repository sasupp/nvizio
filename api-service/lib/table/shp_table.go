package table

import (
	"xtrinio.com/helper"
	"xtrinio.com/model"
)

func ShpHeader() TableOperation {
	row := Row(FromCells([]string{"Date", "Promoter", "Public", "Employee Trusts"}))
	return func(tb *TableBuilder) {
		tb.table.Header = row
	}
}

func ShpRows(entries []model.NseShareholdingDto) TableOperation {
	rows := make([]HtmlRow, 0)
	for _, e := range entries {
		rowOp := FromCells([]string{
			e.AsOnDate,
			helper.ToString(e.PromoterGroup),
			helper.ToString(e.Public),
			helper.ToString(e.EmployeeTrusts),
		})
		rows = append(rows, Row(rowOp))
	}
	return func(tb *TableBuilder) {
		tb.table.Body = append(tb.table.Body, rows...)
	}
}
