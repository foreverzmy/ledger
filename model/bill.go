package model

import (
	"github.com/globalsign/mgo/bson"
)

// Bill 账单
type Bill struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Payment string        `json:"payment" bson:"payment"` // 支出账户
	Time    int64         `json:"time" bson:"time"`       // 时间
	Amount  int64         `json:"amount" bson:"amount"`   // 金额
	Class   int           `json:"class" bson:"class"`     // 支出类型
	Remark  string        `json:"remake" bson:"remark"`   // 备注
}

// BillTypes 账单类型
type BillTypes int

// 账单类型
const (
	_          BillTypes = iota
	Food                 // 餐饮
	Shoppong             // 购物
	Traffic              // 交通
	Vegetables           // 蔬菜
	Fruit                // 水果
	Fun                  // 娱乐
	Auto                 // 自定义
)
