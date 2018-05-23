package dao

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

type UserDAO interface {
	Select(query Query) (model datamodels.User, found bool)
	SelectMany(query Query, limit int) (results []datamodels.User)

	InsertOrUpdate(model datamodels.User) (datamodels.User, error)
	Delete(query Query) (deleted bool)
}

type userRepository struct {
	source *gorm.DB
}

func NewUserDAO(connection *gorm.DB) UserDAO {
	return &userRepository{source: connection}
}

func (r *userRepository) Select(query Query) (datamodels.User, bool) {
	user := datamodels.User{}
	if err := r.source.Where(query).First(&user).Error; err != nil {
		return datamodels.User{}, false
	}
	return user, true
}

func (r *userRepository) SelectMany(query Query, limit int) (results []datamodels.User) {
	users := new([]datamodels.User)
	r.source.Where(query).Find(&users).Limit(limit)
	return *users
}

func (r *userRepository) InsertOrUpdate(user datamodels.User) (_ datamodels.User, err error) {
	var oldUser datamodels.User
	if err := r.source.First(&oldUser, user.ID).Error; err != nil {
		r.source.Create(&user)
	} else {
		r.source.Model(&oldUser).Update(&user)
	}

	return user, err
}

func (r *userRepository) Delete(query Query) bool {
	err := r.source.Delete(datamodels.User{}, query).Error
	return err == nil
}
