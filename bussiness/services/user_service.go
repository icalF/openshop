package services

import (
	"github.com/koneko096/openshop/dao"
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/bussiness/usecases"
)


func NewUserManager(dao dao.UserDAO) usecases.UserManager {
	return &userService{
		dao: dao,
	}
}


type userService struct {
	dao dao.UserDAO
}



func (s *userService) GetAll() []datamodels.User {
	return s.dao.SelectMany(map[string]string{}, 1)
}

func (s *userService) GetByID(id int64) (datamodels.User, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *userService) GetByToken(token string) (datamodels.User, error) {
	user, found := s.dao.Select(map[string]string{
		"token": token,
	})

	if !found {
		newUser, err := s.InsertOrUpdate(datamodels.NewUser(token))
		if err != nil {
			return datamodels.User{}, err
		}
		user = newUser
	}

	return user, nil
}

func (s *userService) InsertOrUpdate(user datamodels.User) (datamodels.User, error) {
	return s.dao.InsertOrUpdate(user)
}

func (s *userService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}