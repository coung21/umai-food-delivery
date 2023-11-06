package common

type Image struct {
	// Id     int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	Ext       string `json:"ext" gorm:"column:ext"`
	CloudName string `json:"cloud_name" gorm:"column:cloud_name"`
}

// func (Image) TableName() string {
// 	return "images"
// }
