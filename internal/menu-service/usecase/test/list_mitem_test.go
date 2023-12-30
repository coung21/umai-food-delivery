package usecase

// func Test_ListMitemByResID(t *testing.T) {
// 	mockRepo := &mocks.RepoMock{
// 		MockListMenuItemByResID: func(ctx context.Context, rid int) ([]model.MenuItem, error) {
// 			foundMenuItems := make([]model.MenuItem, 0)
// 			menuItems := []model.MenuItem{
// 				{ID: primitive.NewObjectID(), RestaurantID: 1, Name: "Chicken"},
// 				{ID: primitive.NewObjectID(), RestaurantID: 1, Name: "Milk Tea"},
// 				{ID: primitive.NewObjectID(), RestaurantID: 2, Name: "Tea"},
// 			}
// 			for _, item := range menuItems {
// 				if item.RestaurantID == rid {
// 					foundMenuItems = append(foundMenuItems, item)
// 				}
// 			}
// 			return foundMenuItems, nil
// 		},
// 	}

// 	mockCacheRepo := &mocks.CacheRepoMock{}

// 	uc := usecase.NewMenuUC(mockRepo, mockCacheRepo)

// 	t.Run("Found restaurant id", func(t *testing.T) {
// 		got, err := uc.ListMenuItemByResID(context.Background(), 1)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if len(*got) != 2 {
// 			t.Errorf("uc.ListMenuItemByResID() should return data len == 2 but got %d", len(*got))
// 		}
// 	})
// 	t.Run("Not found restaurant id", func(t *testing.T) {
// 		notExID := 10
// 		got, err := uc.ListMenuItemByResID(context.Background(), notExID)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if len(*got) != 0 {
// 			t.Errorf("uc.ListMenuItemByResID() should return data len == 0 but got %d", len(*got))
// 		}
// 	})
// }
