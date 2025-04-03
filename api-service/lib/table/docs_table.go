package table

import (
	"strconv"

	"xtrinio.com/lib/url"
	"xtrinio.com/model"
)

func docsIdLink(ticker string, filingId string, docId int) string {
	return url.BuildLink(
		url.LinkPath("/nse/company/"),
		url.LinkPath(ticker),
		url.LinkPath("/financials/"),
		url.LinkPath(filingId),
		url.LinkPath("/docs/"),
		url.LinkPath(strconv.Itoa(docId)),
		url.LinkParam("format", "html"))
}

func DocumentsHeader() TableOperation {
	row := Row(FromCells([]string{"Type", "Title"}))
	return func(tb *TableBuilder) {
		tb.table.Header = row
	}
}

func DocumentsRow(ticker string, doc model.FilingDocumentsDto) TableOperation {
	class := ""
	if doc.IsStatement() {
		class = "highlight-row"
	}

	row := StyledRow(class,
		Cell(doc.Type),
		CellHref(doc.Title, docsIdLink(ticker, doc.FilingId, doc.DocId)))

	return func(tb *TableBuilder) {
		tb.table.Body = append(tb.table.Body, row)
	}
}
