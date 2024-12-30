# 电商商城项目
基于Go语言开发的用于线上购物的电商商城。

## 目录
  - [网站地址](#网站地址)
  - [简介](#简介)
  - [项目技术栈](#项目技术栈)
  - [基础环境](#基础环境)
  - [文件目录](#文件目录)
  - [项目启动步骤](#项目启动步骤)
  - [项目作者](项目作者)
  - [版本控制](版本控制)

## 网站地址
  #### 该链接访问到目前网站：[www.zxiaohuselfshop.com](http://8.153.71.196:8080)

## 简介
  基于Go语言的hertz、kitex框架开发的电商商城项目，实现了用户登录注册(jwt鉴权)等基础功能，商品管理、购物车、订单管理与支付等功能。  
  redis缓存商品信息，减轻数据库压力，运用consul实现多模块分布式服务注册与发现，降低服务的耦合性。  
  基于protobuf的RPC系统定义接口和函数实现，使前后端能更快适配服务，支持多语言的开发。  
  使用NATS消息队列实现购物车添加后的事件通知功能。运用prometheus实现项目数据指标以及多种性能指标的监控。  
## 技术栈
 - 前端
  - Hertz
  - Html
 - 后端
  - Kitex
  - consul
  - Nats
 - 数据库
  - Mysql
  - redis
 - 通用
  - protobuf
  - Dokcer
  - Go
  - prometheus
## 基础环境
  - Go 1.23.2
  - Docker 1.13.1
  - Mysql 5.7.43
  - Redis 5.0.7
  - consul 1.13.9
  - nats 2.9.22

## 文件目录
```go
├── config(配置信息)
│   ├── define.go
│   └── video_user_info.go
├── helper(Md5码,token,短信验证码)
│   └── helper.go
├── middleware(中间件)
│   ├── auto.go (鉴权)
│   └── zookeeper.go 
├── model (结构层)
│   ├── favorite.go
│   ├── feed.go
│   ├── init.go
│   ├── message_basic.go
│   ├── room_basic.go
│   ├── user_basic.go
│   └── user_room.go
├── router(路由)
│   └── router.go
├── service(服务层)
│   ├── chat.go
│   ├── user_basic.go
│   ├── favorite.go
│   ├── user_basic.go
│   ├── video_feed.go
│   └── websocket.go
├── test(测试)
│   ├── mongo.go
│   ├── redis.go
│   ├── websocket.go
│   └── zookeeper.go
├── go.mod
├── go.sum
├── README.md
└── main.go (主启动文件)
```

