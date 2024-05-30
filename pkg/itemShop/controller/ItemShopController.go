package controller

import "github.com/labstack/echo/v4"

type ItemShopController interface{
	Listing (pcrx echo.Context) error
}