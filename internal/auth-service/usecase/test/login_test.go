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

func Test_UcLogin(t *testing.T) {
	now := time.Now()

	tokenprovider := jwt.NewJWTProvider("random-key")

	mockRepo := &mocks.RepoMock{
		MockFindUserByEmail: func(ctx context.Context, email string) (*model.User, error) {
			user := &model.User{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, Name: "Joe Doe", Email: "joed@mail.com", Password: "1234", Role: model.RoleCustomer}
			if email != user.Email {
				return nil, common.NotExistAccount
			}
			user.HashPassword()
			return user, nil
		},
	}

	uc := usecase.NewAuthUC(mockRepo, tokenprovider, 24*10)

	t.Run("Valid login", func(t *testing.T) {
		cred := &model.LoginCredentials{
			Email:    "joed@mail.com",
			Password: "1234",
		}

		got, err := uc.Login(context.Background(), cred)

		if err != nil {
			t.Fatal(err)
		}
		if got.ID != 1 {
			t.Errorf("authUC.Login() should return model.User.ID = 1, but got = %d", got.ID)
		}

		if got.Name != got.Name {
			t.Errorf("authUC.Login() should return model.User.Name = %s, but got = %s", got.Name, got.Name)
		}

		if got.Email != got.Email {
			t.Errorf("authUC.Login() should return model.User.Email = %s, but got = %s", got.Email, got.Email)
		}

		if got.Password != "" {
			t.Errorf("authUC.Login() should return model.User.Password = '', but got = %s", got.Password)
		}
	})

	t.Run("Wrong password login", func(t *testing.T) {
		cred := &model.LoginCredentials{
			Email:    "joed@mail.com",
			Password: "abcd",
		}

		got, err := uc.Login(context.Background(), cred)

		if err == nil && err != common.WrongPassword && got != nil {
			t.Error("authUC.Login() should return error wrong password")
		}
	})
}
