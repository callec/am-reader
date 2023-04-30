package html

import (
	"io"
	"mag/magazine"
)

var mainPage = parse("library.html", "magazine.html")

type MainPageParams struct {
	Title   string
	Results []*magazine.Magazine
}

func MainPage(w io.Writer, p MainPageParams) error {
	return mainPage.Execute(w, p)
}
