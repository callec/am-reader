package auth

import (
	"log"
	"mag/service"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type BasicAuthoriser struct {
	d service.Service
}

func (b *BasicAuthoriser) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// handle error
	}

	user := &service.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	// store user in the database
}

func (b *BasicAuthoriser) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// retrieve the user with this username from the database
	ctx, cf := getTimedContext(300)
	defer cf()

	user, err := b.d.GetUserByName(ctx, username)
	if err != nil {
		log.Printf("LoginHandler: Login unsuccessful: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		// password is incorrect, handle error
	}

	// password is correct, log the user in
}

func (b *BasicAuthoriser) RequireLogin(next http.Handler) http.Handler {
	return nil
}
