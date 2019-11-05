//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/5 1:56 下午
//
//============================================================

package routes

import (
	"github.com/labstack/echo"
	"github.com/yunlian/smart/ctrls/apis/userApi"
)

func InitRoutes(e *echo.Echo) {

	initUserRoute(e)
}

func initUserRoute(e *echo.Echo) {
	//这里可以将一组api 添加拦截器,类似 servlet 中的 filter
	g := e.Group("/api/user")

	g.POST("/list", userApi.List)

}
