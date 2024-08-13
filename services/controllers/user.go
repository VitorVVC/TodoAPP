package controllers

import (
	"api-postgresql/models"
	"api-postgresql/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	validator *validator.Validate
}

func NewUser() *UserController {
	return &UserController{}
}

func (u UserController) Create(ctx echo.Context) error {
	var data models.User

	if err := ctx.Bind(&data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse body")
	}

	if err := u.validator.Struct(data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to validate body")
	}

	return utils.HTTPCreated(ctx, data)
}
