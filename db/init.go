package db

import (
	"github.com/globalsign/mgo"
)

// C 集合
var C *mgo.Collection

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
	C = session.DB("ledge").C("account")
}
