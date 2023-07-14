package handler

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllCategory godoc
// @Summary Get all category
// @Description Get all category
// @Tags category
// @Accept  json
// @Produce  json
// @Success 200 {object} endeus.CategoryResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /category/all [get]
func (h *Handler) GetAllCategory(c echo.Context) error {
	data, err := h.svc.CategoryService.GetAll()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	result := endeus.CategoryResponse{
		Success: true,
		Message: "Success get list category",
		Data:    &data,
	}
	return c.JSON(http.StatusOK, result)
}
