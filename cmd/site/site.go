package main

import (
	"database/sql"
	"fmt"
	"log"
	"mag"
	"mag/internal/chttp"
	"mag/internal/html"
	"mag/internal/nofs"
	"mag/service"
	"net/http"
	"os"

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

	// Mux.
	mux := http.NewServeMux()

	// Middleware.
	var logger log.Logger
	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime)
	loggingMW := chttp.NewLogger(logger)

	// Handlers.
	errRender := func(w http.ResponseWriter, e error) error {
		params := html.ErrorPageParams{
			Title: "ERROR",
			Err:   e.Error(),
		}
		return html.ErrorPage(w, params)
	}

	handler := http.FileServer(nofs.NoBrowseFS{Fs: http.FS(mag.Content())})
	mux.Handle("/", handler)

	// Spaghetti to avoid dependencies between packages.
	homeRender := func(w http.ResponseWriter, ms []*mag.Magazine) error {
		params := html.MainPageParams{
			Title:   "MAIN PAGE",
			Results: ms,
		}
		return html.MainPage(w, params)
	}
	mux.HandleFunc("/main/", chttp.HomeHandler(s, homeRender, errRender))

	viewRender := func(w http.ResponseWriter, m *mag.Magazine) error {
		params := html.ViewPageParams{
			Title:    "VIEWER",
			Magazine: m,
		}
		return html.ViewPage(w, params)
	}
	mux.HandleFunc("/viewer/", chttp.ViewHandler(s, viewRender, errRender))

	mux.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		params := html.AdminPageParams{
			Title:    "ADMIN",
			Verified: true,
		}
		html.AdminPage(w, params)
	})

	wrappedMux := loggingMW(mux)

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
