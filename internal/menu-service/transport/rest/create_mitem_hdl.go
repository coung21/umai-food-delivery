package rest

import (
	"common"
	"menu-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *menuHandler) CreateMenuItemHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mitem model.MenuItem

		if err := ctx.ShouldBind(&mitem); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}
		//call biz
		newMItem, err := h.menuUC.CreateMenuItem(ctx, &mitem)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}
		ctx.JSON(http.StatusCreated, common.NewHttpSuccessResponse(http.StatusCreated, "Create new menu item successfully.", newMItem))
	}
}
