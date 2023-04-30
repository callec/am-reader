package main

import (
	"database/sql"
	"fmt"
	"mag"
	"mag/internal/chttp"
	"mag/internal/html"
	"mag/internal/nofs"
	"mag/magazine"
	"mag/magazine/db"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialise database.
	d, err := sql.Open("sqlite3", "./database/mg.db")
	if err != nil {
		fmt.Printf("TODO ERROR")
		return
	}
	err = magazine.InitSQL(d)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	queries := db.New(d)
	s := magazine.NewService(queries)

	ef := func(w http.ResponseWriter, e error) error {
		params := html.ErrorPageParams{
			Title: "ERROR",
			Err:   e.Error(),
		}
		return html.ErrorPage(w, params)
	}

	handler := http.FileServer(nofs.NoBrowseFS{Fs: http.FS(mag.Content())})
	http.Handle("/", handler)

	// Spaghetti to avoid dependencies between packages.
	mrf := func(w http.ResponseWriter, ms []*magazine.Magazine) error {
		params := html.MainPageParams{
			Title:   "MAIN PAGE",
			Results: ms,
		}
		return html.MainPage(w, params)
	}
	mhf := func(w http.ResponseWriter, r *http.Request, s magazine.Service) {
		chttp.HomeHandler(w, r, s, mrf, ef)
	}
	http.HandleFunc("/main/", chttp.MakeHandler(mhf, s))

	vrf := func(w http.ResponseWriter, m *magazine.Magazine) error {
		params := html.ViewPageParams{
			Title:    "VIEWER",
			Magazine: m,
		}
		return html.ViewPage(w, params)
	}
	vhf := func(w http.ResponseWriter, r *http.Request, s magazine.Service) {
		chttp.ViewHandler(w, r, s, vrf, ef)
	}
	http.HandleFunc("/viewer/", chttp.MakeHandler(vhf, s))

	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		params := html.AdminPageParams{
			Title:    "ADMIN",
			Verified: true,
		}
		html.AdminPage(w, params)
	})

	http.ListenAndServe(":8080", nil)
}
