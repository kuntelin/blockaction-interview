package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UsersController struct{}

// func ListUserController() UsersController {
// 	return UsersController{}
// }

// func (u UsersController) ListUser(c *gin.Context) {
// 	fmt.Println("ListUserController")
// 	users := ListUserService()

//		c.JSON(http.StatusOK, gin.H{
//			"return_code": 0,
//			"msgid":       "success update data",
//			"msgdata":     nil,
//			"data":        users,
//			"data_size":   len(users),
//		})
//	}

/* list user */
func ListUserController(c *gin.Context) {
	logger.Debug("ListUserController")

	users := ListUserService()

	c.JSON(http.StatusOK, gin.H{
		"return_code": 0,
		"msgid":       "success update data",
		"msgdata":     nil,
		"data":        users,
		"data_size":   len(users),
	})
}

/* create user*/
type CreateUserForm struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
	Email    string `json:"email" binding:"required" example:"test123@gmail.com"`
}

func CreateUserController(c *gin.Context) {
	logger.Debug("CreateUserController")

	var createUserForm CreateUserForm

	logger.Debug("createUserForm", createUserForm)

	bindErr := c.BindJSON(&createUserForm)

	logger.Error("bindErr", bindErr)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "failed to create user with username %s, email %s",
			"msgdata": gin.H{
				"username": createUserForm.Username,
				"email":    createUserForm.Email,
			},
			"trace": bindErr.Error(),
		})
		return
	}

	user, createErr := CreateUserService(createUserForm.Username, createUserForm.Password, createUserForm.Email)
	if createErr != nil {
		logger.Debug("createErr: ", createErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "failed to create user with username %s, email %s",
			"msgdata": gin.H{
				"username": createUserForm.Username,
				"email":    createUserForm.Email,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"return_code": 0,
		"msgid":       "success Register",
		"msgdata":     nil,
		"data":        user,
		"data_size":   1,
	})
}

/* get user */
func GetUserController(c *gin.Context) {
	logger.Debug("GetUserController")

	username := c.Params.ByName("username")

	user := GetUserService(&username)

	if user.Username != "" {
		c.JSON(http.StatusOK, gin.H{
			"return_code": 0,
			"msgid":       "success get user by username",
			"msgdata":     nil,
			"data":        user,
			"data_size":   1,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"return_code": 1002,
		"msgid":       "failed to get user by username, username %s",
		"msgdata": gin.H{
			"username": username,
		},
		"trace": "user not found",
	})
}
