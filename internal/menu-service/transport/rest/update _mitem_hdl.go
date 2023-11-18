package rest

import (
	"common"
	"menu-service/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *menuHandler) UpdateMenuItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		mid := ctx.Param("menu_id")
		var upd model.UpdateMenuItem
		if err := ctx.ShouldBind(&upd); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		newMitem, err := h.menuUC.UpdateMenuItem(ctx, rid, mid, &upd)
		if err != nil {
			if err == common.BadQueryParams {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else if err == common.NotFound {
				ctx.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Update menu item successfully.", newMitem))
	}
}
