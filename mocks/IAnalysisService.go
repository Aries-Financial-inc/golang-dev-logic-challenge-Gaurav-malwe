// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"
	mock "github.com/stretchr/testify/mock"
)

// IAnalysisService is an autogenerated mock type for the IAnalysisService type
type IAnalysisService struct {
	mock.Mock
}

// AnalysisLogic provides a mock function with given fields: ctx, contracts
func (_m *IAnalysisService) AnalysisLogic(ctx context.Context, contracts []model.OptionsContract) model.AnalysisResult {
	ret := _m.Called(ctx, contracts)

	if len(ret) == 0 {
		panic("no return value specified for AnalysisLogic")
	}

	var r0 model.AnalysisResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.OptionsContract) model.AnalysisResult); ok {
		r0 = rf(ctx, contracts)
	} else {
		r0 = ret.Get(0).(model.AnalysisResult)
	}

	return r0
}

// NewIAnalysisService creates a new instance of IAnalysisService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAnalysisService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAnalysisService {
	mock := &IAnalysisService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
