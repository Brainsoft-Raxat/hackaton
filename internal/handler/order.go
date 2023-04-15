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
