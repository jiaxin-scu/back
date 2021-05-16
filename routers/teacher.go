// description: 老师相关 api
//
// author: vignetting
// time: 2021/5/10

package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"structure/models"
	"structure/pkg/logger"
	"structure/pkg/result"
)

// @Summary		    获取老师
// @tags			老师
// @Description    	基于老师 id 获取老师信息
// @Param		    id path int false "老师id"
// @Success 	    200 {object} result.Result{data=models.Teacher}
// @Router 		    /teacher/{id} [get]
func getTeacher(context *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(context.Param("id")); err != nil {
		context.JSON(http.StatusOK, result.Fail("id 格式错误或未填写", nil))
		logger.Debug(err.Error())
		return
	}

	if teacher, err := models.GetTeacher(id); err != nil {
		context.JSON(http.StatusOK, result.Fail("老师不存在", nil))
	} else {
		context.JSON(http.StatusOK, result.Ok("获取成功", teacher))
	}

}

// @Summary		    获取老师列表
// @tags			老师
// @Description    	分页获取老师列表
// @Param		    pageNumber query int false "页号"
// @Param		    pageSize query int false "页大小"
// @Success 	    200 {object} result.Result{data=map[string]interface{}}
// @Router 		    /teacher/list [get]
func getTeacherList(context *gin.Context) {
	var pageNumber, pageSize int
	var err error
	if pageNumber, err = strconv.Atoi(context.Query("pageNumber")); err != nil {
		logger.Debug(err.Error())
		context.JSON(http.StatusOK, result.Fail("页号未填写", nil))
		return
	}
	if pageSize, err = strconv.Atoi(context.Query("pageSize")); err != nil {
		logger.Debug(err.Error())
		context.JSON(http.StatusOK, result.Fail("页号未填写", nil))
		return
	}
	pages := map[string]interface{}{"pageNumber": pageNumber, "pageSize": pageSize, "items": models.GetTeachers(pageNumber, pageSize)}
	context.JSON(http.StatusOK, result.Ok("获取成功", pages))
}

// @Summary		    更新老师信息
// @tags			老师
// @Description    	基于老师 ID 更新老师信息，需要更新哪些填写哪些信息，ID 是必要的
// @Param		    pageNumber body models.Teacher false "老师"
// @Success 	    200 {object} result.Result
// @Router 		    /teacher [put]
func updateTeacher(context *gin.Context) {
	var teacher models.Teacher
	if err := context.BindJSON(&teacher); err != nil {
		logger.Debug(err.Error())
		context.JSON(http.StatusOK, result.Fail("缺少用户id", nil))
		return
	}
	models.UpdateTeacher(teacher)
	context.JSON(http.StatusOK, result.Ok("更新成功", nil))
}
