package storage

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/douglaszuqueto/go-api-boilerplate/pkg/util"
)

// UserMemoryStorage UserMemoryStorage
type UserMemoryStorage struct {
	db sync.Map
}

// NewUserMemoryStorage NewUserMemoryStorage
func NewUserMemoryStorage() *UserMemoryStorage {
	log.Println("Storage: UserMemoryStorage")

	return &UserMemoryStorage{
		db: sync.Map{},
	}
}

// ListUser ListUser
func (s *UserMemoryStorage) ListUser(ctx context.Context) ([]User, error) {
	var l []User
	s.db.Range(func(key interface{}, value interface{}) bool {
		u, ok := value.(User)
		if !ok {
			return false
		}

		// hide the password
		u.Password = ""

		l = append(l, u)
		return true
	})

	return l, nil
}

// GetUser GetUser
func (s *UserMemoryStorage) GetUser(ctx context.Context, id string) (User, error) {
	var l User

	value, ok := s.db.Load(id)
	if !ok {
		return l, errors.New("not found")
	}

	l = value.(User)

	// hide the password
	l.Password = ""

	return l, nil
}

// CreateUser CreateUser
func (s *UserMemoryStorage) CreateUser(ctx context.Context, u User) (string, error) {
	u.ID = util.GenerateID()

	s.db.Store(u.ID, u)

	return u.ID, nil
}

// UpdateUser UpdateUser
func (s *UserMemoryStorage) UpdateUser(ctx context.Context, u User) error {
	s.db.Store(u.ID, u)

	return nil
}

// DeleteUser DeleteUser
func (s *UserMemoryStorage) DeleteUser(ctx context.Context, id string) error {
	s.db.Delete(id)

	return nil
}
