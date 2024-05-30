package controller

import (
	"net/http"

	"github.com/kimzey/iskeai-shop/pkg/custom"
	_itemShopModel "github.com/kimzey/iskeai-shop/pkg/itemShop/model"
	_itemShopService "github.com/kimzey/iskeai-shop/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)
type ItemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl (itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &ItemShopControllerImpl{itemShopService}
}

func (c *ItemShopControllerImpl) Listing (pctx echo.Context) error{
	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoReauest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoReauest.Bind(itemFilter); err != nil {
		return custom.CuttomError(pctx,http.StatusBadRequest,err.Error())
	}

	itemModelList ,err := c.itemShopService.Listing(itemFilter)

	if err != nil {
		return custom.CuttomError(pctx,http.StatusInternalServerError,err.Error())
	}
	
	return pctx.JSON(http.StatusOK,itemModelList)
}