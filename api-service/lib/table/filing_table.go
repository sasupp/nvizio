package table

import (
	"strconv"
	"strings"

	"xtrinio.com/lib/url"
	"xtrinio.com/model"
)

const (
	YYYYMMDD = "2006-01-02"
)

func docsLink(ticker string, filingId string) string {
	return url.BuildLink(
		url.LinkPath("/nse/company/"),
		url.LinkPath(ticker),
		url.LinkPath("/financials/"),
		url.LinkPath(filingId),
		url.LinkPath("/docs"))
}

func FilingHeader() TableOperation {
	row := Row(FromCells([]string{"FilingDate", "CompanyName", "FiscalYear", "FiscalPeriod", "Consolidated", "Audited", "Doc"}))
	return func(tb *TableBuilder) {
		tb.table.Header = row
	}
}

func FilingRow(ticker string, filing model.FilingsDto) TableOperation {
	class := ""
	if strings.ToUpper(filing.FiscalPeriod) == "FY" {
		class = "highlight-row"
	}

	row := StyledRow(class,
		Cell(filing.FilingDate.Format("02 Jan, 2006")),
		Cell(filing.CompanyName),
		Cell(strconv.Itoa(filing.FiscalYear)),
		Cell(filing.FiscalPeriod),
		Cell(filing.Props["consolidated"]),
		Cell(filing.Props["audited"]),
		CellHref("Link", docsLink(ticker, filing.FilingId)))

	return func(tb *TableBuilder) {
		tb.table.Body = append(tb.table.Body, row)
	}
}
