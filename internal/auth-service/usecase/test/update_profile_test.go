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

func Test_UpdateUser(t *testing.T) {
	now := time.Now()
	tokenprovider := jwt.NewJWTProvider("random-key")
	mockRepo := &mocks.RepoMock{
		MockUpdateUser: func(ctx context.Context, olduser *model.User, upd *model.UserUpdate) (*model.User, error) {
			if upd.Name != nil {
				olduser.Name = *upd.Name
			}
			if upd.Phone != nil {
				olduser.Phone = *upd.Phone
			}
			if upd.Avatar != nil {
				olduser.Avatar = upd.Avatar
			}
			return olduser, nil
		},
		MockFindUserByID: func(ctx context.Context, id int) (*model.User, error) {
			var foundUser model.User
			rows := []model.User{
				{SqlModel: common.SqlModel{ID: 1, CreatedAt: now, UpdatedAt: now}, Name: "Joe Doe", Email: "joed@mail.com", Password: "12345", Role: model.RoleCustomer},
				{SqlModel: common.SqlModel{ID: 2, CreatedAt: now, UpdatedAt: now}, Name: "Alice", Email: "alice@mail.com", Password: "12345", Role: model.RoleCustomer},
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

	t.Run("Valid update profile", func(t *testing.T) {
		newName := "Joe Dijk"
		newPhoneN := "0987654"
		upd := &model.UserUpdate{
			Name:  &newName,
			Phone: &newPhoneN,
			Avatar: &common.Image{
				Url:       "https://www.baeldung.com/wp-content/uploads/sites/4/2021/05/Blank-diagram.svg",
				Width:     123,
				Height:    123,
				Ext:       ".jpg",
				CloudName: "test",
			},
		}
		got, err := uc.UpdateProfile(context.Background(), 1, upd)

		if err != nil {
			log.Fatalln(err)
		}

		if got.Name != newName {
			t.Errorf("authUC.UpdateProfile() should return %s but got = %s", newName, got.Name)
		}

		if got.Phone != newPhoneN {
			t.Errorf("authUC.UpdateProfile() should return %s but got = %s", newPhoneN, got.Phone)
		}

		if got.Avatar == nil {
			t.Error("authUC.UpdateProfile() should return not nil Avatar")
		}

		if got.Password != "" {
			t.Error("authUC.UpdateProfile() should return Password = ''")
		}
	})

	t.Run("Not found user", func(t *testing.T) {
		nonExistId := 5
		newName := "Joe Dijk"
		newPhoneN := "0987654"
		upd := &model.UserUpdate{
			Name:  &newName,
			Phone: &newPhoneN,
			Avatar: &common.Image{
				Url:       "https://www.baeldung.com/wp-content/uploads/sites/4/2021/05/Blank-diagram.svg",
				Width:     123,
				Height:    123,
				Ext:       ".jpg",
				CloudName: "test",
			},
		}
		got, err := uc.UpdateProfile(context.Background(), nonExistId, upd)

		if got != nil && err == nil {
			t.Error("authUC.UpdateProfile() should return NoExistAccount error")
		}
	})
}
