package repository

import (
	"errors"

	"github.com/Zeta201/identity-server/infrastructure"
	jwtservice "github.com/Zeta201/identity-server/jwt-service"
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

func (repository UserRepository) LogIn(user model.User) (string, error) {
	var found model.User
	err := repository.db.DB.Debug().Model(&model.User{}).Where(&user).Take(&found).Error
	if err != nil {
		return "", errors.New("access denied. incorrect credentials")
	}
	signedAccessToken, err := jwtservice.NewAccessToken(found)
	if err != nil {
		return "", errors.New("Error generating the token")
	}

	return signedAccessToken, nil
}

// func (repository UserRepository) Validate(string accessToken) ()
