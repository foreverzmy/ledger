package object

import (
	"github.com/foreverzmy/ledger/db"
	"github.com/foreverzmy/ledger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/graphql-go/graphql"
)

var addAssetsField = graphql.Field{
	Type:        graphql.Boolean,
	Description: "添加资产",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "账户 id",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Description: "资产名",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"icon": &graphql.ArgumentConfig{
			Description:  "图标",
			Type:         graphql.String,
			DefaultValue: "",
		},
		"amount": &graphql.ArgumentConfig{
			Description:  "资产数",
			DefaultValue: 0,
			Type:         graphql.NewNonNull(graphql.Int),
		},
		"class": &graphql.ArgumentConfig{
			Description: "资产类型",
			Type:        graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)
		name := p.Args["name"].(string)
		icon := p.Args["icon"].(string)
		amount := int64(p.Args["amount"].(int))
		class := p.Args["class"].(int)

		assets := &model.Assets{
			ID:     bson.NewObjectId(),
			Name:   name,
			Icon:   icon,
			Amount: amount,
			Class:  model.AssetsTypes(class),
		}

		err := db.AddAssets(id, assets)
		if err != nil {
			return false, err
		}
		return true, nil
	},
}

var assetField = graphql.Fields{
	"id": &graphql.Field{
		Description: "资产 ID",
		Type:        graphql.String,
	},
	"name": &graphql.Field{
		Description: "资产名",
		Type:        graphql.String,
	},
	"icon": &graphql.Field{
		Description: "图标",
		Type:        graphql.String,
	},
	"amount": &graphql.Field{
		Description: "资产数",
		Type:        graphql.Int,
	},
	"class": &graphql.Field{
		Description: "资产类型",
		Type:        graphql.String,
	},
}

var queryAssetField = graphql.Field{
	Description: "用户资产",
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name:   "asset",
		Fields: assetField,
	}),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "资产 ID",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)

		asset, err := db.GetAsset(id)

		return asset, err
	},
}

var queryAssetsField = graphql.Field{
	Description: "用户资产列表",
	Type: graphql.NewList(
		graphql.NewObject(graphql.ObjectConfig{
			Name:   "assets",
			Fields: assetField,
		},
		),
	),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "资产 ID",
			Type:        graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		assetsList, err := db.GetAssets()
		return assetsList, err
	},
}
