package http

import (
	"github.com/gin-gonic/gin"
	"strings"
	"sushi/helpers"
	"sushi/internal/core/domain"
	"sushi/internal/core/ports"
	"sushi/internal/core/services"
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

	c.JSON(200, domain.HttpResponse{
		Error: domain.Error{
			Status:  200,
			Message: "You are successfully logged in",
		},
		Data: gin.H{
			"access_token": jwt,
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
	userID, err := helpers.GetIdentity(strings.Split(c.Request.Header["Authorization"][0], " ")[1])

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

	if user.ID == 0 {
		c.JSON(404, domain.HttpResponse{
			Error: domain.Error{
				Status:  404,
				Message: "User not found",
			},
		})
		return
	}

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
			Message: "User found",
		},
		Data: user,
	})

}

func NewUserAdapters(service services.UserService) authControllers {
	return authControllers{
		UserService: &service,
	}
}
