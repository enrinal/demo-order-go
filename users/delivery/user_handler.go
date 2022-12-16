package delivery

import (
	"net/http"

	"github.com/enrinal/demo-order-go/domain"

	"github.com/enrinal/demo-order-go/entity"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userSvc domain.UserService
}

func NewUserHandler(userSvc domain.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

func (h *UserHandler) Login(c echo.Context) error {
	// decode request
	var req entity.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	// call service
	res, err := h.userSvc.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (h *UserHandler) Register(c echo.Context) error {
	var req entity.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	err := h.userSvc.Register(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
	})
}
