package auth

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type SessionStore struct {
	store map[uuid.UUID]string
	mu    sync.RWMutex
}

func newSessionStore() *SessionStore {
	return &SessionStore{
		store: make(map[uuid.UUID]string),
	}
}

// Store stores a new session with a given id and username.
func (s *SessionStore) Store(id uuid.UUID, username string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.store[id]; ok {
		return fmt.Errorf("session id %v already exists", id)
	}

	s.store[id] = username
	return nil
}

// Get retrieves a username by session id.
func (s *SessionStore) Get(id uuid.UUID) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	username, ok := s.store[id]
	return username, ok
}

// Delete removes a session by id.
func (s *SessionStore) Delete(id uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.store, id)
}
