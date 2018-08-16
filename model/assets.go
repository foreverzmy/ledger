package model

import "github.com/globalsign/mgo/bson"

// Assets 资产
type Assets struct {
	ID     bson.ObjectId `bson:"_id"`
	Name   string        `json:"name" bson:"name"`   // 资产名
	Icon   string        `json:"icon" bson:"icon"`   // 图标
	Amount int64         `json:"amount" bson:"amount"` // 资产数
	Class  AssetsTypes   `json:"class" bson:"class"`  // 资产类型
}

// AssetsTypes 资产类型
type AssetsTypes int

// 资产类型
const (
	_       AssetsTypes = iota
	Cash                // 现金
	Bedit               // 借记卡
	Credit              // 信用卡
	Virtual             // 虚拟账户
	Invest              // 投资账户
	Loan                // 负债账户
	Lend                // 债权账户
)
