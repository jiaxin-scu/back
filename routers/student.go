// description: 用户 api
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

// @Summary		    获取学生
// @tags			学生
// @Description    	基于学生 id 获取学生信息
// @Param		    id path int false "学生id"
// @Success 	    200 {object} result.Result{data=models.Student}
// @Router 		    /student/{id} [get]
func getStudent(context *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(context.Param("id")); err != nil {
		context.JSON(http.StatusOK, result.Fail("id 格式错误或未填写", nil))
		logger.Debug(err.Error())
		return
	}

	if student, err := models.GetStudent(id); err != nil {
		context.JSON(http.StatusOK, result.Fail("该学生不存在", nil))
	} else {
		context.JSON(http.StatusOK, result.Ok("获取成功", student))
	}
}

// @Summary		    获取学生列表
// @tags			学生
// @Description    	分页获取学生列表
// @Param		    pageNumber query int false "页号"
// @Param		    pageSize query int false "页大小"
// @Success 	    200 {object} result.Result{data=map[string]interface{}}
// @Router 		    /student/list [get]
func getStudentList(context *gin.Context) {
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
	pages := map[string]interface{}{"pageNumber": pageNumber, "pageSize": pageSize, "items": models.GetStudents(pageNumber, pageSize)}
	context.JSON(http.StatusOK, result.Ok("获取成功", pages))
}

// @Summary		    更新学生信息
// @tags			学生
// @Description    	基于学生 ID 更新学生信息，需要更新哪些填写哪些信息，ID 是必要的
// @Param		    student body models.Student false "学生"
// @Success 	    200 {object} result.Result
// @Router 		    /student [put]
func updateStudent(context *gin.Context) {
	var student models.Student
	if err := context.BindJSON(&student); err != nil {
		logger.Debug(err.Error())
		context.JSON(http.StatusOK, result.Fail("缺少用户id", nil))
		return
	}
	models.UpdateStudent(student)
	context.JSON(http.StatusOK, result.Ok("更新成功", nil))
}
