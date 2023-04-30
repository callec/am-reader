package mag

import (
	"embed"
	"io/fs"
)

// Content returns the website's static content.
func Content() fs.FS {
	return subdir(files, "_content")
}

//go:embed _content
var files embed.FS

func subdir(fsys fs.FS, path string) fs.FS {
	s, err := fs.Sub(fsys, path)
	if err != nil {
		panic(err)
	}
	return s
}
