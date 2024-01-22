package handler

import (
	"Homework2/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserService interface {
	GetAll() []*model.User
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) InitRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAll)
}

// GetAll godoc
// @Summary 	Get Users
// @Description Get list of users
// @Tags 		users
// @Produce 	json
// @Success 	200 {object} 		[]model.User
// @Router 		/users [get]
func (handler *UserHandler) GetAll(ctx *fiber.Ctx) error {
	users := handler.service.GetAll()

	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"users": users,
		})
}
