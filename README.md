# Basic PDF Viewer
Complete website for basic PDF viewing.
Currently supports:
- Library page
- Two page PDF reader

## Basics
The server is written in golang.
The code for the executable is located in `cmd/site/`, and its dependencies are factored out to `internal/<package name>`.
Packages in `internal/` should (preferably) have no dependencies between each other.
Additional packages and code that are used globally is located in the root directory.

## Running
```
git clone git@github.com:callec/pdf-reader.git
cd pdf-reader
go get .
go run ./...
```

## Dependencies
golang, sqlite3, sqlc (if you want to add/edit queries).

## TODO

### Actual TODO
- Uploading and deletion of PDFs
  - Scanning PDFs for unwanted content (can use `nsfwjs` but I think it will be VERY slow for long PDFs).
- Testing so that everything works properly in case of change
- CSS & js
  - Make site look nice.
  - Make buttons work as intended.
  - Make PDFs render properly.
- CSP (and other web security stuff)
- Docker
- Make sure it runs properly on actual webserver
  - Concurrency?
  - Uptime?
  - Memory leaks?

### Maybe TODO
- Search
  - Scan PDF files for words and search using sqlite3's `LIKE`.
  - Can use `tesseract` (`gosseract` library?) to scan PDFs.
  - Further perhaps implement some similarity system to score search queries.
- Extended library
  - Folders.
  - Custom images.
