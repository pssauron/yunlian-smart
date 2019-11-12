//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/5 1:54 下午
//
//============================================================

package userApi

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/yunlian/smart/models"

	"github.com/yunlian/smart/constants"
	"github.com/yunlian/smart/internal/storages"

	"github.com/pssauron/gocore/rs"

	"github.com/yunlian/smart/models/userModel"

	"github.com/labstack/echo"
)

func List(ctx echo.Context) error {
	page := ctx.Param("page")
	size := ctx.Param("size")

	p, err := strconv.Atoi(page)

	if err != nil {
		panic(rs.NewApiErr(998, "分页参数异常"))
	}

	g, err := strconv.Atoi(size)

	if err != nil {
		panic(rs.NewApiErr(998, "分页参数异常"))
	}

	return ctx.JSON(http.StatusOK, rs.NewResult(userModel.QueryPage(p, g)))
}

func Register(ctx echo.Context) error {

	//清除 cookie
	ctx.SetCookie(&http.Cookie{
		Name:     constants.TokenName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	req := userModel.RegReq{}

	if err := ctx.Bind(&req); err != nil {
		panic(rs.NewApiErr(999, "注册参数绑定异常"))
	}

	//TODO: 校验传参
	if req.Phone.IsEmpty() {
		models.PanicErr("注册手机号不能为空")
	}

	u := userModel.Register(&req)

	setCookie(ctx, u)
	return ctx.JSON(http.StatusOK, rs.NewResult(u))
}

func setCookie(ctx echo.Context, u *userModel.UserDTO) {
	tk := uuid()
	c := &http.Cookie{
		Name:     constants.TokenName,
		Value:    tk,
		Path:     "/",
		MaxAge:   int(24 * 60 * time.Minute), //不支持 IE6,7,8
		HttpOnly: true,
	}
	//将 token 保存在 redis 中
	err := storages.UseRedis.SetValueWithTimeout(constants.TokenPre+tk, u, constants.TokenExpire)
	if err != nil {
		panic(rs.NewApiErr(999, "token异常"))
	}
	ctx.SetCookie(c)
}

func uuid() string {
	t := time.Now().Unix()

	rand.Seed(t)

	s := fmt.Sprintf("%5d", rand.Intn(10000))

	s = fmt.Sprintf("%d%s", t, s)

	return base64.StdEncoding.EncodeToString([]byte(s))
}
