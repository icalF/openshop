package dao

import (
	"errors"
	"sync"

	"github.com/imdario/mergo"

	"github.com/icalF/openshop/models"
)

type Query func(user models.User) bool

type UserDAO interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (user models.User, found bool)
	SelectMany(query Query, limit int) (results []models.User)

	InsertOrUpdate(user models.User) (updatedMovie models.User, err error)
	Delete(query Query, limit int) (deleted bool)
}

func NewUserDAO(source map[int64]models.User) UserDAO {
	return &userMemoryRepository{source: source}
}

type userMemoryRepository struct {
	source map[int64]models.User
	mu     sync.RWMutex
}

const (
	ReadOnlyMode = iota
	ReadWriteMode
)

func (r *userMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, user := range r.source {
		ok = query(user)
		if ok {
			if action(user) {
				if actionLimit >= loops {
					break
				}
			}
		}
	}

	return
}

func (r *userMemoryRepository) Select(query Query) (user models.User, found bool) {
	found = r.Exec(query, func(m models.User) bool {
		user = m
		return true
	}, 1, ReadOnlyMode)

	if !found {
		user = models.User{}
	}

	return
}

func (r *userMemoryRepository) SelectMany(query Query, limit int) (results []models.User) {
	r.Exec(query, func(m models.User) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

func (r *userMemoryRepository) InsertOrUpdate(user models.User) (models.User, error) {
	id := user.ID

	if id == 0 {
		var lastID int64

		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		user.ID = id

		r.mu.Lock()
		r.source[id] = user
		r.mu.Unlock()

		return user, nil
	}

	current, exists := r.Select(func(m models.User) bool {
		return m.ID == id
	})
	if !exists {
		return models.User{}, errors.New("failed to update a nonexistent user")
	}

	r.mu.Lock()
	err := mergo.MergeWithOverwrite(&current, user)
	r.source[id] = current
	r.mu.Unlock()

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m models.User) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
