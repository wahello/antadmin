package main

import (
	_ "github.com/antbiz/antadmin/boot"
	_ "github.com/antbiz/antadmin/router"
	"github.com/gogf/gf/frame/g"
)

// @title       `antadmin`服务API
// @version     1.0
// @description `antadmin`服务API接口文档。
// @schemes     http
func main() {
	g.Server().Run()
}
