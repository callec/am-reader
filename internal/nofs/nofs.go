// Custom FS wrapper that does not allow browsing of static files.
// Used to hide static file tree from users.
package nofs

import (
	"log"
	"net/http"
)

// From stackoverflow https://stackoverflow.com/a/51170557.
// Used to block browsing of “public” files.
type NoBrowseFS struct {
	Fs     http.FileSystem
	Logger *log.Logger
}

func (n NoBrowseFS) Open(name string) (result http.File, err error) {
	notFoundFile, notFoundErr := http.Dir("non-existing-path").Open("non-existing-path")

	f, err := n.Fs.Open(name)
	if err != nil {
		n.Logger.Println("Error opening file: ", err)
		return
	}

	fi, err := f.Stat()
	if err != nil {
		n.Logger.Println("Error retrieving file info: ", err)
		return
	}
	if fi.IsDir() {
		return notFoundFile, notFoundErr
	}
	return f, nil
}
