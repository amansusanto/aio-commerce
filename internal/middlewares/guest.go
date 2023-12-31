package middlewares

import (
	UserRepository "aio-commerce/internal/modules/user/repositories"
	"aio-commerce/pkg/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}
		// before request

		c.Next()
	}
}
