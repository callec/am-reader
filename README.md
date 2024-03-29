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
Before running you will have to create a file `internal/service/setup.sql`, which contains basic setup for a user.
Without this file it will not be possible to access the admin page as registration of new users can only be performed by another user (i.e. an admin).

The server itself only requires docker to be run, and you should run it using the `docker-compose up` command, or `docker-compose up --build` if it needs to be rebuilt.

### `setup.sql` example
This is an example of how setup.sql might look.
```sql
INSERT OR IGNORE
INTO users (id, pwd)
VALUES (
    "cc0ab179-3748-40c2-99e9-c9f29266e6b6",
    "$2a$10$9T/3wAFMWFmWLfL4nYoedOdgmOEteuv/2sAYASrwRbbQQxmeSu0Iq"
);

INSERT OR IGNORE
INTO unames (uid, uname)
VALUES ("cc0ab179-3748-40c2-99e9-c9f29266e6b6", "username");
```

To find uuid and hash you can, for example, use [this](https://go.dev/play/p/ZT_VuZVD4cu) go program.
```go
package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
	uid := uuid.New().String()
	fmt.Println(uid)
}
```

## Dependencies
Docker

## TODO
- Make site look nice
- Search
- Extended library
  - Folders
  - Custom images
