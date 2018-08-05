package model

import (
	"github.com/globalsign/mgo/bson"
)

// Bill 账单
type Bill struct {
	Payment bson.ObjectId // 支出账户
	Time    int64         // 时间
	Amount  int64         // 金额
	Class   int           // 支出类型
	Remake  string        // 备注
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
