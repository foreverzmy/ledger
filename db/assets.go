package db

import (
	"github.com/globalsign/mgo/bson"

	"github.com/foreverzmy/ledger/model"
)

// AddAssets 添加资产
func AddAssets(id string, assets *model.Assets) error {

	err := AssetsC.Insert(assets)
	if err != nil {
		return err
	}

	err = AccountC.UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{
		"assets": assets.ID,
	}})

	return err
}

// GetAssets 获取资产
func GetAssets(id bson.ObjectId) (model.Assets, error) {
	assets := model.Assets{}

	err := AssetsC.FindId(id).One(&assets)

	return assets, err
}
