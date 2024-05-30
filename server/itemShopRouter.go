package server

import (
	_itemShopRepository "github.com/kimzey/iskeai-shop/pkg/itemShop/repository"
	_itemShopService "github.com/kimzey/iskeai-shop/pkg/itemShop/service"
	_itemShopController "github.com/kimzey/iskeai-shop/pkg/itemShop/controller"

)

func (s *echoServer) initItemShopRouter () {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db,s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopControler := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("/",itemShopControler.Listing)
}