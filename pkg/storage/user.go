package storage

import (
	"context"
	"time"
)

// UserStorage UserStorage
type UserStorage interface {
	ListUser(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, id string) (User, error)
	CreateUser(ctx context.Context, u User) (string, error)
	UpdateUser(ctx context.Context, u User) error
	DeleteUser(ctx context.Context, id string) error
}

// User User
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	State     uint32    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db UserStorage

// GetDB getDB
func GetDB() UserStorage {
	return db
}

// Open open
func Open() {
	GetStorageType()
}

// GetStorageType getStorageType
func GetStorageType() {
	storageType := "memory"

	switch storageType {
	case "memory":
		db = NewUserMemoryStorage()
	default:
		panic("unknow storage type: " + storageType)
	}
}
