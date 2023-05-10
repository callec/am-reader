package html

import (
	"io"
	"mag"
)

var viewPage = parse("viewer.html")

type ViewPageParams struct {
	Title    string
	Magazine *mag.Magazine
}

func ViewPage(w io.Writer, p ViewPageParams) error {
	return viewPage.Execute(w, p)
}
