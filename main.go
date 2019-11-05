//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/11/5 1:10 下午
//
//============================================================

package main

import (
	"flag"
	"runtime"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"

	"github.com/yunlian/smart/internal/storages"

	"github.com/pssauron/log4go"
	"github.com/yunlian/smart/internal/cfg"
)

func main() {
	//一般将 GOMAXPROCS 的个数设置为 CPU 的核数
	runtime.GOMAXPROCS(runtime.NumCPU())

	//解析命令行参数
	flag.Parse()

	//初始化 配置信息
	initLog()

	//初始化配置
	initCfg()

	//初始化数据库组件
	storages.InitStorage()

	//初始化 http 框架 参考 https://echo.labstack.com
	e := echo.New()

	//设置 recover,捕获业务代码中的 panic 错误,这里将重写recover 将http json 返回一致
	e.Use(middleware.Recover())

	if err := e.Start(":9999"); err != nil {
		panic("启动HTTP服务器异常")
	}

}

//初始化配置文件
func initCfg() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "./cfg.yaml", "配置文件地址")
	cfg.InitCfg(cfgPath)
}

func initLog() {
	var logPath string
	//初始化日志
	flag.StringVar(&logPath, "log", "./log.xml", "日志文件地址,接受命令行 -log 参数")

	//初始化日志,前端直接调用 log4go.Error||Info ... 记录日志
	log4go.LoadConfiguration(logPath)

}
