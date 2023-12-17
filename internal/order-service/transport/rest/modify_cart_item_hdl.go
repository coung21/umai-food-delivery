package rest

import (
	"common"
	"net/http"
	"order-service/model"
	"order-service/transport/grpc"

	"github.com/gin-gonic/gin"
)

func (h *orderHandler) ModifyCartHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.Value(common.CuserId).(int)

		var newCartItem model.CartItem
		if err := ctx.ShouldBind(&newCartItem); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		_, err := grpc.GetMenuItemHdl(h.grpcC.MenuC, newCartItem.ItemID)
		if err != nil {
			if err == common.NotFound {
				ctx.JSON(http.StatusConflict, common.NewRestErr(http.StatusConflict, err.Error(), err))
				return
			} else if err == common.BadQueryParams {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			}
		}

		val := h.orderUc.ModifyCart(ctx, uid, newCartItem.ItemID, newCartItem.Quantity)

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Modify Cart Successfully", val))

	}
}
