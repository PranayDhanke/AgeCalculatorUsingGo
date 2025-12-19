package handler

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/pranaydhanke/go-user-api/internal/models"
	"github.com/pranaydhanke/go-user-api/internal/repository"
	"github.com/pranaydhanke/go-user-api/internal/service"
)

type UserHandler struct {
	repo *repository.UserRepository
	v    *validator.Validate
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
		v:    validator.New(),
	}
}

/* -------------------- CREATE -------------------- */

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if err := h.v.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.repo.Create(c.Context(), req.Name, req.Dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob,
	})
}

/* -------------------- GET BY ID -------------------- */

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.repo.Get(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob,
		Age:  age,
	})
}

/* -------------------- LIST ALL -------------------- */

func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.repo.List(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := make([]models.UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			Dob:  user.Dob,
			Age:  service.CalculateAge(user.Dob),
		})
	}

	return c.JSON(response)
}

/* -------------------- UPDATE -------------------- */

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if err := h.v.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.repo.Update(
		c.Context(),
		int32(id),
		req.Name,
		req.Dob,
	)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob,
	})
}

/* -------------------- DELETE -------------------- */

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.repo.Delete(c.Context(), int32(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
