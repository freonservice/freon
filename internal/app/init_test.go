package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/powerman/structlog"
	_ "github.com/smartystreets/goconvey/convey"
)

func testNew(t *testing.T) (func(), Appl, *MockRepo, *MockAuth, *MockPassword) {
	ctrl := gomock.NewController(t)

	mockRepo := NewMockRepo(ctrl)
	mockAuth := NewMockAuth(ctrl)
	mockPassword := NewMockPassword(ctrl)
	mockSettingRepo := NewMockSettingRepo(ctrl)
	mockStorage := NewMockStorage(ctrl)
	mockTranslation := NewMockTranslation(ctrl)

	svc := NewSvc(mockRepo, mockAuth, mockPassword, mockSettingRepo, mockTranslation, mockStorage)
	a := New(svc, Config{}, structlog.New())
	return ctrl.Finish, a, mockRepo, mockAuth, mockPassword
}
