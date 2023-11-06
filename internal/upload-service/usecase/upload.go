package usecase

import (
	"bytes"
	"common"
	"context"
	"fmt"
	"image"
	_ "image/jpeg" // Import package xử lý hình ảnh JPEG
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
	"upload-service/component/uploadprovider"
	"upload-service/model"
)

type uploadUC struct {
	provider uploadprovider.UploadProvider
}

func NewUploadUC(provider uploadprovider.UploadProvider) *uploadUC {
	return &uploadUC{
		provider: provider,
	}
}

func (u *uploadUC) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	reader := bytes.NewReader(data)
	w, h, err := getImageDimension(reader)

	if err != nil {
		return nil, model.ErrFileIsNotImage
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := u.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, model.ErrCannotSaveFile
	}

	img.Width = w
	img.Height = h
	img.Ext = fileExt

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
