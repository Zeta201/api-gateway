package util

import (
	jwtservice "github.com/Zeta201/identity-server/jwt-service"
	"github.com/Zeta201/identity-server/model"
)

func BindClaimsToUser(claims jwtservice.UserClaims) model.User {
	return model.User{
		ID:       claims.ID,
		Username: claims.Username,
		Password: claims.Password,
	}
}
