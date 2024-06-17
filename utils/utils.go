package utils

import (
	"math"
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/constants"
)

var (
	RR1001 = setErrorFields(http.StatusBadRequest, constants.RR1001, "BAD REQUEST")
)

func setErrorFields(httpStatus int, code string, errMessage string) map[string]interface{} {
	return map[string]interface{}{
		"HTTP_STATUS": httpStatus,
		"CODE":        code,
		"ERR_MESSAGE": errMessage,
	}
}

func CustomErrorFields(setErrorFields map[string]interface{}, customMessage string) map[string]interface{} {
	setErrorFields["ERR_MESSAGE"] = customMessage
	return setErrorFields
}

func GetEnvOrDefaultFloat64(key string, defaultValue float64) float64 {
	cfg := config.GetConfig()
	value := cfg.GetFloat64(key)
	if value == 0 {
		value = defaultValue
	}
	return value
}

func RoundToDecimal(num float64, decimalPlaces int) float64 {
	precision := math.Pow10(decimalPlaces)
	return math.Round(num*precision) / precision
}
