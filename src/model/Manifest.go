package model

import "gorm.io/gorm"

type Manifest struct {
	gorm.Model
	Label       string `json:"label"`
	Attribution string `json:"attribution"`
	License     string `json:"license"`
	Thumbnail   string `json:"thumbnail"`
	UserId      uint   `json:"userId"` //Manifestを登録したユーザーのID
	Url         string `json:"url"`
}
