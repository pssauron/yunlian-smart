//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/5 1:43 下午
//
//============================================================

package storages

import (
	"github.com/pssauron/gocore/stores"
	"github.com/yunlian/smart/internal/cfg"
)

var (
	UseMysql *stores.MyStore
	UseRedis *stores.RedisStore
)

func InitStorage() {
	UseMysql = stores.NewMyStore(cfg.Cfg.DB)
	UseRedis = stores.NewRedisStore(cfg.Cfg.Redis)
}
