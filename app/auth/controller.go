package auth

import (
	"net/http"

	"blockaction-api/app/users"

	"github.com/gin-gonic/gin"
)

/* sign in user */
type signInUserForm struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
}

func SignInUserController(c *gin.Context) {
	logger.Debug("SignInUserController")

	formData := signInUserForm{}

	bindErr := c.BindJSON(&formData)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "username and password are required",
			"msgdata": gin.H{
				"username": formData.Username,
			},
			"trace": bindErr.Error(),
		})
		return
	}
	user := users.GetUserService(&formData.Username)
	if user.Username == "" || user.Password != formData.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "username or password is incorrect",
			"msgdata":     nil,
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

	tokenValue, extractErr := ExtractBearerToken(c)

	// no authorization header, nothing to do
	if extractErr != nil {
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
type signUpUserForm struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
	Email    string `json:"email" binding:"required" example:"test123@gmail.com"`
}

func SignUpUserController(c *gin.Context) {
	logger.Debug("SignUpUserController")

	var formData = signUpUserForm{}

	// check user input data
	bindErr := c.BindJSON(&formData)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "failed to sign up user with username %s, email %s",
			"msgdata": gin.H{
				"username": formData.Username,
				"email":    formData.Email,
			},
		})
		return
	}

	// create user with users service
	user, createUserErr := users.CreateUserService(formData.Username, formData.Password, formData.Email)
	if createUserErr != nil {
		logger.Debug("createUserErr: ", createUserErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "failed to create user with username %s, email %s",
			"msgdata": gin.H{
				"username": user.Username,
				"email":    user.Email,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"return_code": 0,
		"msgid":       "success sign up",
		"msgdata":     nil,
		"data":        user,
		"data_size":   1,
	})
}
