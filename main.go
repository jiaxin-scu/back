// description:
//
// author: vignetting
// time: 2021/5/10

package main

import (
	"structure/models"
	"structure/pkg/logger"
	"structure/pkg/mail"
	"structure/pkg/setting"
	"structure/routers"
)

func init() {
	setting.Setup()
	logger.Setup()
	mail.Setup()
	models.SetUp()
}

// @title 			项目名
// @version 		版本
// @description 	简介
// @basePath		/
func main() {
	routers.Run()
}
