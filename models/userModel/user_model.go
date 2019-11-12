//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/5 1:55 下午
//
//============================================================

package userModel

import (
	"github.com/pssauron/gocore/libs"
	"github.com/pssauron/gocore/rs"
	"github.com/pssauron/gocore/utils/strutils"
	"github.com/yunlian/smart/internal/storages"
	"github.com/yunlian/smart/models"
)

func Register(reg *RegReq) *UserDTO {

	//检测手机号是否存在
	pe := CheckPhoneExists(reg.Phone.Get())

	if pe {
		models.PanicErr("手机号已被占用")
	}

	u := new(UserDTO)

	u.Phone = reg.Phone
	u.UserName = reg.UserName
	u.Password = libs.NewString(strutils.EncodePassword(reg.Password.Get()))

	err := storages.UseMysql.Insert(u)

	models.CheckErr(err, "注册用户异常")

	err = storages.UseMysql.Get(u, "select * from t_user where phone = ?", u.Phone.Get())
	models.CheckErr(err, "获取用户异常")

	return u

}

func QueryPage(page, size int) *rs.PageData {
	q := `select * from t_user`

	users := make([]UserDTO, 0)

	result, err := storages.UseMysql.QueryPage(&users, q, page, size)

	models.CheckErr(err, "查询用户分页异常")

	return result
}

func CheckPhoneExists(phone string) bool {
	q := `select count(*) from t_user where phone = ?`

	var count int

	err := storages.UseMysql.Get(&count, q, phone)

	models.CheckErr(err, "检测手机号是否存在发生异常")

	if count > 0 {
		return true
	}

	return false

}

func Login(req *LoginReq) *UserDTO {

	q := `select * from t_user where phone = ?`

	u := UserDTO{}

	err := storages.UseMysql.Get(&u, q, req.Phone.Get())

	models.CheckErr(err, "用户不存在")

	eq := strutils.ComparePassword(req.Password.Get(), u.Password.Get())
	if !eq {
		models.PanicErr("密码错误")
	}
	return &u
}
