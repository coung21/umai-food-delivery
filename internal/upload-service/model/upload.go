package model

import (
	"common"
	"errors"
)

type Upload struct {
	common.SqlModel
	common.Image
}

var (
	ErrFileTooLarge   = errors.New("file too large")
	ErrFileIsNotImage = errors.New("file is not image")
	ErrCannotSaveFile = errors.New("cannot save uploaded file")
)
