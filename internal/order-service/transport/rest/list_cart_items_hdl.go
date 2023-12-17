package rest

import (
	"common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *orderHandler) ListCartItemsHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Value(common.CuserId).(int)
		cartitem, err := h.orderUc.ListCartItems(ctx, id)
		if cartitem == nil {
			fmt.Println("ok")
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}

		// cart := []model.CartItem{}
		// fmt.Println(cart)
		// fmt.Println(cartitem) // mảng rỗng
		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "List cart successfully", cartitem))
	}
}
