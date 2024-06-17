package service

import "github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"

// This is the service interface
type (
	Service interface {
		IAnalysisService
	}

	service struct {
		config *config.Config
	}
)

func New(config *config.Config) Service {
	return &service{
		config: config,
	}
}
