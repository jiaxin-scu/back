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

// @title 			项目管理-人脸识别签到系统
// @version 		0.0.1
// @description 	基于 go 实现的后端，用于处理与保存人脸识别相关信息
// @basePath		/
func main() {
	routers.Run()
}
