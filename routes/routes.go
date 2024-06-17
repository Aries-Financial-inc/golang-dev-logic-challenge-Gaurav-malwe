package routes

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/controllers"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/middleware"
	"github.com/gin-gonic/gin"
)

type RiskRewardServer struct {
	Router *gin.Engine
}

func NewRiskRewardServer() *RiskRewardServer {
	server := new(RiskRewardServer)
	server.Router = gin.Default()
	return server
}

func (s *RiskRewardServer) SetupRouter(c controllers.Controller, middlewares ...gin.HandlerFunc) {
	group := s.Router.Group("/",
		middleware.TransactionInMiddleware(), // transaction middleware
		middleware.RequestLogger(),           // logger middleware
	) // group for all routes and middlewares. Here the global middlewares are applied for all route groups
	group.Use(middlewares...)

	group.POST("/analyze", c.AnalysisHandler)

}
