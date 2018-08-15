package model

import "github.com/globalsign/mgo/bson"

// Account 账户
type Account struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`       // 用户名
	Avatar     string        `bson:"avatar"`     // 头像
	CreateTime int64         `bson:"createTime"` // 创建时间
	Assets     []string      `bson:"assets"`     // 资产 _id
	Bills      []string      `bson:"bills"`      // 账单 _id
}
