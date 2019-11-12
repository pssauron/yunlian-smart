//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/11 10:53 上午
//
//============================================================

package userModel

import (
	"github.com/pssauron/gocore/libs"
)

type UserDTO struct {
	UserID        libs.Int    `db:"user_id" json:"userId" primary:"true"`
	UserName      libs.String `db:"user_name" json:"userName"`
	Email         libs.String `db:"email" json:"email"`
	Phone         libs.String `db:"phone" json:"phone"`
	Password      libs.String `db:"password" json:"-"`
	LoginIP       libs.String `db:"login_ip" json:"loginIp"`
	LastLoginTime libs.Time   `db:"last_login_time" json:"lastLoginTime"`
	TS            libs.Time   `db:"ts" json:"ts"`
}

func (UserDTO) TableName() string {

	return "t_user"
}

type LoginReq struct {
	Phone    libs.String `json:"phone"`
	Password libs.String `json:"password"`
}

type RegReq struct {
	UserName libs.String `json:"userName"`
	Phone    libs.String `json:"phone"`
	Password libs.String `json:"password"`
}
