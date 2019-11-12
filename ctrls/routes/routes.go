//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/5 1:56 下午
//
//============================================================

package routes

import (
	"net/http"

	"github.com/yunlian/smart/ctrls/apis/userApi"

	"github.com/labstack/echo"
	"github.com/pssauron/gocore/rs"
	"github.com/yunlian/smart/ctrls/middlewares"
)

func InitRoutes(e *echo.Echo) {

	e.HTTPErrorHandler = ApiErrHandler
	e.Use(middlewares.Recover())
	initUserRoute(e)
}

func initUserRoute(e *echo.Echo) {
	//这里可以将一组api 添加拦截器,类似 servlet 中的 filter
	g := e.Group("/api/user")

	g.POST("/register", userApi.Register)

	g.GET("/page/:page/:size", userApi.List)
}

func ApiErrHandler(e error, c echo.Context) {

	if err, ok := e.(*rs.ApiErr); ok {
		c.JSON(http.StatusOK, err)
	} else if err, ok := e.(*echo.HTTPError); ok {
		c.JSON(err.Code, err.Message)
	} else {
		c.JSON(http.StatusInternalServerError, e.Error())
	}

}
