package table

type RowOperation func(*HtmlRow)

func FromCells(texts []string) RowOperation {
	cells := make([]HtmlCell, len(texts))
	for i, text := range texts {
		cells[i] = HtmlCell{Text: text}
	}
	return func(r *HtmlRow) {
		r.Cells = cells
	}
}

func Cell(text string) RowOperation {
	return func(r *HtmlRow) {
		r.Cells = append(r.Cells, HtmlCell{Text: text})
	}
}

func CellHref(text string, href string) RowOperation {
	return func(r *HtmlRow) {
		r.Cells = append(r.Cells, HtmlCell{Text: text, Href: href})
	}
}
