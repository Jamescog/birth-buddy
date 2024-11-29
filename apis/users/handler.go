package users

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Jamescog/birth-buddy/ent"
	"github.com/Jamescog/birth-buddy/ent/users"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *ent.Client) {
	userGroup := rg.Group("/users")
	userGroup.GET("/", func(c *gin.Context) { GetAllUsers(c, db) })
	userGroup.GET("/:telegram_id", func(c *gin.Context) { GetSingleUser(c, db) })
	userGroup.POST("/register", func(c *gin.Context) { RegisterNewUser(c, db) })
}

// UsersDetailDoc godoc

// @Summary		Return all users in the system
// @Description	Returns all users
// @Tags			User
// @Produce		json
// @Success		200	{object}	[]UserDetailOutDTO
// @Failure		500
// @Router			/users/	[get]
func GetAllUsers(c *gin.Context, db *ent.Client) {
	var usersDTO []UserDetailOutDTO
	users, err := db.Users.Query().All(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for _, user := range users {
		usersDTO = append(usersDTO, UserDetailOutDTO{
			TelegramID:    user.TelegramID,
			Username:      user.Username,
			FullName:      user.FullName,
			BirthdayCount: user.BirthdayCount,
			IsPremium:     user.IsPremium,
		})

	}
	c.JSON(200, usersDTO)
}

// UsesrRegistrationDoc godoc

// @Summary		Create a user
// @Description	Create  a new telegram user in the system
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body	UserRegistrationRequestDTO	true	"User Data"
// @Success		201
// @Failure		400
// @Router			/users/register	[post]
func RegisterNewUser(c *gin.Context, db *ent.Client) {
	var user UserRegistrationRequestDTO

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	usr, err := db.Users.Create().
		SetFullName(user.FullName).
		SetNillableUsername(&user.Username).
		SetTelegramID(user.TelegramID).Save(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"user": usr})
}

// SingleUserDetail godoc

// @Summary		Returns single user
// @Description	Return single user by telegram_id
// @Tags			User
// @Produce		json
// @Param			telegram_id	path		int	true	"User Telegram ID"
// @Success		200			{object}	UserDetailOutDTO
// @Failure		400
// @Failure		404
// @Faulure		500
//
// @Router			/users/{telegram_id} [get]
func GetSingleUser(c *gin.Context, db *ent.Client) {
	var u UserDetailOutDTO

	telegram_id := c.Param("telegram_id")

	i, err := strconv.Atoi(telegram_id)

	if err != nil {

		c.JSON(400, gin.H{"error": "Ivalid telegram id type Expected int type"})
		return

	}

	user, err := db.Users.Query().
		Where(users.TelegramID(i)).
		Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		return

	}

	u = UserDetailOutDTO{
		TelegramID:    user.TelegramID,
		Username:      user.Username,
		FullName:      user.FullName,
		BirthdayCount: user.BirthdayCount,
		IsPremium:     user.IsPremium,
	}

	c.JSON(200, u)

}
