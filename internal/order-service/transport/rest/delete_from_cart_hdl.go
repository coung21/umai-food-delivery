package rest

import (
	"common"
	"net/http"
	"order-service/model"

	"github.com/gin-gonic/gin"
)

func (h *orderHandler) DeleteItemFromCartHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.Value(common.CuserId).(int)

		del := model.DeletingReq{}

		if err := ctx.ShouldBind(&del); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		val := h.orderUc.DeleteItemFromCart(ctx, uid, del.ItemsID)

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Delete Items From Cart Successfully", val))
	}
}
