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

func NewProductHandler(e *echo.Echo, productService domain.ProductService) *ProductHandler {
	handler := &ProductHandler{ProductService: productService}

	basePath := "/api/v1/products"
	e.GET(basePath, handler.GetAll)
	e.GET(basePath+"/:id", handler.GetById)

	return handler
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
