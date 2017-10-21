package services

import (
	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/dao"
)

type UserService interface {
	GetAll() []models.User
	GetByID(id int64) (models.User, bool)
	InsertOrUpdate(user models.User) (models.User, error)
	DeleteByID(id int64) bool
}

func NewUserService(dao dao.UserDAO) UserService {
	return &userService{
		dao: dao,
	}
}

type userService struct {
	dao dao.UserDAO
}

func (s *userService) GetAll() []models.User {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *userService) GetByID(id int64) (models.User, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *userService) InsertOrUpdate(user models.User) (models.User, error) {
	return s.dao.InsertOrUpdate(user)
}

func (s *userService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}