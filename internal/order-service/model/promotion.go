package model

import (
	"common"
	"time"
)

type Promotion struct {
	common.SqlModel
	Code     string
	Discount float32
	Exp      time.Time
}
