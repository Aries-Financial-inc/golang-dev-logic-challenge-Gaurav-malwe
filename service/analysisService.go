package service

import "github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"

type IAnalysisService interface {
	AnalysisLogic(contracts []model.OptionsContract) model.AnalysisResult
}

func (s *service) AnalysisLogic(contracts []model.OptionsContract) model.AnalysisResult {
	// Your code here
	return model.AnalysisResult{}
}

func calculateXYValues(contracts []model.OptionsContract) []model.GraphPoint {
	// Your code here
	return nil
}

func calculateMaxProfit(contracts []model.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateMaxLoss(contracts []model.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateBreakEvenPoints(contracts []model.OptionsContract) []float64 {
	// Your code here
	return nil
}
