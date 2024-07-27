package auth

import (
	"net/http"

	"blockaction-api/app/users"

	"github.com/gin-gonic/gin"
)

/* sign in user */
type SignInUserForm struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
}

func SignInUserController(c *gin.Context) {
	logger.Debug("SignInUserController")

	signInUserForm := SignInUserForm{}

	bindErr := c.BindJSON(&signInUserForm)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "username and password are required",
			"msgdata": gin.H{
				"username": signInUserForm.Username,
			},
			"trace": bindErr.Error(),
		})
		return
	}
	user := users.GetUserService(signInUserForm.Username)
	if user.Username == "" || user.Password != signInUserForm.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "username or password is incorrect",
			"msgdata":     nil,
			"trace":       nil,
		})
		return
	}

	token := CreateTokenService(&user.Username)
	c.JSON(http.StatusOK, gin.H{
		"return_code": 0,
		"msgid":       "success sign in",
		"msgdata":     nil,
		"data":        token,
		"data_size":   1,
	})
}

/* sign out user */
func SignOutUserController(c *gin.Context) {
	logger.Debug("SignOutUserController")

	tokenValue, parseErr := ExtractBearerToken(c.GetHeader("Authorization"))

	// no authorization header, nothing to do
	if parseErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"return_code": 0,
			"msgid":       "",
			"msgdata":     nil,
		})
		return
	}

	// delete token with value
	DeleteTokenService(&tokenValue)

	c.JSON(http.StatusOK, gin.H{
		"return_code": 0,
		"msgid":       "success sign out",
		"msgdata":     nil,
	})
}

/* sign up user*/
type SignUpUserForm struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
	Email    string `json:"email" binding:"required" example:"test123@gmail.com"`
}

func SignUpUserController(c *gin.Context) {
	logger.Debug("SignUpUserController")

	// var createUserForm CreateUserForm

	// bindErr := c.BindJSON(&createUserForm)
	// if bindErr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"return_code": 1001,
	// 		"msgid":       "failed to create user with username %s, email %s",
	// 		"msgdata": gin.H{
	// 			"username": createUserForm.Username,
	// 			"email":    createUserForm.Email,
	// 		},
	// 		"trace": bindErr.Error(),
	// 	})
	// 	return
	// }

	// user, createErr := CreateUserService(createUserForm.Username, createUserForm.Password, createUserForm.Email)

	// if createErr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"return_code": 1001,
	// 		"msgid":       "failed to create user with username %s, email %s",
	// 		"msgdata": gin.H{
	// 			"username": createUserForm.Username,
	// 			"email":    createUserForm.Email,
	// 		},
	// 		"trace": createErr.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"return_code": 0,
		"msgid":       "success sign up",
		"msgdata":     nil,
		"data":        nil,
		"data_size":   0,
	})
}
