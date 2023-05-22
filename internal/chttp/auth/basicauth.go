package auth

import (
	"log"
	"mag/service"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type BasicAuthoriser struct {
	d            service.Service
	emptyFun     func(http.ResponseWriter, error) error
	sessionStore SessionStore
}

func NewBasicAuthoriser(
	d service.Service,
	ef func(http.ResponseWriter, error) error,
) *BasicAuthoriser {
	return &BasicAuthoriser{
		d:            d,
		emptyFun:     ef,
		sessionStore: *newSessionStore(),
	}
}

func (b *BasicAuthoriser) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		b.emptyFun(w, err)
		return
	}

	ctx, cf := getTimedContext(300)
	defer cf()
	b.d.RegisterUser(ctx, username, string(hashedPassword))
}

func (b *BasicAuthoriser) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	ctx, cf := getTimedContext(300)
	defer cf()

	user, err := b.d.GetUserByName(ctx, username)
	if err != nil {
		log.Printf("LoginHandler: Login unsuccessful: %s", err.Error())
		b.emptyFun(w, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Printf("LoginHandler: Authentication error: %s", err.Error())
		b.emptyFun(w, err)
		return
	}

	// password is correct, log the user in.
	sessionID := uuid.New()

	err = b.sessionStore.Store(sessionID, username)
	if err != nil {
		log.Printf("LoginHandler: Error storing session: %s", err.Error())
		b.emptyFun(w, err)
		return
	}

	// Send the session ID to the client in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (b *BasicAuthoriser) validateSessionToken(value string) bool {
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

func (b *BasicAuthoriser) isUserLoggedIn(r *http.Request) bool {
	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		return false
	}

	return b.validateSessionToken(sessionTokenCookie.Value)
}

func (b *BasicAuthoriser) RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !b.isUserLoggedIn(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
