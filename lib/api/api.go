package api

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rover10/payment/lib/constant"
)

// TransactionHistory fetch recent transactions of the user
func (s *Server) TransactionHistory(context echo.Context) error {
	// Assuming that the userId is validated & belongs to loggedin user
	ID := context.Param("userId")
	log.Printf("TransactionHistory: Getting payment history for userId: %v", ID)
	userId, err := strconv.Atoi(ID)
	errResponse := []map[string]interface{}{}
	if err != nil {
		log.Printf("TransactionHistory: error while converting userId: err: %v, userId: %v", err, ID)
		return context.JSON(http.StatusBadRequest, errResponse)
	}
	// offset & limit to control the response size
	// limit default value is 10
	offSet := offset(context)
	limit := limit(context)
	res, err := s.PaymentService.PaymentHistory(userId, offSet, limit)
	if err != nil {
		log.Printf("TransactionHistory: error getting payment history: err: %v, userId: %v", err, userId)
		context.JSON(http.StatusInternalServerError, errResponse)
	}
	return context.JSON(http.StatusOK, res)
}

func limit(context echo.Context) int64 {
	var limit int64 = constant.DEFAULTLIMIT
	paramLimit := context.QueryParam("limit")
	if paramLimit != "" {
		v, err := strconv.ParseInt(paramLimit, 10, 64)
		if err == nil && isLimitInRange(v) {
			limit = v
		}
	}
	return limit
}

func offset(context echo.Context) int64 {
	paramOffSet := context.QueryParam("offset")
	var offSet int64 = math.MaxInt64
	if paramOffSet != "" {
		v, err := strconv.ParseInt(paramOffSet, 10, 64)
		if err == nil && isValidOffset(v) {
			offSet = v
		}
	}
	return offSet
}

func isLimitInRange(limit int64) bool {
	return (limit <= constant.MAXLIMIT && limit >= constant.MINLIMIT)
}

func isValidOffset(offset int64) bool {
	return offset > 0
}
