package handler

import (
	"github.com/Brainsoft-Raxat/hacknu/pkg/data"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) CheckIIN(c echo.Context) error {
	var req data.CheckIINRequest
	err := c.Bind(&req)
	if err != nil {
		return handleError(c, http.StatusBadRequest, err)
	}

	req.IIN = c.Param("iin")

	ctx := c.Request().Context()
	resp, err := h.service.OrderService.CheckIIN(ctx, req)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetClientData(c echo.Context) error {
	var req data.GetClientDataRequest
	err := c.Bind(&req)
	if err != nil {
		return handleError(c, http.StatusBadRequest, err)
	}

	req.IIN = c.Param("iin")

	ctx := c.Request().Context()
	resp, err := h.service.OrderService.GetClientData(ctx, req)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetBranches(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.OrderService.GetDeliveryServices(ctx)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) CreateOrder(c echo.Context) error {
	var req data.CreateOrderRequest
	err := c.Bind(&req)
	if err != nil {
		return handleError(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	resp, err := h.service.OrderService.CreateOrder(ctx, req)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetCoordinates(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.OrderService.GetCoordinates(ctx, data.GetCoordinatesRequest{Street: "Kazakhstan, Astana, Kenesary 9"})
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}
