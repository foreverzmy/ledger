package db

import (
	"github.com/globalsign/mgo"
)

// AccountC 账户集合
var AccountC *mgo.Collection

// AssetsC 资产集合
var AssetsC *mgo.Collection

// BillC 账单集合
var BillC *mgo.Collection

// init 初始化数据库连接
func init() {
	// 与数据库建立连接
	var session, err = mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// 使用 ledge 这个数据库
	ledger := session.DB("ledge")
	AccountC = ledger.C("account")
	AssetsC = ledger.C("assets")
	BillC = ledger.C("bill")
}
