package auth

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

const sessionExpiration = 60 * time.Minute

type Session struct {
	uname  string
	expiry time.Time
}

type SessionStore struct {
	store map[uuid.UUID]Session
	mu    sync.RWMutex
}

func newSessionStore() *SessionStore {
	ss := &SessionStore{
		store: make(map[uuid.UUID]Session),
	}
	go ss.startCleanupRoutine()
	return ss
}

func (s *SessionStore) startCleanupRoutine() {
	ticker := time.NewTicker(5 * time.Minute) // cleanup every 5 minutes
	defer ticker.Stop()

	for range ticker.C {
		s.cleanup()
	}
}

func (s *SessionStore) cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for id, session := range s.store {
		if now.After(session.expiry) {
			delete(s.store, id)
		}
	}
}

// Store stores a new session with a given id and username.
func (s *SessionStore) Store(id uuid.UUID, username string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.store[id]; ok {
		return fmt.Errorf("session id %v already exists", id)
	}

	expiry := time.Now().Add(sessionExpiration)
	s.store[id] = Session{username, expiry}
	return nil
}

// Get retrieves a username by session id.
func (s *SessionStore) Get(id uuid.UUID) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.store[id]
	if !ok {
		return "", false
	}

	if time.Now().After(session.expiry) {
		return "", false
	}

	return session.uname, true
}

// Delete removes a session by id.
func (s *SessionStore) Delete(id uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.store, id)
}
