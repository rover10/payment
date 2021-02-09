package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (s *Server) TransactionHistory(context echo.Context) error {
	ID := context.Param("userId")
	userId, err := strconv.Atoi(ID)
	if err != nil {
		// Empty array
		return context.JSON(http.StatusBadRequest, []map[string]interface{}{})
	}

	res, _ := s.Database.PaymentHistory(userId)
	return context.JSON(http.StatusOK, res)
}
