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
		return context.JSON(http.StatusBadRequest, []map[string]interface{}{})
	}
	offSet := offset(context)
	limit := limit(context)
	res, _ := s.PaymentService.PaymentHistory(userId, offSet, limit)
	return context.JSON(http.StatusOK, res)
}

func limit(context echo.Context) int64 {
	var limit int64 = 10
	paramLimit := context.QueryParam("limit")
	if paramLimit != "" {
		if v, err := strconv.ParseInt(paramLimit, 10, 64); err == nil {
			if isLimitInRange(v) {
				limit = v
			}
		}
	}
	return limit
}

func offset(context echo.Context) int64 {
	paramOffSet := context.QueryParam("offset")
	var offSet int64 = math.MaxInt64
	if paramOffSet != "" {
		if v, err := strconv.ParseInt(paramOffSet, 10, 64); err == nil {
			if isValidOffset(v) {
				offSet = v
			}
		}
	}
	return offSet
}

func isLimitInRange(limit int64) bool {
	var maxLimit int64 = 100
	var minLimit int64 = 0
	return (limit <= maxLimit && limit > minLimit)
}

func isValidOffset(offset int64) bool {
	return offset > 0
}
