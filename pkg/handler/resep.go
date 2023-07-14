package handler

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllResep godoc
// @Summary Get all resep
// @Description Get all resep
// @Tags resep
// @Accept  json
// @Produce  json
// @Param search query string false "Filter by search"
// @Success 200 {object} endeus.ResepAllResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /resep/all [get]
func (h *Handler) GetAllResep(c echo.Context) error {
	param := c.QueryParam("search")

	data, err := h.svc.ResepService.GetAll(param)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	result := endeus.ResepAllResponse{
		Success: true,
		Message: "Success get all resep",
		Data:    &data,
	}
	return c.JSON(http.StatusOK, result)
}

// GetAllResepByCategoryID godoc
// @Summary Get all resep
// @Description Get all resep
// @Tags resep
// @Accept  json
// @Produce  json
// @Param id path string true "id from category"
// @Success 200 {object} endeus.ResepAllResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /resep/category/{id} [get]
func (h *Handler) GetAllResepByCategoryID(c echo.Context) error {
	id, err := utils.ConvertID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	data, err := h.svc.GetAllByCategoryID(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	result := endeus.ResepAllResponse{
		Success: true,
		Message: "Success get list resep",
		Data:    &data,
	}
	return c.JSON(http.StatusOK, result)
}

// GetAllResepByID godoc
// @Summary Get one resep
// @Description Get one resep
// @Tags resep
// @Accept  json
// @Produce  json
// @Param id path string true "id from resep"
// @Success 200 {object} endeus.ResepSingleResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /resep/single/{id} [get]
func (h *Handler) GetResepByID(c echo.Context) error {
	id, err := utils.ConvertID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	data, err := h.svc.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	result := endeus.ResepSingleResponse{
		Success: true,
		Message: "Success get resep",
		Data:    &data,
	}
	return c.JSON(http.StatusOK, result)
}

// CreateResep godoc
// @Summary Create one resep
// @Description Create one resep
// @Tags resep
// @Accept  json
// @Produce  json
// @Param resep body endeus.ResepParam true "Resep to create"
// @Success 200 {object} utils.Success
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /resep/create [post]
func (h *Handler) CreateResep(c echo.Context) error {
	var param endeus.ResepParam
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := c.Validate(param); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewValidatorError(err))
	}

	err := h.svc.CreateNew(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessMsg("Success create resep"))
}

// UnpublishResep godoc
// @Summary Change publis status resep
// @Description Change publis status resep
// @Tags resep
// @Accept  json
// @Produce  json
// @Param id path string true "id from resep"
// @Success 200 {object} utils.Success
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /resep/unpublish/{id} [put]
func (h *Handler) UnpublishResep(c echo.Context) error {
	id, err := utils.ConvertID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.svc.Unpublish(id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessMsg("Success unpublish resep"))
}

// DeleteResep godoc
// @Summary Delete single resep
// @Description Delete single resep
// @Tags resep
// @Accept  json
// @Produce  json
// @Param id path string true "id from resep"
// @Success 200 {object} utils.Success
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /resep/delete/{id} [delete]
func (h *Handler) DeleteResep(c echo.Context) error {
	id, err := utils.ConvertID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.svc.Delete(id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessMsg("Success delete resep"))
}

// UpdateResep godoc
// @Summary Update one resep
// @Description Update one resep
// @Tags resep
// @Accept  json
// @Produce  json
// @Param id path string true "id from resep"
// @Param resep body endeus.ResepParam true "Resep to update"
// @Success 200 {object} utils.Success
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /resep/update/{id} [put]
func (h *Handler) UpdateResep(c echo.Context) error {
	var param endeus.ResepParam
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := c.Validate(param); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewValidatorError(err))
	}

	id, err := utils.ConvertID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	err = h.svc.Update(id, param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessMsg("Success update resep"))
}
