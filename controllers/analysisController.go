package controllers

import (
	"fmt"
	"strings"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/log"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
)

type IAnalysisController interface {
	AnalysisHandler(ginCtx *gin.Context)
}

func (c *controller) AnalysisHandler(ginCtx *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(ginCtx.Request.Context(), "IAnalysisController::AnalysisHandler")
	defer span.Finish()

	log := log.GetLogger(ctx)
	log.Info("Controller::Customer::UploadCustomerFromCSV")

	payload, err := validateModelRequest(ginCtx)
	if checkError(ginCtx, err) {
		return
	}

	// func(c *gin.Context) {
	// 	var contracts []model.OptionsContract

	// 	if err := c.ShouldBindJSON(&contracts); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	// Your code here

	// 	c.JSON(http.StatusOK, gin.H{"message": "Your code here"})
	// }
}

func validateModelRequest(ginCtx *gin.Context) ([]model.OptionsContract, error) {
	var payload []model.OptionsContract
	var err error

	// check binding
	if err := ginCtx.ShouldBind(&payload); err != nil {
		return payload, err
	}

	validate := validator.New()

	err = validate.Struct(payload)
	if err != nil {
		arr := listErrors(err)
		return payload, fmt.Errorf("%#v", utils.CustomErrorFields(utils.RR1001, ("Invalid/missing input parameters: "+arr)))
	}
	return payload, nil
}

func listErrors(err error) string {
	var arr []string
	for _, err := range err.(validator.ValidationErrors) {
		arr = append(arr, err.Field())
	}
	str := strings.Join(arr, ", ")
	return str
}
