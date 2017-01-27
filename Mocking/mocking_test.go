package mocking

import (
	"testing"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

//go:generate mockery -inpkg -testonly -name=UserStore

func TestMyFunction(t *testing.T) {
	mockedUserStore := MockUserStore{}

	mockedUserStore.On("GetUser", "Cube").Return(&User{Name: "Jakub", Surname: "Martin", Age: 18}, nil)
	mockedUserStore.On("GetUser", "Cube2").Return(nil, errors.Errorf("User not found."))
	mockedUserStore.On("SetUser", "Cube", mock.AnythingOfType("*mocking.User")).Return(errors.Errorf("User already exists."))
	mockedUserStore.On("SetUser", "Cube2", mock.AnythingOfType("*mocking.User")).Return(nil)

	user, err := mockedUserStore.GetUser("Cube")
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("%v", user)
	}
	user, err = mockedUserStore.GetUser("Cube2")
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("%v", user)
	}
	err = mockedUserStore.SetUser("Cube", &User{Name: "Jakub", Surname: "Martin", Age:18})
	if err != nil {
		t.Log(err)
	}
	err = mockedUserStore.SetUser("Cube2", &User{Name: "Jakub", Surname: "Martin2", Age:18})
	if err != nil {
		t.Log(err)
	}
}
