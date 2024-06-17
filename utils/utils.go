package utils

import (
	"net/http"

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
