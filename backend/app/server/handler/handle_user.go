package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/raveger/NuxtTodoApp/backend/domain/service"
	"github.com/pkg/errors"
)

type UserHandler interface {
	Users(c echo.Context) error
}

type UserHandler struct {
	user service.User
}

func NewUserHandler(user service.User) UserHandler {
	return &userHandler(user: user)
}

func (u *userHandler) Users(c echo.Context) error {
	paramID := c.QueryParam("id")
	if paramID == "" {
		users, err := u.user.Users()
		if err != nil {
			err = errors.Wrap(err, "cannot get users")
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		}
		return c.JSON(http.StatusOK, users)
	}
	id, err := strconv.Atoi(paramID)
	if err != nil {
		err = errors.Wrap(ErrUserIDIsNotNumber, "Invalid user id")
		log.Println(err)
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	user, err := u.user.User(id)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("cannot get user %d", id))
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, user)
}