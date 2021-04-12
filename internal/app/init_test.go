package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/smartystreets/goconvey/convey"
)

func testNew(t *testing.T) (func(), Appl, *MockRepo, *MockAuth, *MockPassword) {
	ctrl := gomock.NewController(t)

	mockRepo := NewMockRepo(ctrl)
	mockAuth := NewMockAuth(ctrl)
	mockPassword := NewMockPassword(ctrl)

	a := New(mockRepo, mockAuth, mockPassword)
	return ctrl.Finish, a, mockRepo, mockAuth, mockPassword
}
