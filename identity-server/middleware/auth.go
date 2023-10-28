package middleware

import (
	"net/http"
	"strings"

	jwtservice "github.com/Zeta201/identity-server/jwt-service"
	"github.com/Zeta201/identity-server/util"
	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	IDToken string `header:"Authorization"`
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := AuthHeader{}

		err := c.ShouldBindHeader(&h)
		if err != nil {
			util.ErrorJSON(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		idTokenHeader := strings.Split(h.IDToken, "Bearer ")
		if len(idTokenHeader) < 2 {
			util.ErrorJSON(c, http.StatusBadRequest, "Must provide Authorization Header with format `Bearer {token}`")
			c.Abort()
			return
		}

		claims, err := jwtservice.ParseAccessToken(idTokenHeader[1])
		if err != nil {
			util.ErrorJSON(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		boundUser := util.BindClaimsToUser(*claims)
		c.Set("user", boundUser)
		c.Next()
	}
}
