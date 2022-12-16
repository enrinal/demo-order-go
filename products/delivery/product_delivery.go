package delivery

import (
	"net/http"

	"github.com/enrinal/demo-order-go/domain"
	"github.com/enrinal/demo-order-go/entity"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductService domain.ProductService
}

func NewProductHandler(productService domain.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	res, err := h.ProductService.GetAll(c.Request().Context())
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

func (h *ProductHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	res, err := h.ProductService.GetById(c.Request().Context(), id)
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
