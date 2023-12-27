package rest

import (
	"common"
	"menu-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func (h *menuHandler) CreateMenuItemHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// id, err := strconv.Atoi(ctx.Param("id"))
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
		// 	return
		// }

		var mitem model.MenuItem

		if err := ctx.ShouldBind(&mitem); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}
		mitem.RestaurantID = ctx.Value(common.CResId).(int)

		validate := validator.New()
		validate.Struct(mitem)
		//call biz
		newMItem, err := h.menuUC.CreateMenuItem(ctx, &mitem)
		if err != nil {
			if err == model.ErrInvalidCategory {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}
		ctx.JSON(http.StatusCreated, common.NewHttpSuccessResponse(http.StatusCreated, "Create new menu item successfully.", newMItem))
	}
}
