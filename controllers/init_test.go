package controllers

import (
	"runtime"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"

	"github.com/pangpanglabs/echosample/filters"
	"github.com/pangpanglabs/echosample/models"
)

var (
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
)

func init() {
	runtime.GOMAXPROCS(1)
	xormEngine, err := xorm.NewEngine("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	xormEngine.Sync(new(models.Discount))
	echoApp = echo.New()
	echoApp.Validator = &filters.Validator{}

	handleWithFilter = func(handlerFunc echo.HandlerFunc, c echo.Context) error {
		return filters.SetLogger("test")(filters.SetDbContext(xormEngine)(handlerFunc))(c)
	}
}