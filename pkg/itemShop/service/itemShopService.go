package service

import (
	_itemShopModel "github.com/kimzey/iskeai-shop/pkg/itemShop/model"
)

type ItemShopService interface{
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult,error)
}