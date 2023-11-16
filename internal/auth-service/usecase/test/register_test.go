package usecase

import (
	"common"
	"context"
	"testing"
	"time"
	jwt "umai-auth-service/component"
	"umai-auth-service/mocks"
	"umai-auth-service/model"
	"umai-auth-service/usecase"
)

func Test_UcRegister(t *testing.T) {
	now := time.Now()
	tokenprovider := jwt.NewJWTProvider("ramdom-key")
	mockRepo := &mocks.RepoMock{
		MockFindUserByEmail: func(ctx context.Context, email string) (*model.User, error) {
			foundUser := &model.User{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, Name: "Joe Doe", Email: "joed@mail.com", Password: "12345", Role: model.RoleCustomer}
			if email != foundUser.Email {
				return nil, common.NotExistAccount
			}
			return foundUser, nil
		},
		MockInsertUser: func(ctx context.Context, user *model.User) (int, error) {
			createdUser := &model.User{SqlModel: common.SqlModel{ID: 2, CreatedAt: now, UpdatedAt: now}, Name: user.Name, Email: user.Email, Password: user.Password, Role: user.Role}
			return createdUser.ID, nil
		},
	}
	uc := usecase.NewAuthUC(mockRepo, tokenprovider, 24*10)

	t.Run("Valid registration", func(t *testing.T) {
		userInput := &model.User{Name: "Alice", Email: "alice@mail.com", Password: "12345"}

		got, err := uc.Register(context.Background(), userInput)

		if err != nil {
			t.Fatal(err)
		}

		if got != 2 {
			t.Errorf("authUC.Register() should return model.User.ID = 2, but got = %d", got)
		}
	})

	t.Run("Existed email registration", func(t *testing.T) {
		userInput := &model.User{Name: "Joe Don", Email: "joed@mail.com", Password: "12345"}

		got, err := uc.Register(context.Background(), userInput)

		if got != 0 && err == nil {
			t.Errorf("authUC.Register() should return error existed email")
		}
	})
}
