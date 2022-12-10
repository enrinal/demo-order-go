package delivery

import (
	"github.com/enrinal/demo-order-go/entity"
	"github.com/enrinal/demo-order-go/users/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewUserHandler(e *echo.Echo, userSvc service.UserService) *UserHandler {
	handler := &UserHandler{userSvc}
	e.POST("/api/v1/users/login", handler.Login)
	e.POST("/api/v1/users/register", handler.Register)
	return handler
}

func (h *UserHandler) Login(c echo.Context) error {

	var req entity.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

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
