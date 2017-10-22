package dao

import (
	"github.com/icalF/openshop/models"
	"github.com/jinzhu/gorm"
)

type UserDAO interface {
	Select(query Query) (model models.User, found bool)
	SelectMany(query Query, limit int) (results []models.User)

	InsertOrUpdate(model models.User) (models.User, error)
	Delete(query Query) (deleted bool)
}

type userRepository struct {
	source *gorm.DB
}

func NewUserDAO(connection *gorm.DB) UserDAO {
	return &userRepository{source: connection}
}

func (r *userRepository) Select(query Query) (models.User, bool) {
	user := models.User{}
	if err := r.source.Where(query).First(&user).Error; err != nil {
		return models.User{}, false
	}
	return user, true
}

func (r *userRepository) SelectMany(query Query, limit int) (results []models.User) {
	users := new([]models.User)
	r.source.Where(query).Find(&users).Limit(limit)
	return *users
}

func (r *userRepository) InsertOrUpdate(user models.User) (_ models.User, err error) {
	var oldUser models.User
	if err := r.source.First(&oldUser).Error; err != nil {
		r.source.Create(&user)
	} else {
		r.source.Model(&oldUser).Update(&user)
	}

	return user, err
}

func (r *userRepository) Delete(query Query) bool {
	if err := r.source.Delete(models.User{}, query).Error; err != nil {
		return false
	}
	return true
}
