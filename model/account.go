package model

import "github.com/globalsign/mgo/bson"

// Account 账户
type Account struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `json:"name" bson:"name"`             // 用户名
	Avatar     string        `json:"avatar" bson:"avatar"`         // 头像
	CreateTime int64         `json:"createTime" bson:"createTime"` // 创建时间
	Assets     []string      `json:"assets" bson:"assets"`         // 资产 _id
	Bills      []string      `json:"bills bson:"bills"`            // 账单 _id
}
