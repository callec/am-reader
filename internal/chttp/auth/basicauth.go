package auth

import (
	"log"
	"mag/service"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type basicAuthoriser struct {
	d            service.Service
	sessionStore SessionStore
}

func NewBasicAuthoriser(
	d service.Service,
) *basicAuthoriser {
	return &basicAuthoriser{
		d:            d,
		sessionStore: *newSessionStore(),
	}
}

func (b *basicAuthoriser) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx, cf := getTimedContext(300)
	defer cf()
	err = b.d.RegisterUser(ctx, username, string(hashedPassword))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login/", http.StatusOK)
}

func (b *basicAuthoriser) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	ctx, cf := getTimedContext(300)
	defer cf()

	user, err := b.d.GetUserByName(ctx, username)
	if err != nil {
		log.Printf("LoginHandler: Login unsuccessful: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Printf("LoginHandler: Authentication error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// password is correct, log the user in.
	sessionID := uuid.New()

	err = b.sessionStore.Store(sessionID, username)
	if err != nil {
		log.Printf("LoginHandler: Error storing session: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the session ID to the client in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}

func (b *basicAuthoriser) validateSessionToken(value string) bool {
	uid, err := uuid.Parse(value)
	if err != nil {
		log.Printf("validateSessionToken: Illegal uuid: %s", err.Error())
		return false
	}

	if _, ok := b.sessionStore.Get(uid); ok {
		return true
	}
	return false
}

func (b *basicAuthoriser) isUserLoggedIn(r *http.Request) bool {
	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		return false
	}

	return b.validateSessionToken(sessionTokenCookie.Value)
}

func (b *basicAuthoriser) RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !b.isUserLoggedIn(r) {
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
