package rest

import (
	"common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *menuHandler) DeleteMenuItemHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		mid, err := strconv.Atoi(ctx.Param("menu_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		delCount, err := h.menuUC.DeleteMenuItem(ctx, mid, rid)
		if err != nil {
			if err == common.BadQueryParams {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else if err == common.Forbidden {
				ctx.JSON(http.StatusForbidden, common.NewRestErr(http.StatusForbidden, err.Error(), err))
				return
			} else if err == common.NotFound {
				ctx.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Delete menu item successfully.", delCount))
	}
}
