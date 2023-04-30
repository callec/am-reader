package html

import "io"

var errorPage = parse("error.html")

type ErrorPageParams struct {
	Title string
	Err   string
}

func ErrorPage(w io.Writer, p ErrorPageParams) error {
	return errorPage.Execute(w, p)
}
