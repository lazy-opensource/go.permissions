package main

import (
	"core/code/controller/system"
	"lzy/framework/web"
)

//初始化
func Dinit() {
	//缓存html模板
	web.InitCacheTemplate(web.HtmlDirPath)
}

func main() {

	//初始化路由
	controller.InitRoute()

	//启动路由
	web.StartRoute()

}
