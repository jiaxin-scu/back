// description: 用于统一加载配置
//
// author: vignetting
// time: 2021/5/10

package setting

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
	"runtime"
	"time"
)

type Log struct {
	Level            string `validate:"required,oneof=debug info warn error dpanic panic fatal"`
	FilePath         string `validate:"required"`
	FileMaxSize      int    `validate:"min=1,max=1024"`
	MaxBackups       int    `validate:"min=0,max=1024"`
	MaxRetentionTime int    `validate:"min=0,max=365"`
	Compress         bool
	DeveloperMode    bool
}

var LogSetting = &Log{}

type Mail struct {
	Host     string        `validate:"required,fqdn"`
	Port     int           `validate:"min=1,max=65535"`
	UserName string        `validate:"required,email"`
	Password string        `validate:"required"`
	PoolSize int           `validate:"min=1,max=1024"`
	Timeout  time.Duration `validate:"required"`
}

var MailSetting = &Mail{}

type Server struct {
	Ip             string `validate:"ipv4"`
	Port           int    `validate:"min=0,max=65535"`
	EnableLogger   bool
	EnableRecovery bool
}

var ServerSetting = &Server{}

type Database struct {
	MaxIdleConnection int           `validate:"min=1,max=1024"`
	MaxOpenConnection int           `validate:"min=1,max=4096"`
	MaxIdleTime       time.Duration `validate:"required"`
	MaxLifeTime       time.Duration `validate:"required"`
	Url               string        `validate:"required"`
}

var DatabaseSetting = &Database{}

func Setup() {
	// 1. 创建配置读取文件
	v := viper.New()

	// 2. 设置文件名与类型
	v.SetConfigName("bootstrap")
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf/")

	// 3. 设置默认值
	v.Set("log.level", "info")
	v.Set("log.filePath", "./logs/info.log")
	v.Set("log.fileMaxSize", 32)
	v.Set("log.maxBackups", 0)
	v.Set("log.maxRetentionTime", 0)
	v.Set("log.compress", true)
	v.Set("log.developerMode", false)

	v.SetDefault("mail.poolSize", runtime.NumCPU()/2+1)
	v.SetDefault("mail.timeout", 3000000000)

	v.SetDefault("server.ip", "127.0.0.1")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.enableLogger", true)
	v.SetDefault("server.enableRecovery", true)

	v.SetDefault("database.maxIdleConnection", runtime.NumCPU())
	v.SetDefault("database.maxIdleConnection", 2*runtime.NumCPU())
	v.SetDefault("database.maxIdleTime", 360)
	v.SetDefault("database.maxLifeTime", 60)

	// 4. 读取配置
	var err error
	if err = v.ReadInConfig(); err != nil {
		panic("配置文件读取失败，" + err.Error())
	}

	// 5. 载入配置
	if err = v.UnmarshalKey("log", LogSetting); err != nil {
		panic("日志配置载入失败" + err.Error())
	}
	if err = v.UnmarshalKey("mail", MailSetting); err != nil {
		panic("邮件配置载入失败" + err.Error())
	}
	if err = v.UnmarshalKey("server", ServerSetting); err != nil {
		panic("服务器配置载入失败" + err.Error())
	}
	if err = v.UnmarshalKey("database", DatabaseSetting); err != nil {
		panic("数据库配置载入失败" + err.Error())
	}

	// 6. 格式校验
	validate := validator.New()
	if err = validate.Struct(LogSetting); err != nil {
		panic("日志配置不合规范" + err.Error())
	}
	if err = validate.Struct(MailSetting); err != nil {
		panic("邮件配置不合规范" + err.Error())
	}
	if err = validate.Struct(ServerSetting); err != nil {
		panic("服务器配置不合规范" + err.Error())
	}
	if err = validate.Struct(DatabaseSetting); err != nil {
		panic("数据库配置不合规范" + err.Error())
	}
}
