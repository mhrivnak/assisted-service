// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package connectivity is a generated GoMock package.
package connectivity

import (
	reflect "reflect"

	common "github.com/filanov/bm-inventory/internal/common"
	"github.com/filanov/bm-inventory/internal/validators"
	models "github.com/filanov/bm-inventory/models"
	gomock "github.com/golang/mock/gomock"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// IsSufficient mocks base method
func (m *MockValidator) IsSufficient(host *models.Host, cluster *common.Cluster) (*validators.IsSufficientReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSufficient", host, cluster)
	ret0, _ := ret[0].(*validators.IsSufficientReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSufficient indicates an expected call of IsSufficient
func (mr *MockValidatorMockRecorder) IsSufficient(host, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSufficient", reflect.TypeOf((*MockValidator)(nil).IsSufficient), host, cluster)
}

// GetHostValidInterfaces mocks base method
func (m *MockValidator) GetHostValidInterfaces(host *models.Host) ([]*models.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostValidInterfaces", host)
	ret0, _ := ret[0].([]*models.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostValidInterfaces indicates an expected call of GetHostValidInterfaces
func (mr *MockValidatorMockRecorder) GetHostValidInterfaces(host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostValidInterfaces", reflect.TypeOf((*MockValidator)(nil).GetHostValidInterfaces), host)
}
