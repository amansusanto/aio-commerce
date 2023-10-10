package helpers

import (
	UserRepository "aio-commerce/internal/modules/user/repositories"
	UserResponse "aio-commerce/internal/modules/user/responses"
	"aio-commerce/pkg/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	var userRepo = UserRepository.New()
	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return response
	}

	return UserResponse.ToUser(user)
}
