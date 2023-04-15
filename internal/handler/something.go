package handler

import (
	"github.com/labstack/echo/v4"
	"hackaton/internal/models"
	"hackaton/pkg/data"
	"net/http"
)

func (h *handler) HandleSomething(c echo.Context) error {
	var req data.DoSomethingRequest
	err := c.Bind(&req)
	if err != nil {
		return handleError(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	resp, err := h.service.SomeService.DoSomething(ctx, req)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func handleError(c echo.Context, code int, err error) error {
	return c.JSON(code, models.ErrorMessage{
		Code:    code,
		Message: err.Error(),
	})
}
