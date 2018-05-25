package usecases

import "github.com/koneko096/openshop/models/datamodels"

type UserManager interface {
	GetAll() []datamodels.User
	GetByID(id int64) (datamodels.User, bool)
	GetByToken(token string) (datamodels.User, error)
	InsertOrUpdate(user datamodels.User) (datamodels.User, error)
	DeleteByID(id int64) bool
}
