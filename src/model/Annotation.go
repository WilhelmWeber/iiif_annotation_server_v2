package model

import "gorm.io/gorm"

type Annotation struct {
	gorm.Model
	UserId     uint   `json:"userId"`     //登録したユーザーのID
	ManifestId uint   `json:"manifestId"` //紐づけられているマニフェストのID
	CanvasPage uint   `json:"canvasPage"` //紐づけられているキャンバスの番号(annnotationListの管理に必要)
	Source     string `json:"canvasUri"`  //紐づけられているキャンバスのURI
	Motivation string `json:"motivation"`
	Selector   string `json:"selector"`
	Body       string `json:"body"` //注釈本文
}
