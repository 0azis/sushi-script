package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sushi/helpers"
	"sushi/internal/core/domain"
	"sushi/internal/core/ports"
)

type authControllers struct {
	UserService ports.UserService
}

func (ac *authControllers) Login(c *gin.Context) {
	var credentials domain.User

	err := c.BindJSON(&credentials)
	if err != nil || !credentials.Validate() {
		c.JSON(400, domain.HttpResponse{
			Error: domain.Error{
				Status:  400,
				Message: "Bad Request",
			},
		})
		return
	}

	ID, err := ac.UserService.LoginUser(credentials)

	if err != nil {
		c.JSON(500, domain.HttpResponse{
			Error: domain.Error{
				Status:  500,
				Message: "An error while login in account",
			},
		})
		return
	}

	jwt, err := helpers.SignJWT(ID)
	if err != nil {
		c.JSON(500, domain.HttpResponse{
			Error: domain.Error{
				Status:  500,
				Message: "An error while login in account",
			},
		})
		return
	}

	c.SetCookie("token", jwt, 100, "/", "localhost", false, false)
	c.JSON(200, domain.HttpResponse{
		Error: domain.Error{
			Status:  200,
			Message: "You are successfully logged in",
		},
	})
}

func (ac *authControllers) UpdateProfile(c *gin.Context) {
	var credentials domain.User

	err := c.BindJSON(&credentials)
	if err != nil || !credentials.Validate() {
		c.JSON(400, domain.HttpResponse{
			Error: domain.Error{
				Status:  400,
				Message: "Bad Request",
			},
		})
		return
	}

	token, _ := c.Cookie("token")
	userID, err := helpers.GetIdentity(token)

	credentials.ID = userID

	err = ac.UserService.UpdateProfile(credentials)
	if err != nil {
		c.JSON(500, domain.HttpResponse{
			Error: domain.Error{
				Status:  500,
				Message: "Internal Error",
			},
		})
		return
	}

	c.JSON(200, domain.HttpResponse{
		Error: domain.Error{
			Status:  200,
			Message: "Profile was successfully updated",
		},
	})
}

func (ac *authControllers) GetMyProfile(c *gin.Context) {
	token, _ := c.Cookie("token")
	userID, err := helpers.GetIdentity(token)

	if err != nil {
		c.JSON(500, domain.HttpResponse{
			Error: domain.Error{
				Status:  500,
				Message: "Internal Error",
			},
		})
		return
	}

	user, err := ac.UserService.Profile(userID)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, domain.HttpResponse{
			Error: domain.Error{
				Status:  500,
				Message: "Internal Error",
			},
		})
		return
	}

	c.JSON(200, domain.HttpResponse{
		Error: domain.Error{
			Status:  200,
			Message: "User found",
		},
		Data: user,
	})

}

func NewUserAdapters(service ports.UserService) authControllers {
	return authControllers{
		UserService: service,
	}
}
