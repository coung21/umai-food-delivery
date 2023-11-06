package upload

import (
	"common"
	"context"
)

type Usecase interface {
	Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error)
}
