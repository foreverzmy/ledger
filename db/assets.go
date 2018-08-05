package db

import (
	"github.com/globalsign/mgo/bson"

	"github.com/foreverzmy/ledger/model"
)

// AddAssets 添加资产
func AddAssets(id string, assets *model.Assets) error {
	err := C.UpdateId(bson.ObjectIdHex(id), bson.M{"$push": bson.M{
		"assets": assets,
	}})
	return err
}

// GetAssets 获取资产
func GetAssets(id bson.ObjectId) (model.Assets, error) {
	assets := model.Assets{}

	err := C.FindId(id).One(&assets)

	return assets, err
}
