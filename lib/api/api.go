package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) TransactionHistory(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello")
}
