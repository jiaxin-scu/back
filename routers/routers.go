// description: 路由
//
// author: vignetting
// time: 2021/5/10

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/unrolled/secure"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	_ "structure/docs"
	"structure/pkg/logger"
	"structure/pkg/result"
	"structure/pkg/setting"
	"time"
)

var router *gin.Engine

func init() {
	router = gin.New()

	if setting.ServerSetting.EnableLogger {
		router.Use(ginLogger(logger.Logger()))
	}
	if setting.ServerSetting.EnableRecovery {
		router.Use(ginRecovery(logger.Logger(), true))
	}
	if setting.ServerSetting.CrossDomain {
		router.Use(Cors())
	}
	if setting.ServerSetting.Ssl {
		router.Use(TlsHandler())
	}

	router.NoRoute(noRoute)
	router.NoMethod(noMethod)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func Run() {
	router.POST("/record", insertRecord)
	router.GET("/record/list", getRecords)

	router.GET("/course/student/list", getStudentList)
	router.GET("/course/list", getCourses)

	router.GET("/student/:id", getStudent)
	router.GET("/student/list", getStudentList)
	router.PUT("/student", updateStudent)

	router.GET("/teacher/:id", getTeacher)
	router.GET("/teacher/list", getTeacherList)
	router.PUT("/teacher", updateTeacher)

	if setting.ServerSetting.Ssl {
		if err := router.RunTLS(setting.ServerSetting.Ip+":"+strconv.Itoa(setting.ServerSetting.Port), setting.ServerSetting.SslPemPath, setting.ServerSetting.SslKeyPath); err != nil {
			panic("运行失败")
		}
	} else {
		if err := router.Run(setting.ServerSetting.Ip + ":" + strconv.Itoa(setting.ServerSetting.Port)); err != nil {
			panic("运行失败")
		}
	}
}

// 基于 zap 实现的 gin 日志
func ginLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// 基于 zap 实现的 gin recovery
func ginRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok")
		}

		defer func() {
			if err := recover(); err != nil {
				panic("跨域中间件出问题了，" + fmt.Sprintf("%v", err))
			}
		}()

		c.Next()
	}
}

// 引入 HTTPS
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     setting.ServerSetting.Ip + ":" + strconv.Itoa(setting.ServerSetting.Port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}

// 处理 404 请求
func noRoute(context *gin.Context) {
	context.JSON(http.StatusOK, result.Any(404, "无相应 url，请检查请求路径是否设置正确", nil))
}

// 处理 405 请求
func noMethod(context *gin.Context) {
	context.JSON(http.StatusOK, result.Any(405, "无对应 method，请检查请求方式是否设置正确", nil))
}
