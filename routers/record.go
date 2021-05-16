// description: 记录
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
	"time"
)

// @Summary		    创建记录
// @Tags            记录
// @Description     基于学生id、教室id 插入记录
// @Param		    studentId query int false "学生id"
// @Param		    classRoomId query int false "教室id"
// @Success 	    200 {object} result.Result
// @Router 		    /record [post]
func insertRecord(context *gin.Context) {
	var studentId, classRoomId, courseId int
	var err error
	if studentId, err = strconv.Atoi(context.Query("studentId")); err != nil {
		context.JSON(http.StatusOK, result.Fail("studentId 格式错误或未填写", nil))
		logger.Debug(err.Error())
		return
	}
	if classRoomId, err = strconv.Atoi(context.Query("classRoomId")); err != nil {
		context.JSON(http.StatusOK, result.Fail("classRoomId 格式错误或未填写", nil))
		logger.Debug(err.Error())
		return
	}

	if courseId, err = models.GetCourseId(studentId, classRoomId); err != nil {
		context.JSON(http.StatusOK, result.Fail("当前学生没有这个课", nil))
		logger.Debug(err.Error())
		return
	}

	models.InsertRecords(models.Record{StudentId: studentId, CourseId: courseId})
	context.JSON(http.StatusOK, result.Ok("插入成功", nil))
}

type getRecordsParam struct {
	// 日期，格式为YYYY-MM-DD，非必须
	Date *time.Time `json:"date" form:"date" binding:"omitempty" time_format:"2006-01-02"`
	// 学生id，非必须
	StudentId *int `json:"studentId" form:"studentId" binding:"omitempty,min=1"`
	// 课程id，非必须
	CourseId *int `json:"courseId" form:"courseId" binding:"omitempty,min=1"`
	// 页号，从一开始
	PageNumber int `json:"pageNumber" form:"pageNumber" binding:"required,min=1"`
	// 页大小
	PageSize int `json:"pageSize" form:"pageSize" binding:"required,min=1"`
}

// @Summary		    查询记录
// @Tags            记录
// @Description     基于学生id、课程id 查询记录
// @Param		    getRecordsParam query routers.getRecordsParam false "查询记录请求对象"
// @Success 	    200 {object} result.Result{data=map[string]interface{}}
// @Router 		    /record/list [get]
func getRecords(context *gin.Context) {
	var param getRecordsParam
	var err error

	if err = context.ShouldBindQuery(&param); err != nil {
		context.JSON(http.StatusOK, result.Fail("参数不合规范"+err.Error(), nil))
		logger.Debug(err.Error())
		return
	}

	pages := map[string]interface{}{"pageNumber": param.PageNumber, "pageSize": param.PageSize, "items": models.GetRecords(param.StudentId, param.CourseId, param.Date, param.PageNumber, param.PageSize)}
	context.JSON(http.StatusOK, result.Ok("获取成功", pages))
}
