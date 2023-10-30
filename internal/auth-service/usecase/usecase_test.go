package usecase

import (
	"common"
	"context"
	"strconv"
	"testing"
	"time"
	jwt "umai-auth-service/component"
	"umai-auth-service/mocks"
	"umai-auth-service/model"
)

func Test_UcRegister(t *testing.T) {
	now := time.Now()
	tokenprovider := jwt.NewJWTProvider("ramdom-key")
	mockRepo := &mocks.UserRepoMock{
		MockFindUserByEmail: func(ctx context.Context, email string) (*model.User, error) {
			foundUser := &model.User{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, Name: "Joe Doe", Email: "joed@mail.com", Password: "12345", Role: "customer"}
			if email != foundUser.Email {
				return nil, common.ExistsEmailError
			}
			return foundUser, nil
		},
		MockInsertUser: func(ctx context.Context, user *model.User) (*model.User, error) {
			createdUser := &model.User{SqlModel: common.SqlModel{ID: 2, CreatedAt: now, UpdatedAt: now}, Name: user.Name, Email: user.Email, Password: user.Password}
			return createdUser, nil
		},
	}
	uc := NewAuthUC(mockRepo, tokenprovider)

	t.Run("Valid registration", func(t *testing.T) {
		userInput := &model.User{Name: "Alice", Email: "alice@mail.com", Password: "12345"}

		got, err := uc.Register(context.Background(), userInput)

		if err != nil {
			t.Fatal(err)
		}

		if got.ID != 2 {
			t.Errorf("authUC.Register() should return model.User.ID = 2, but got = %d", got.ID)
		}

		if got.Name != userInput.Name {
			t.Errorf("authUC.Register() should return model.User.Name = %s, but got = %s", userInput.Name, got.Name)
		}

		if got.Email != userInput.Email {
			t.Errorf("authUC.Register() should return model.User.Email = %s, but got = %s", userInput.Email, got.Email)
		}

		if got.FakeID == strconv.Itoa(2) {
			t.Errorf("authUC.Register() should return model.User.FakeID is a encoded id, but got = %v", got.FakeID)
		}

		if got.Password == userInput.Password || got.Password != "" {
			t.Errorf("authUC.Register() should return model.User.Password = '', but got = %s", got.Password)
		}
	})

	t.Run("Existed email registration", func(t *testing.T) {
		userInput := &model.User{Name: "Joe Don", Email: "joed@mail.com", Password: "12345"}

		got, err := uc.Register(context.Background(), userInput)

		if got != nil && err == nil {
			t.Errorf("authUC.Register() should return model.User.Password = '', but got = %s", got.Password)
		}
	})
}