package constants

import "time"

const (
	TokenName   = "token"               //token名
	TokenPre    = "session_"            //token redis 前缀
	TokenExpire = 24 * 60 * time.Minute //token 过期时间
)
