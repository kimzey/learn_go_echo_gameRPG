package repository

import (
	"github.com/kimzey/iskeai-shop/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	_itemShopException "github.com/kimzey/iskeai-shop/pkg/itemShop/exception"
	_itemShopModel "github.com/kimzey/iskeai-shop/pkg/itemShop/model"

)

type ItemShopRepositoryImpl struct{
	db *gorm.DB
	logger echo.Logger

}

func NewItemShopRepositoryImpl (db *gorm.DB,Logger echo.Logger) ItemShopRepository{
	return &ItemShopRepositoryImpl{db,Logger}
}

func (r *ItemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item,error){

	itemList := make([]*entities.Item,0)

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?",false)

	if itemFilter.Name != ""{
		query = query.Where("name ilike ?","%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != ""{
		query = query.Where("description ilike ?","%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page-1)*itemFilter.Size)
	limit := int(itemFilter.Size)


	if err := query.Offset(offset).Limit(limit).Find(&itemList).Order("id asc").Error; err != nil {
		r.logger.Error("Fail to list item : %s",err.Error())
		return nil,&_itemShopException.ItemListing{}
	}

	return itemList,nil
}

func (r *ItemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64,error){
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?",false)

	if itemFilter.Name != ""{
		query = query.Where("name ilike ?","%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != ""{
		query = query.Where("description ilike ?","%"+itemFilter.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Couning item failed : %s",err.Error())
		return -1,&_itemShopException.ItemCounting{}
	}

	return count,nil
}