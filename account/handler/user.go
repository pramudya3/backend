package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/pramudya3/backend/payment/domain"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uc domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *userHandler {
	return &userHandler{
		uc: userUsecase,
	}
}

func (u *userHandler) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := &domain.Signup{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		newUser := &domain.User{
			Name:     payload.Name,
			Username: payload.Username,
			Email:    payload.Email,
			Password: payload.Password,
		}

		if err := u.uc.Create(c, newUser); err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		c.Status(http.StatusCreated)
	}
}

func (u *userHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := domain.Login{}
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		user, err := u.uc.GetById(c, login.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, domain.ResponseFailed("id not found"))
				return
			}
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}
		if user.Password != login.Password {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed("id/password wrong"))
			return
		}

		ses, err := session.CreateNewSession(c.Request, c.Writer, "", strconv.Itoa(int(login.ID)), nil, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		c.JSON(http.StatusOK, domain.ResponseSuccess(ses.GetAccessToken()))
	}
}

func (u *userHandler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
		userId := sessionContainer.GetUserID()

		if _, err := session.RevokeAllSessionsForUser(userId, nil, nil); err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		err := supertokens.DeleteUser(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}
		c.Status(http.StatusOK)
	}
}

func (u *userHandler) FindByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
		userId := sessionContainer.GetUserID()
		id, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		user, err := u.uc.GetById(c, uint64(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		c.JSON(http.StatusOK, domain.ResponseSuccess(user))
	}
}
