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

func Test_UcLogin(t *testing.T) {
	now := time.Now()

	tokenProvider := jwt.NewJWTProvider("random-key")

	mockRepo := &mocks.UserRepoMock{
		MockFindUserByEmail: func(ctx context.Context, email string) (*model.User, error) {
			user := &model.User{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, Name: "Joe Doe", Email: "joed@mail.com", Password: "1234", Role: model.RoleCustomer}
			if email != user.Email {
				return nil, common.NotExistAccout
			}
			user.HashPassword()
			return user, nil
		},
	}

	uc := NewAuthUC(mockRepo, tokenProvider, 24*10)

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
			t.Errorf("authUC.Login() should return model.User.ID = 2, but got = %d", got.ID)
		}

		if got.Name != got.Name {
			t.Errorf("authUC.Login() should return model.User.Name = %s, but got = %s", got.Name, got.Name)
		}

		if got.Email != got.Email {
			t.Errorf("authUC.Login() should return model.User.Email = %s, but got = %s", got.Email, got.Email)
		}

		if got.FakeID == strconv.Itoa(2) {
			t.Errorf("authUC.Login() should return model.User.FakeID is a encoded id, but got = %v", got.FakeID)
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
