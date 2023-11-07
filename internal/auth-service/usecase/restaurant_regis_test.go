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
)

func Test_RetaurantRes(t *testing.T) {
	now := time.Now()
	tokenprovider := jwt.NewJWTProvider("random-key")

	mockRepo := &mocks.RepoMock{
		MockFindUserByID: func(ctx context.Context, id int) (*model.User, error) {
			var foundUser model.User
			rows := []model.User{
				{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, Name: "Joe Doe", Email: "joed@mail.com", Password: "12345", Role: model.RoleCustomer},
				{SqlModel: common.SqlModel{ID: 3, CreatedAt: now, UpdatedAt: now}, Name: "Alice", Email: "alice@mail.com", Password: "12345", Role: model.RoleRestaurant},
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
		MockUpdateRole: func(ctx context.Context, user *model.User) error {
			user.Role = model.RoleRestaurant
			user.UpdatedAt = now
			if user.Role != model.RoleRestaurant {
				return common.InternalServerError
			}
			return nil
		},
		MockInsertRestaurant: func(ctx context.Context, res *model.Restaurant) (*model.Restaurant, error) {
			restaurant := &model.Restaurant{SqlModel: common.SqlModel{ID: 2, CreatedAt: now, UpdatedAt: now}, UserID: res.UserID, RestaurantName: res.RestaurantName}
			return restaurant, nil
		},
	}

	uc := NewAuthUC(mockRepo, tokenprovider, 24*10)

	t.Run("Valid restaurant registration", func(t *testing.T) {
		resInput := &model.Restaurant{RestaurantName: "JChick", UserID: 1}

		got, err := uc.RestaurantRegis(context.Background(), resInput)

		if err != nil {
			log.Fatal(err)
		}

		if got.ID != 2 {
			t.Errorf("authUC.RestaurantRegis() should return model.Restaurant.ID = 2 but got %d", got.ID)
		}

		if got.UserID != resInput.UserID {
			t.Errorf("authUC.RestaurantRegis() should return model.Restaurant.UserID = 1 but got %d", got.UserID)
		}

		if got.RestaurantName != got.RestaurantName {
			t.Errorf("authUC.RestaurantRegis() should return model.Restaurant.RestaurantName = %s but got %s", got.RestaurantName, got.RestaurantName)
		}
	})

	t.Run("Invalid user role restaurant registration", func(t *testing.T) {
		resInput := &model.Restaurant{RestaurantName: "JChick", UserID: 3}

		got, err := uc.RestaurantRegis(context.Background(), resInput)

		if got != nil && err != common.BadRequest {
			t.Errorf("authUC.RestaurantRegis() should return error: %s", common.BadRequest.Error())
		}
	})

	t.Run("Not found user restaurant registration", func(t *testing.T) {
		resInput := &model.Restaurant{RestaurantName: "JChick", UserID: 10}
		got, err := uc.RestaurantRegis(context.Background(), resInput)

		if got != nil && err != common.NotExistAccount {
			t.Errorf("authUC.RestaurantRegis() should return error: %s", common.NotExistAccount.Error())
		}
	})

}
