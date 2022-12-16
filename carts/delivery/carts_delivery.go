package delivery

import (
	"net/http"

	"github.com/enrinal/demo-order-go/domain"
	"github.com/enrinal/demo-order-go/entity"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartSvc domain.CartService
}

func NewCartHandler(cartSvc domain.CartService) *CartHandler {
	return &CartHandler{cartSvc: cartSvc}
}

func (c *CartHandler) AddCart(e echo.Context) error {
	var req entity.CartRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	err := c.cartSvc.AddCart(e.Request().Context(), req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, entity.BaseResponse{
		Message: "success",
	})
}

func (c *CartHandler) GetCartById(e echo.Context) error {
	id := e.Param("id")
	res, err := c.cartSvc.GetCartById(e.Request().Context(), id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return e.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}
