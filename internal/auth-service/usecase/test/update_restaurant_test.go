package usecase

import (
	"common"
	"context"
	"log"
	"testing"
	"time"
	jwt "umai-auth-service/component"
	"umai-auth-service/mocks"
	"umai-auth-service/model"
	"umai-auth-service/usecase"
)

func Test_UpdateRestaurant(t *testing.T) {
	now := time.Now()
	tokenprovider := jwt.NewJWTProvider("random-key")

	mockRepo := &mocks.RepoMock{
		MockUpdateRestaurant: func(ctx context.Context, oldres *model.Restaurant, upd *model.RestaurantUpdate) (*model.Restaurant, error) {
			if upd.RestaurantName != nil {
				oldres.RestaurantName = *upd.RestaurantName
			}
			if upd.Slogan != nil {
				oldres.Slogan = *upd.Slogan
			}
			if upd.OpenHour != nil {
				oldres.OpenHour = upd.OpenHour
			}
			if upd.CloseHour != nil {
				oldres.OpenHour = upd.CloseHour
			}
			if upd.Cover != nil {
				oldres.Cover = upd.Cover
			}
			return oldres, nil
		},
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

	uc := usecase.NewAuthUC(mockRepo, tokenprovider, 24*10)
	ctx := context.WithValue(context.Background(), common.CurrentUser, &model.User{SqlModel: common.SqlModel{ID: 1}})

	t.Run("Valid update", func(t *testing.T) {
		newName := "BeefDamn"
		got, err := uc.UpdateRestaurant(ctx, 1, &model.RestaurantUpdate{RestaurantName: &newName})

		if err != nil {
			log.Fatalln(err)
		}

		if got.RestaurantName != newName {
			t.Errorf("authUC.UpdateRestaurant() should return %s but got = %s", newName, got.RestaurantName)
		}

		if got.Slogan == "" {
			t.Errorf("authUC.UpdateRestaurant() should return '' but got = %s", got.Slogan)

		}
	})

	t.Run("Not found", func(t *testing.T) {
		nonExistId := 5
		newName := "ABCIMC"

		got, err := uc.UpdateRestaurant(ctx, nonExistId, &model.RestaurantUpdate{RestaurantName: &newName})

		if got != nil && err == nil {
			t.Error("authUC.UpdateRestaurant() should return NoExistAccount error")
		}
	})
}
