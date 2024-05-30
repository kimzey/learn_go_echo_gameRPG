package repository

import (
	"github.com/kimzey/iskeai-shop/entities"	
	_itemShopModel "github.com/kimzey/iskeai-shop/pkg/itemShop/model"
)


type ItemShopRepository interface{
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item,error)	
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64,error)
}