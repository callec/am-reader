package html

import "io"

var loginPage = parse("login.html")

type LoginPageParams struct {
	Title string
}

func LoginPage(w io.Writer, p LoginPageParams) error {
	return loginPage.Execute(w, p)
}
