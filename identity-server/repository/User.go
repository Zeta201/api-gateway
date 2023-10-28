package repository

import (
	"github.com/Zeta201/identity-server/infrastructure"
	"github.com/Zeta201/identity-server/model"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repository UserRepository) Save(user model.User) error {
	return repository.db.DB.Create(&user).Error
}

func (repository UserRepository) LogIn(user model.User) (model.User, error) {
	var found model.User
	err := repository.db.DB.Debug().Model(&model.User{}).Where(&user).Take(&found).Error
	return found, err
}
