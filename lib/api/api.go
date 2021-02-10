package api

import (
	"math"
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

	// offset
	paramOffSet := context.QueryParam("offset")
	var offSet int64 = math.MaxInt64
	if paramOffSet != "" {
		if v, err := strconv.ParseInt(paramOffSet, 10, 64); err == nil && v > 0 {
			offSet = v
		}
	}

	// limit
	var limit int64 = 10
	var maxAllowed int64 = 100
	paramlimit := context.QueryParam("limit")
	if paramlimit != "" {
		if v, err := strconv.ParseInt(paramlimit, 10, 64); err == nil && v <= maxAllowed && v > 0 {
			limit = v
		}
	}

	res, _ := s.Database.PaymentHistory(userId, offSet, limit)
	return context.JSON(http.StatusOK, res)
}
