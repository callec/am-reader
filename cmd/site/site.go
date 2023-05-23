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

	// Render error.
	errRender := func(w http.ResponseWriter, e error) error {
		params := html.ErrorPageParams{
			Title: "ERROR",
			Err:   e.Error(),
		}
		return html.ErrorPage(w, params)
	}

	// Mux.
	mux := http.NewServeMux()

	// Middleware.
	var logger log.Logger
	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime)
	loggingMW := chttp.NewLogger(logger)

	authMW := chttp.NewAuth(s, errRender)

	// Handlers.

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

	adminRender := func(w http.ResponseWriter, verified bool) error {
		params := html.AdminPageParams{
			Title:    "ADMIN",
			Verified: verified,
		}
		return html.AdminPage(w, params)
	}
	mux.Handle("/admin/", authMW(chttp.AdminHandler(s, adminRender, errRender)))

	loginRender := func(w http.ResponseWriter) error {
		params := html.LoginPageParams{Title: "LOGIN"}
		return html.LoginPage(w, params)
	}
	mux.HandleFunc("/login/", chttp.LoginHandler(s, loginRender, errRender))
	mux.HandleFunc("/login_process/", chttp.LoginProcessHandler(s, errRender))

	regRender := func(w http.ResponseWriter) error {
		params := html.RegPageParams{Title: "REGISTRATION"}
		return html.RegPage(w, params)
	}
	mux.HandleFunc("/register/", chttp.RegisterHandler(s, regRender, errRender))
	mux.HandleFunc("/register_process/", chttp.RegisterProcessHandler(s, errRender))

	// Wrap in logger.

	wrappedMux := loggingMW(mux)

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
