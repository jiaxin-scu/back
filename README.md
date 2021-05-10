# common

## 简介

一个通用的工具包

注意：配置文件名称应该为 bootstrap.yaml，它应该被放在 conf 目录下

```yaml
# 日志配置
log:
  level: info                 # 日志保存级别，默认为 info
  filePath: ./logs/info.log   # 日志存储位置，默认为 ./logs/info.log
  fileMaxSize: 32             # 文件最大大小，单位为 MB，默认为 32
  maxBackups: 0               # 统一文件夹下文件的最大数量， 默认为 0 暨无最大文件现在
  maxRetentionTime: 0         # 一个日志文件最大保留时间， 默认为 0 暨无最大存活时间
  compress: false             # 文件是否压缩，默认为 true
  developerMode: false        # 是否为开发者模式，默认为 false

# 邮件配置
mail:
  host: host
  port: port
  userName: username
  password: password
  poolSize: n                 # 邮件连接池连接的数目，默认为 CPU 核数除二加一
  timeout: time               # 单位是毫秒，用于指定邮件发送的超时时间，默认为 3S

# 服务器配置
server:
  ip: 127.0.0.1               # 默认为 127.0.0.1
  port: 8080                  # 默认为 8080
  enableLogger: true          # 启用日志中间件，默认为 true
  enableRecovery: true        # 启用 recovery 中间件，用于捕获 panic，默认为 true

# 数据库配置
database:
  maxIdleConnection: 8        # 最大空闲连接数，默认为 CPU 核数
  maxOpenConnection: 8        # 最大连接数，默认为两倍 CPU 核数
  maxIdleTime: 60s            # 空闲连接最长存活时间，默认为 60s
  maxLifeTime: 60m            # 一个连接最长使用时间，默认为 60min
  url: userName:passWord@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local
```