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

func Test_GetProfile(t *testing.T) {
	now := time.Now()
	tokenprovider := jwt.NewJWTProvider("random-key")
	mockRepo := &mocks.RepoMock{
		MockFindUserByID: func(ctx context.Context, id int) (*model.User, error) {
			var foundUser model.User
			rows := []model.User{
				{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, Name: "Joe Doe", Email: "joed@mail.com", Password: "12345", Role: model.RoleCustomer},
				{SqlModel: common.SqlModel{ID: 2, CreatedAt: now, UpdatedAt: now}, Name: "Alice", Email: "alice@mail.com", Password: "12345", Role: model.RoleRestaurant},
			}

			for _, user := range rows {
				if user.ID == id {
					foundUser = user
				}
			}

			if foundUser.ID == 0 {
				return nil, common.NotExistAccount
			}

			return &foundUser, nil
		},
	}

	cMockRepo := &mocks.CacheRepoMock{}

	uc := usecase.NewAuthUC(mockRepo, cMockRepo, tokenprovider, 24*10)

	t.Run("Valid get user data", func(t *testing.T) {
		got, err := uc.GetProfile(context.Background(), 1)
		if err != nil && got == nil {
			t.Error("authUC.GetProfile() should not return nil value")
		}
	})

	t.Run("Invalid get user data", func(t *testing.T) {
		NotExistId := 10
		got, err := uc.GetProfile(context.Background(), NotExistId)
		if err == nil && got != nil {
			t.Error("authUC.GetProfile() should return nil value")
		}
	})
}
