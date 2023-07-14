package handler

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	config endeus.Config
	svc    *service.Service
}

func NewHandler(cfg endeus.Config, service *service.Service) *Handler {
	return &Handler{
		config: cfg,
		svc:    service,
	}
}

func (h *Handler) Register(route *echo.Group) {
	route.GET("/category/all", h.GetAllCategory)

	resepRoute := route.Group("/resep")
	resepRoute.GET("/all", h.GetAllResep)
	resepRoute.GET("/category/:id", h.GetAllResepByCategoryID)
	resepRoute.GET("/single/:id", h.GetResepByID)
	resepRoute.POST("/create", h.CreateResep)
	resepRoute.PUT("/unpublish/:id", h.UnpublishResep)
	resepRoute.DELETE("/delete/:id", h.DeleteResep)
	resepRoute.PUT("/update/:id", h.UpdateResep)

}
