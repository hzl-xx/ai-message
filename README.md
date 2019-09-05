# ai-message
## 项目概述
- 名称：消息推送服务
- 项目代号：go-api

## 运行环境
- GO 1.12+
- Redis 3.0+
- RabbitMQ

## 开发环境部署
本项目使用go-micro+gin框架

### 1、安装 Protocol Buffers

> 参考官方文档

### 2、安装 Protobuf Golang 包

```$xslt
go get -u -v github.com/golang/protobuf/protoc-gen-go
```

#### 检查安装结果
```$xslt
protoc --version && which protoc-gen-go
```

#### 安装 gRPC 包

```$xslt
go get -u -v google.golang.org/grpc
```

#### 安装 golang.org/x/*
```$xslt
#!/bin/bash
MODULES="crypto net oauth2 sys text tools"
for module in ${MODULES}
do
    wget https://github.com/golang/${module}/archive/master.tar.gz -O ${GOPATH}/src/golang.org/x/${module}.tar.gz
    cd ${GOPATH}/src/golang.org/x && tar zxvf ${module}.tar.gz && mv ${module}-master/ ${module}
done
```

#### 安装 google.golang.org/genproto

### 3、安装 go-micro

```$xslt
go get github.com/micro/go-micro
```
#### 安装 protoc-gen-micro
```$xslt
go get github.com/micro/protoc-gen-micro
```
### 4、安装 gin
```$xslt
go get -u github.com/gin-gonic/gin
```

### 5、安装 govendor
```$xslt
go get -u -v github.com/kardianos/govendor
```

### 6、克隆源代码
> 克隆源码至 `GOPATH` 的 `src` 目录下

执行：
```$xslt
govendor sync
```

## 目录结构
- common  公共常量、函数目录
- conf  配置文件
- controller    控制器
- grpc  rpc服务
    - message  
        - conf  配置文件
        - configure 配置加载
        - handle    rpc服务handle
        - protos    proto文件
        - services  
        - storage   日志
- model     
- router    
- storage   
- utils     工具函数

## 服务启动

### rpc服务和队列
可执行文件目录 `ai-message/grpc/message`
```$xslt
go run server.go

go run consume.go
```

### 启动http服务
```$xslt
go run main.go
```

## 服务打包
```$xslt
go build main.go
go build server.go
go build consumego
```

## 配置文件

### http 服务配置

> /ai-message/conf/app.ini

```$xslt

[app]
PAGE_SIZE = 10  // 分页大小
JWT_SECRET = 23347$040412   // 加密key
JWT_EXPIRE_TIME = 24    // token有效期

[server]
HTTP_PORT = 8888    // 端口
READ_TIMEOUT = 60   
WRITE_TIMEOUT = 60

// mysql 配置
[database]
TYPE = mysql
USER = root
PASSWORD = hz251013
#127.0.0.1:3306
HOST = 127.0.0.1:3306
NAME = blog
TABLE_PREFIX =
CHARSET = utf8mb4

// redis配置
[redis]
HOST = 127.0.0.1:6379

```

### rpc 服务配置

> ai-message/grpc/message/conf/config.ini

```$xslt
// rpc服务名称 启动端口
[qywx_server]
SERVER_NAME = srv.send.message
HOST = 127.0.0.1:8881

// 企业微信配置
[company_wechat]
AGENTID = 1000017
SECRET = 3Q0ZLGo6nmt9jiEBKckkMke0Ua_M_UMxLbQgLCI1Z0k
CORPID = ww3386e86778990e91

[redis]
HOST = 127.0.0.1:6379

// token加密可以
[jwt_key]
KEY = aiGrpcToken

// 队列
[rabbit_mq]
RABBMITMQ_HOST = amqp://xiongwei:xinchan2@39.105.79.110:5672/go
RABBMITMQ_NAME = message-go
```

#### 服务目前支持：

- 企业微信群机器人消息推送（通用消息推送接口）
- 企业微信应用消息推送
- 邮件（开发中）

> 均支持 rpc 和 http

http接口：

`/v1/send/message`需要token验证

`/send/message`不需要token验证
> 请求参数参考 message.proto 文件结构体
