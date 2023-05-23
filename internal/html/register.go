package html

import "io"

var regPage = parse("register.html")

type RegPageParams struct {
	Title string
}

func RegPage(w io.Writer, p RegPageParams) error {
	return regPage.Execute(w, p)
}
