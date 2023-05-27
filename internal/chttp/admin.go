package chttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mag"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	validFileName = regexp.MustCompile("(.*/)*.+\\.pdf$")
)

func AdminHandler(
	s mag.Service,
	renderFun func(http.ResponseWriter, string) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		message := ""
		cookie, err := r.Cookie("message")
		if err == nil && cookie.Value != "" {
			message = cookie.Value

			// Reset cookie.
			http.SetCookie(w, &http.Cookie{
				Name:   "message",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
			})
		}

		renderFun(w, message) // Middleware does auth.
	}
	return validateHandler(fn)
}

func validateFile(handler *multipart.FileHeader) error {
	m := validFileName.FindStringSubmatch(handler.Filename)
	if m == nil {
		return errors.New("Invalid file format, got: " + filepath.Ext(handler.Filename))
	}
	return nil
}

func getUploadedFile(r *http.Request, elem string) (multipart.File, string, error) {
	// File key == imageFile.
	upload, handler, err := r.FormFile(elem)
	if err != nil {
		return nil, "", err
	}

	err = validateFile(handler)
	if err != nil {
		return nil, "", err
	}

	fname := handler.Filename
	return upload, fname, nil
}

func createFile(fname string, upload multipart.File) (string, error) {
	fext := filepath.Ext(fname)
	newfile, err := os.CreateTemp("/uploads/", "upload-*"+fext)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(newfile, upload)
	if err != nil {
		return "", err
	}

	err = newfile.Sync()
	if err != nil {
		return "", err
	}

	err = newfile.Close()
	if err != nil {
		return "", err
	}

	nname := newfile.Name()

	return nname, nil
}

func UploadHandler(
	s mag.Service,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var message, fname, handler string
		var file multipart.File
		var ctx context.Context
		var cf context.CancelFunc
		var t time.Time

		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		number := r.FormValue("number")
		date := r.FormValue("date")

		i, err := strconv.Atoi(number)
		if err != nil {
			message = err.Error()
			goto redirect
		}

		t, err = time.Parse("2006-01-02T15:04", date)
		if err != nil {
			message = err.Error()
			goto redirect
		}

		err = r.ParseMultipartForm(50 << 20) // Limit upload size.
		if err != nil {
			message = err.Error()
			goto redirect
		}

		file, handler, err = getUploadedFile(r, "magazine")
		if err != nil {
			message = err.Error()
			goto redirect
		}
		defer file.Close()

		fname, err = createFile(handler, file)
		if err != nil {
			message = err.Error()
			goto redirect
		}
		ctx, cf = getTimedContext(300)
		defer cf()

		err = s.AddMagazine(ctx, i, t, fname)
		if err != nil {
			message = err.Error()
			goto redirect
		}
		message = "Upload successful."

	redirect:
		http.SetCookie(w, &http.Cookie{
			Name:  "message",
			Value: message,
			Path:  "/",
		})
		http.Redirect(w, r, "/admin/", http.StatusSeeOther)
	}
	return validateHandler(fn)
}

func deleteMagazineFromDisk(s mag.Service, id uuid.UUID) error {
	ctx, cf := getTimedContext(300)
	defer cf()

	// Get the magazine from the service
	magazine, err := s.GetMagazine(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get magazine: %w", err)
	}

	// Delete the file from the disk
	err = os.Remove(magazine.Location)
	if err != nil {
		return fmt.Errorf("failed to delete magazine file: %w", err)
	}

	return nil
}

func DeleteHandler(
	s mag.Service,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var message, mid string
		var id uuid.UUID
		var err error
		var ctx context.Context
		var cf context.CancelFunc

		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		mid = r.FormValue("mid")
		id, err = uuid.Parse(mid)
		if err != nil {
			message = err.Error()
			goto redirect
		}

		ctx, cf = getTimedContext(300)
		defer cf()

		err = deleteMagazineFromDisk(s, id)
		if err != nil {
			message = err.Error()
			goto redirect
		}

		err = s.RemoveMagazine(ctx, id)
		if err != nil {
			message = err.Error()
			goto redirect
		}
		message = "Deletion successful."

	redirect:
		http.SetCookie(w, &http.Cookie{
			Name:  "message",
			Value: message,
			Path:  "/",
		})
		http.Redirect(w, r, "/admin/", http.StatusSeeOther)
	}
	return validateHandler(fn)
}
