package html

import (
	"io"
	"mag/magazine"
)

var viewPage = parse("viewer.html")

type ViewPageParams struct {
	Title    string
	Magazine *magazine.Magazine
}

func ViewPage(w io.Writer, p ViewPageParams) error {
	return viewPage.Execute(w, p)
}
