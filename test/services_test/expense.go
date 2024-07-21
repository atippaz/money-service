package test

import (
	"testing"
)

type test_service struct {
}

func (test_service) TestMyService_DoSomething(t *testing.T) {
	// mockRepo := new(MockRepository)
	// mockRepo.On("SomeMethod").Return("mocked result")

	// service := MyService{repo: mockRepo}
	// result := service.DoSomething()

	// assert.Equal(t, "mocked result", result)
	// mockRepo.AssertExpectations(t)
}
