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

func Test_GetRestaurant(t *testing.T) {
	now := time.Now()
	tokenprovider := jwt.NewJWTProvider("random-key")
	mockRepo := &mocks.RepoMock{
		MockFindRestaurantByID: func(ctx context.Context, id int) (*model.Restaurant, error) {
			var foundRes model.Restaurant
			rows := []model.Restaurant{
				{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, UserID: 1, RestaurantName: "Checkdamn", Slogan: "This is slogan"},
				{SqlModel: common.SqlModel{ID: 2, CreatedAt: now, UpdatedAt: now}, UserID: 1, RestaurantName: "Checkdamn", Slogan: "This is slogan"},
			}

			for _, res := range rows {
				if res.ID == id {
					foundRes = res
				}
			}

			if foundRes.ID == 0 {
				return nil, common.NotExistAccount
			}

			return &foundRes, nil
		},
	}

	cMockRepo := &mocks.CacheRepoMock{
		MockGet: func(ctx context.Context, id int) (*model.Restaurant, error) {
			return nil, common.ErrMissCache
		},
		MockSet: func(ctx context.Context, id int, res *model.Restaurant, ttl time.Duration) error {
			return nil
		},
	}

	uc := usecase.NewAuthUC(mockRepo, cMockRepo, tokenprovider, 24*10)

	t.Run("Valid get user data", func(t *testing.T) {
		got, err := uc.GetRestaurant(context.Background(), 1)
		if err != nil && got == nil {
			t.Error("authUC.GetRestaurant() should not return nil value")
		}
	})

	t.Run("Invalid get user data", func(t *testing.T) {
		NotExistId := 10
		got, err := uc.GetRestaurant(context.Background(), NotExistId)
		if err == nil && got != nil {
			t.Error("authUC.GetRestaurant() should return nil value")
		}
	})
}
