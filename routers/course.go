// description: 课程
//
// author: vignetting
// time: 2021/5/11

package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"structure/models"
	"structure/pkg/logger"
	"structure/pkg/result"
)

// @Summary		    查询课程
// @Tags            课程
// @Description     基于学生id 查询课程
// @Param		    studentId query int false "学生id"
// @Param		    pageNumber query int false "页号"
// @Param		    pageSize query int false "页大小"
// @Success 	    200 {object} result.Result{data=map[string]interface{}}
// @Router 		    /course/student/list [get]
func getStudentCourses(context *gin.Context) {
	var studentId, pageNumber, pageSize int
	var err error
	if studentId, err = strconv.Atoi(context.Param("studentId")); err != nil {
		context.JSON(http.StatusOK, result.Fail("studentId 格式错误或未填写", nil))
		logger.Debug(err.Error())
		return
	}
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

	pages := map[string]interface{}{"pageNumber": pageNumber, "pageSize": pageSize, "items": models.GetStudentCourses(studentId, pageNumber, pageSize)}
	context.JSON(http.StatusOK, result.Ok("获取成功", pages))
}

// @Summary		    查询课程
// @Tags            课程
// @Description     基于老师id、教室id 查询课程
// @Param		    teacherId query int false "老师id，非必要"
// @Param		    classRoomId query int false "教室id，非必要"
// @Param		    pageNumber query int false "页号"
// @Param		    pageSize query int false "页大小"
// @Success 	    200 {object} result.Result{data=map[string]interface{}}
// @Router 		    /course/list [get]
func getCourses(context *gin.Context) {
	var teacherId, classRoomId, pageNumber, pageSize int
	var err error
	teacherId, _ = strconv.Atoi(context.Query("teacherId"))
	classRoomId, _ = strconv.Atoi(context.Query("classRoomId"))
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

	pages := map[string]interface{}{"pageNumber": pageNumber, "pageSize": pageSize, "items": models.GetCourses(teacherId, classRoomId, pageNumber, pageSize)}
	context.JSON(http.StatusOK, result.Ok("获取成功", pages))
}
