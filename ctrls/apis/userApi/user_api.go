//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/5 1:54 下午
//
//============================================================

package userApi

import (
	"net/http"

	"github.com/labstack/echo"
)

func List(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "hello world")
}
