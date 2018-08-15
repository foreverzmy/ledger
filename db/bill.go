package db

import (
	"errors"

	"github.com/foreverzmy/ledger/model"
	"github.com/globalsign/mgo/bson"
)

// AddBill 添加账单
func AddBill(id string, bill *model.Bill) error {

	var assets model.Assets

	err := AssetsC.FindId(bson.ObjectIdHex(bill.Payment)).One(&assets)
	if err != nil {
		return err
	}

	if assets.Amount < bill.Amount {
		return errors.New("支出大与存款")
	}

	err = BillC.Insert(bill)
	if err != nil {
		return err
	}

	err = AccountC.UpdateId(
		bson.ObjectIdHex(id),
		bson.M{
			"$addToSet": bson.M{
				"bills": bill.ID,
			},
		},
	)

	if err != nil {
		return err
	}

	err = AssetsC.UpdateId(
		bson.ObjectIdHex(bill.Payment),
		bson.M{
			"$set": bson.M{
				"amount": assets.Amount - bill.Amount,
			},
		},
	)

	return err
}

// GetBill 获取账单
func GetBill(id bson.ObjectId) (model.Bill, error) {
	bill := model.Bill{}

	err := BillC.FindId(id).One(&bill)

	return bill, err
}
