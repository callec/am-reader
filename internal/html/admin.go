package html

import "io"

var adminPage = parse("admin.html")

type AdminPageParams struct {
	Title    string
	Verified bool
}

func AdminPage(w io.Writer, p AdminPageParams) error {
	return adminPage.Execute(w, p)
}
