//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/11 2:09 下午
//
//============================================================

package middlewares

import (
	"fmt"

	"github.com/labstack/echo"
)

func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {

					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					c.Error(err)
				}
			}()

			return next(c)
		}
	}
}
