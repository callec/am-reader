package main

import (
	"database/sql"
	"fmt"
	"mag"
	"mag/internal/chttp"
	"mag/internal/html"
	"mag/internal/nofs"
	"mag/service"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: It's all spaghetti here.
func main() {
	// Initialise database.
	d, err := sql.Open("sqlite3", "./database/mg.db")
	if err != nil {
		fmt.Printf("TODO ERROR")
		return
	}
	err = service.InitSQL(d)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	queries := service.Queries(d)
	s := service.NewService(queries)

	errRender := func(w http.ResponseWriter, e error) error {
		params := html.ErrorPageParams{
			Title: "ERROR",
			Err:   e.Error(),
		}
		return html.ErrorPage(w, params)
	}

	handler := http.FileServer(nofs.NoBrowseFS{Fs: http.FS(mag.Content())})
	http.Handle("/", handler)

	// Spaghetti to avoid dependencies between packages.
	homeRender := func(w http.ResponseWriter, ms []*mag.Magazine) error {
		params := html.MainPageParams{
			Title:   "MAIN PAGE",
			Results: ms,
		}
		return html.MainPage(w, params)
	}
	http.HandleFunc("/main/", chttp.HomeHandler(s, homeRender, errRender))

	viewRender := func(w http.ResponseWriter, m *mag.Magazine) error {
		params := html.ViewPageParams{
			Title:    "VIEWER",
			Magazine: m,
		}
		return html.ViewPage(w, params)
	}
	http.HandleFunc("/viewer/", chttp.ViewHandler(s, viewRender, errRender))

	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		params := html.AdminPageParams{
			Title:    "ADMIN",
			Verified: true,
		}
		html.AdminPage(w, params)
	})

	http.ListenAndServe(":8080", nil)
}
