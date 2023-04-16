package handler

import (
	"github.com/Brainsoft-Raxat/hacknu/internal/service"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service.Service
}

type Handler interface {
	Register(e *echo.Echo)
}

func New(services *service.Service) Handler {
	return &handler{service: services}
}

func (h *handler) Register(e *echo.Echo) {
	e.Use()
	api := e.Group("/api")
	{
		api.GET("/check/:iin", h.CheckIIN)
		api.GET("/client/:iin", h.GetClientData)
		api.GET("/branches", h.GetBranches)
		api.POST("/document", h.DocumentReady)
		api.POST("/coordinates", h.GetCoordinates)
		api.POST("/orders/create", h.CreateOrder)
		api.POST("/orders/confirm", h.ConfirmOrder)
		api.GET("/orders", h.GetOrders)
	}
}
