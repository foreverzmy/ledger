package db

import (
	"github.com/foreverzmy/ledger/model"
	"github.com/globalsign/mgo/bson"
	"time"
)

// AddAccount 创建账户
func AddAccount(name string, avatar string) error {
	err := C.Insert(&model.Account{
		ID:         bson.NewObjectId(),
		Name:       name,
		Avatar:     avatar,
		CreateTime: time.Now().UnixNano() / 1e6,
	})

	return err
}

// GetAccount 获取账户信息
func GetAccount(id string) (interface{}, error) {
	account := model.Account{}
	err := C.FindId(id).One(&account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// UpdateAccount 更新账户信息
func UpdateAccount(id bson.ObjectId) error {
	err := C.Update(bson.M{"_id": id},
		bson.M{"$set": bson.M{}})
	if err != nil {
		return err
	}
	return nil
}

// DeleteAccount 删除账户
func DeleteAccount(id bson.ObjectId) error {
	err := C.RemoveId(id)
	if err != nil {
		return err
	}
	return nil
}
