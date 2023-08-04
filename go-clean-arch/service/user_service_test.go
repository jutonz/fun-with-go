package service

import (
	"context"
	"example/clean-arch/model"
	"example/clean-arch/model/mocks"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
  t.Run("Success", func(t *testing.T) {
    uid, _ := uuid.NewRandom()

    mockUser := &model.User{
      UID: uid,
      Email: "me@t.co",
      Name: "Testing Tester",
    }
    mockUserRepo := new(mocks.MockUserRepository)
    us := NewUserService(&USConfig{
      UserRepository: mockUserRepo,
    })
    mockUserRepo.On("FindById", mock.Anything, uid).Return(mockUser, nil)

    ctx := context.TODO()
    u, err := us.Get(ctx, uid)

    assert.NoError(t, err)
    assert.Equal(t, u, mockUser)
    mockUserRepo.AssertExpectations(t)
  })

  t.Run("NotFound", func(t *testing.T) {
    uid, _ := uuid.NewRandom()

    mockUserRepo := new(mocks.MockUserRepository)
    us := NewUserService(&USConfig{
      UserRepository: mockUserRepo,
    })
    mockUserRepo.On("FindById", mock.Anything, uid).Return(nil, fmt.Errorf("something went wrong"))

    ctx := context.TODO()
    u, err := us.Get(ctx, uid)

    assert.Nil(t, u)
    assert.Error(t, err)
    mockUserRepo.AssertExpectations(t)
  })
	// success
	// not found
}
