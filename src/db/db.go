package db

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"-"`
	FirstName string
	LastName  string
	Biography string
}

type DB struct {
	mutex sync.Mutex
	Users map[uuid.UUID]User
}

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Biography string
}

type UpdateUserRequest struct {
	FirstName string
	LastName  string
	Biography string
}

func Create() DB {
	return DB{
		Users: make(map[uuid.UUID]User),
	}
}

func (db *DB) Insert(u CreateUserRequest) (User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	id := uuid.New()

	user := User{
		ID:        id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Biography: u.Biography,
	}

	db.Users[id] = user

	return user, nil
}

func (db *DB) FindAll() map[uuid.UUID]User {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Create a copy of the map to avoid concurrent map read/write issues
	usersCopy := make(map[uuid.UUID]User)
	for k, v := range db.Users {
		usersCopy[k] = v
	}

	return usersCopy
}

func (db *DB) FindById(id uuid.UUID) (User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	user, ok := db.Users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (db *DB) Delete(id uuid.UUID) (User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	user, ok := db.Users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	delete(db.Users, id)

	return user, nil
}

func (db *DB) Update(id uuid.UUID, u UpdateUserRequest) (User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	_, ok := db.Users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	user := User{
		ID:        id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Biography: u.Biography,
	}

	db.Users[id] = user

	return user, nil
}
