package service

import (
	"github.com/kimzey/iskeai-shop/entities"
	_itemShopModel "github.com/kimzey/iskeai-shop/pkg/itemShop/model"
	_itemShopRepository "github.com/kimzey/iskeai-shop/pkg/itemShop/repository"
)

type ItemShopServiceImpl struct{
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &ItemShopServiceImpl{itemShopRepository}
}

func (s *ItemShopServiceImpl) Listing (itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult,error){
	itemList,err := s.itemShopRepository.Listing(itemFilter)

	if  err!= nil {
		return nil,err
	}

	itemCouting,err := s.itemShopRepository.Counting(itemFilter)
	if err!=nil {
		return nil,err
	}

	totalPage := s.totalPageCalculation(itemCouting,int64(itemFilter.Size))

	return s.toItemResultResponse(itemList,int64(itemFilter.Page),totalPage),nil
}

func (s *ItemShopServiceImpl) totalPageCalculation(totalItems int64 , size int64) int64 {
	totalPage := totalItems / size

	if totalItems %size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *ItemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item,page,totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item,0)

	for _,item := range itemEntityList {
		itemModelList = append(itemModelList,item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page: page,
			TotalPage: totalPage,
		},
	}
}