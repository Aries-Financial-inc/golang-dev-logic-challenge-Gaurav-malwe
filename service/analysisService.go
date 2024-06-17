package service

import (
	"context"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"
)

type IAnalysisService interface {
	AnalysisLogic(ctx context.Context, contracts []model.OptionsContract) model.AnalysisResult
}

func (s *service) AnalysisLogic(ctx context.Context, contracts []model.OptionsContract) model.AnalysisResult {
	// Your code here
	return model.AnalysisResult{}
}

func calculateXYValues(ctx context.Context, contracts []model.OptionsContract) []model.GraphPoint {
	// Your code here
	return nil
}

func calculateMaxProfit(ctx context.Context, contracts []model.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateMaxLoss(ctx context.Context, contracts []model.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateBreakEvenPoints(ctx context.Context, contracts []model.OptionsContract) []float64 {
	// Your code here
	return nil
}
