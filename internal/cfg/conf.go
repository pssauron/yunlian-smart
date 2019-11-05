//============================================================
// 描述: 配置文件
// 作者: Simon
// 日期: 2019/11/5 1:21 下午
//
//============================================================

package cfg

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/pssauron/gocore/stores"
)

var Cfg *Conf

type Conf struct {
	DB    *stores.StoreConf `yaml:"db"`
	Redis *stores.RedisConf `yaml:"redis"`
	//这里可以加一些其他配置
}

func InitCfg(cfgPath string) {
	bs, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(errors.New(fmt.Sprintf("配置文件加载异常；未找到文件[%s]", cfgPath)))
	}
	err = yaml.Unmarshal(bs, &Cfg)

	if err != nil {
		panic(errors.New("配置文件解析异常"))
	}

}
