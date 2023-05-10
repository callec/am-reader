package html

import (
	"io"
	"mag"
)

var mainPage = parse("library.html", "magazine.html")

type MainPageParams struct {
	Title   string
	Results []*mag.Magazine
}

func MainPage(w io.Writer, p MainPageParams) error {
	return mainPage.Execute(w, p)
}
