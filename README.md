# Basic PDF Viewer
Complete website for basic PDF viewing.
Currently supports:
- Library page
- Two page PDF reader

## Running
```
git clone git@github.com:callec/pdf-reader.git
cd pdf-reader
go get .
go run ./...
```

## Dependencies
sqlite3, sqlc (if you want to add/edit queries).

## TODO

### Actual TODO
- Uploading and deletion of PDFs
  - Proper log in system (i.e. hashing, salting, and storage of passwords).
  - Scanning PDFs for unwanted content (can use `nsfwjs` but I think it will be VERY slow for long PDFs).
- CSS & js
  - Make site look nice.
  - Make buttons work as intended.
  - Make PDFs render properly.
- Middleware
  - Logging etc
- CSP (and other web security stuff)
- Documentation
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
