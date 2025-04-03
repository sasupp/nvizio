package service

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"xtrinio.com/helper"
	"xtrinio.com/lib/table"
	"xtrinio.com/model"
)

type ITransformService interface {
	TransformFilingsHtml(ticker string, filings []model.FilingsDto) (*table.HtmlTable, error)
	TransformDocumentsHtml(ticker string, filings []model.FilingDocumentsDto) (*table.HtmlTable, error)
	TransformDocumentDataHtml(ticker string, docs []model.FilingDocuments) ([]*table.HtmlTable, error)
}

type transformServiceImpl struct {
	numPrinter *message.Printer
}

func (a transformServiceImpl) TransformFilingsHtml(ticker string, filings []model.FilingsDto) (*table.HtmlTable, error) {
	builder := table.NewTableBuilder()
	builder.Add(table.FilingHeader())
	for _, filing := range filings {
		builder.Add(table.FilingRow(ticker, filing))
	}
	return builder.Build(), nil
}

func (a transformServiceImpl) TransformDocumentsHtml(ticker string, docs []model.FilingDocumentsDto) (*table.HtmlTable, error) {
	builder := table.NewTableBuilder()
	builder.Add(table.DocumentsHeader())
	for _, doc := range docs {
		if doc.IsStatement() {
			builder.Add(table.DocumentsRow(ticker, doc))
		}
	}
	return builder.Build(), nil
}

func (a transformServiceImpl) buildRecord(item model.Item, decimal int, format string, prefix string, ignore bool) table.HtmlRow {
	var ops []table.RowOperation

	class := ""
	if item.Concept.IsTotal && !ignore {
		class = "total-row"
	} else if item.Concept.IsAbstract && !ignore {
		class = "abstract-row"
	}

	hasNonNil := false
	if len(item.Facts) > 0 {
		ops = append(ops, table.Cell(prefix+item.Concept.Label))

		for j, fact := range item.Facts {
			if j >= 4 {
				break
			}
			text := ""
			if !fact.IsNil {
				if fact.IsText {
					text = fact.Text
				} else {
					if fact.Value != 0 {
						if !strings.Contains(strings.ToLower(fact.UnitShort), "shares") {
							if false && fact.Value < 0 {
								// TODO: we can also use sec formatting of brackets instead of -1
								fact.Value *= -1
								negatedFormat := "(" + format + ")"
								text = a.numPrinter.Sprintf(negatedFormat, fact.Value*math.Pow10(decimal))
							} else {
								text = a.numPrinter.Sprintf(format, fact.Value*math.Pow10(decimal))
							}
						} else {
							text = a.numPrinter.Sprintf("%0.2f", fact.Value)
						}
					}
				}
				if text != "" {
					hasNonNil = true
				}

			}
			ops = append(ops, table.Cell(text))
		}
	}
	if hasNonNil || class != "" {
		return table.StyledRow(class, ops...)
	} else {
		return table.HtmlRow{}
	}
}

func (a transformServiceImpl) buildHeader(doc model.FilingDocuments, decimals int, unit string) []string {
	var header []string = make([]string, 0)

	cols := 1 + len(doc.Data.Contexts)
	if cols > 5 {
		cols = 5
	}

	title := doc.Title
	if decimals == -3 {
		title = fmt.Sprintf("%s<br />(%s in thousands)", doc.Title, unit)
	} else if decimals == -5 {
		title = fmt.Sprintf("%s<br />(%s in Lakhs)", doc.Title, unit)
	} else if decimals == -6 {
		title = fmt.Sprintf("%s<br />(%s in millions)", doc.Title, unit)
	} else if decimals == -9 {
		title = fmt.Sprintf("%s<br />(%s in billions)", doc.Title, unit)
	}

	header = append(header, title)
	for i, ctx := range doc.Data.Contexts {
		if i >= 4 {
			break
		}
		if ctx.Type == "duration" {
			header = append(header, fmt.Sprintf("%d month<br />%s", ctx.Duration, ctx.EndDate))
		} else {
			header = append(header, ctx.EndDate)
		}
	}
	return header
}

func (a transformServiceImpl) buildSegmentTable(doc model.FilingDocuments) ([]*table.HtmlTable, error) {
	var result []*table.HtmlTable

	// build a table for each context
	for c := range doc.Data.Contexts {
		// rows per axis
		rowsByAxis := make(map[string]map[string]map[int]string)
		indexByAxis := make(map[string][]string)
		for _, item := range doc.Data.Items {
			if len(item.Concept.Dimension) == 1 {
				dim := item.Concept.Dimension[0]
				var rows map[string]map[int]string
				var row map[int]string
				if _, ok := rowsByAxis[dim.Axis]; !ok {
					rowsByAxis[dim.Axis] = make(map[string]map[int]string)
				}
				rows = rowsByAxis[dim.Axis]
				if _, ok := rows[item.Concept.Label]; !ok {
					rows[item.Concept.Label] = make(map[int]string)
				}
				row = rows[item.Concept.Label]
				factValueStr := ""
				if !(item.Facts[c].IsNil) {
					if item.Facts[c].IsText {
						factValueStr = item.Facts[c].Text
					} else {
						factValueStr = a.numPrinter.Sprintf("%0.0f", item.Facts[c].Value)
					}
				}
				if _, ok := indexByAxis[dim.Axis]; !ok {
					indexByAxis[dim.Axis] = make([]string, 0)
				}
				index := 0
				found := false
				for index = 0; index < len(indexByAxis[dim.Axis]); index++ {
					if dim.Member == indexByAxis[dim.Axis][index] {
						found = true
						break
					}
				}
				if !found {
					indexByAxis[dim.Axis] = append(indexByAxis[dim.Axis], dim.Member)

				}
				row[index] = factValueStr
			}
		}
		// if number of tables > 0
		if len(rowsByAxis) > 0 {
			for _, rows := range rowsByAxis {
				if len(rows) > 0 {
					builder := table.NewTableBuilder()
					for s, r := range rows {
						numCols := 1 + len(r)
						var ops []table.RowOperation = make([]table.RowOperation, numCols)
						ops[0] = table.Cell(s)
						// ops = append(ops, table.Cell(s))
						for i, col := range r {
							ops[i+1] = table.Cell(col)
							//ops = append(ops, table.Cell(col))
						}
						builder.Add(table.TableRow(table.Row(ops...)))
					}
					result = append(result, builder.Build())
				}
			}
		}
	}
	return result, nil
}

func (a transformServiceImpl) TransformDocumentDataHtml(ticker string, docs []model.FilingDocuments) ([]*table.HtmlTable, error) {
	var result []*table.HtmlTable

	for _, doc := range docs {
		builder := table.NewTableBuilder()

		tableUnit := helper.CommonUnit(doc.Data.Items)
		tableDecimals, formatStr := -5, "%.1f"

		builder.Add(
			table.Header(
				table.Row(
					table.FromCells(a.buildHeader(doc, tableDecimals, tableUnit)))))

		var disclosureTable *table.HtmlTable

		for _, item := range doc.Data.Items {
			if len(item.Concept.Dimension) == 0 {
				if item.Concept.Name == "DisclosureOfNotesOnFinancialResultsExplanatoryTextBlock" {
					for _, df := range item.Facts {
						if df.IsNil {
							continue
						}
						disclosureBuilder := table.NewTableBuilder()
						disclosureRow := table.FromCells([]string{"Disclosure", df.Text})
						disclosureBuilder.Add(table.TableRow(table.Row(disclosureRow)))
						disclosureTable = disclosureBuilder.Build()
						// result = append(result, disclosureBuilder.Build())
					}
				} else {
					records := a.buildRecord(item, tableDecimals, formatStr, "", false)
					builder.Add(table.TableRow(records))
				}
			}
		}

		result = append(result, builder.Build())
		if disclosureTable != nil {
			result = append(result, disclosureTable)
		}

		// segmentTables, err := a.buildSegmentTable(doc)
		// if err == nil {
		// 	result = append(result, segmentTables...)
		// }
	}
	return result, nil
}

func NewTransformService() ITransformService {
	return transformServiceImpl{
		numPrinter: message.NewPrinter(language.English),
	}
}
