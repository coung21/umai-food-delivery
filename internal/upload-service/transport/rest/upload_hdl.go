package rest

import (
	"common"
	"net/http"
	upload "upload-service/interfaces"
	"upload-service/model"

	"github.com/gin-gonic/gin"
)

type uploadHdl struct {
	uploadUC upload.Usecase
}

func NewUploadHdl(uploadUC upload.Usecase) *uploadHdl {
	return &uploadHdl{
		uploadUC: uploadUC,
	}
}

func (h *uploadHdl) UploadHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		folder := ctx.DefaultPostForm("folder", "image")

		file, err := fileHeader.Open()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {

		}

		img, err := h.uploadUC.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			if err == model.ErrFileIsNotImage || err == model.ErrFileTooLarge {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Upload image successfully.", img))
	}
}
