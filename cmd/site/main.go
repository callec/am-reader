package main

import (
	"database/sql"
	"log"
	"mag"
	"mag/internal/chttp"
	"mag/internal/html"
	"mag/internal/nofs"
	"mag/service"
	"net/http"
	"os"

	_ "modernc.org/sqlite" //"github.com/mattn/go-sqlite3"
)

var (
	dbloc = "./database/mg.db"
)

func initDB(loc string) (service.Service, error) {
	d, err := sql.Open("sqlite", "./database/mg.db")
	if err != nil {
		return nil, err
	}
	err = service.InitSQL(d)
	if err != nil {
		return nil, err
	}

	queries := service.Queries(d)
	return service.NewService(queries), nil
}

func newGenericLogger(title string) *log.Logger {
	return log.New(os.Stdout, title+" ", log.Ldate|log.Ltime)
}

// TODO: It's all spaghetti here.
func main() {
	// Initialise database.
	s, err := initDB(dbloc)
	if err != nil {
		log.Fatal(err)
	}

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
	logger := newGenericLogger("Server")
	loggingMW := chttp.NewLogger(logger)
	authMW := chttp.NewAuth(s)
	errorMW := chttp.NewError(errRender)

	// Handlers.
	nfslogger := newGenericLogger("NoBrowseFS")
	nfs := nofs.NoBrowseFS{Fs: http.FS(mag.Content()), Logger: nfslogger}
	handler := http.FileServer(nfs)
	mux.Handle("/", handler)
	mux.Handle("/uploads/",
		http.StripPrefix("/uploads/",
			http.FileServer(nofs.NoBrowseFS{
				Fs:     http.Dir("./uploads"),
				Logger: nfslogger,
			})))

	// Spaghetti to avoid dependencies between packages.
	homeRender := func(w http.ResponseWriter, ms []*mag.Magazine) error {
		params := html.MainPageParams{
			Title:   "MAIN PAGE",
			Results: ms,
		}
		return html.MainPage(w, params)
	}
	mux.HandleFunc("/main/", chttp.HomeHandler(s, homeRender))

	viewRender := func(w http.ResponseWriter, m *mag.Magazine) error {
		params := html.ViewPageParams{
			Title:    "VIEWER",
			Magazine: m,
		}
		return html.ViewPage(w, params)
	}
	mux.HandleFunc("/viewer/", chttp.ViewHandler(s, viewRender))

	adminRender := func(w http.ResponseWriter, msg string) error {
		params := html.AdminPageParams{
			Title:    "ADMIN",
			Verified: true,
			Message:  msg,
		}
		return html.AdminPage(w, params)
	}
	mux.Handle("/admin/", authMW(chttp.AdminHandler(s, adminRender)))

	loginRender := func(w http.ResponseWriter) error {
		params := html.LoginPageParams{Title: "LOGIN"}
		return html.LoginPage(w, params)
	}
	mux.HandleFunc("/login/", chttp.LoginHandler(s, loginRender))
	mux.HandleFunc("/login_process/", chttp.LoginProcessHandler(s))

	// IMPORTANT: Registration of new users should _only_ be performed by an admin.
	// regRender := func(w http.ResponseWriter) error {
	// 	params := html.RegPageParams{Title: "REGISTRATION"}
	// 	return html.RegPage(w, params)
	// }
	// mux.HandleFunc("/register/", chttp.RegisterHandler(s, regRender, errRender))

	mux.Handle("/register_process/", authMW(chttp.RegisterProcessHandler(s)))
	mux.Handle("/magazine_upload/", authMW(chttp.UploadHandler(s)))
	mux.Handle("/magazine_delete/", authMW(chttp.DeleteHandler(s)))

	// Wrap in middleware.
	wrappedMux := loggingMW(errorMW(mux))

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
