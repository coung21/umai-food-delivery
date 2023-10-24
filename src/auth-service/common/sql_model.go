package common

import (
	"time"

	"github.com/sqids/sqids-go"
)

type SqlModel struct {
	ID        uint64    `json:"-" gorm:"column:id"`
	FakeID    string    `json:"id" gorm:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
}

func (s *SqlModel) EncodeId() {
	sq, _ := sqids.New(sqids.Options{
		Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		MinLength: 12,
	})

	fakeID, _ := sq.Encode([]uint64{s.ID})
	s.FakeID = fakeID
}

func (s *SqlModel) DecodeId() uint64 {
	sq, _ := sqids.New(sqids.Options{
		Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		MinLength: 12,
	})
	id := sq.Decode(s.FakeID)

	var realId uint64

	for _, num := range id {
		realId = realId*10 + num
	}

	return realId
}
