package db

import (
	"github.com/foreverzmy/ledger/model"
	"github.com/globalsign/mgo/bson"
)

// AddBill 添加账单
func AddBill(id string, bill *model.Bill) error {
	err := C.UpdateId(bson.ObjectIdHex(id), bson.M{"$push": bson.M{
		"bills": bill,
	}})
	return err
}

// GetBill 获取账单
func GetBill(id bson.ObjectId) (model.Bill, error) {
	bill := model.Bill{}

	err := C.FindId(id).One(&bill)

	return bill, err
}
