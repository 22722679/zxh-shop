# 电商商城项目
基于Go语言开发的用于线上购物的电商商城。

## 目录
  - [网站地址](#网站地址)
  - [简介](#简介)
  - [项目技术栈](#项目技术栈)
  - [基础环境](#基础环境)
  - [文件目录](#文件目录)
  - [项目启动步骤](#项目启动步骤)
  - [项目作者](#项目作者)
  - [版本控制](#版本控制)
  - [鸣谢](#鸣谢)

## 网站地址
  #### 该链接访问到目前网站：[www.zxiaohuselfshop.com](http://www.zxiaohuselfshop.com)
  #### 基于该项目的自动化测试演示以及测试源码：[UI自动化测试演示](https://www.bilibili.com/video/BV19ScceHEMa/?vd_source=4cbd2d924b1efbce235b9f288896cf6f)，[测试源码](https://github.com/22722679/UI-automation-zxh-shop)

## 简介
  **基于Go语言的hertz、kitex框架开发的电商商城项目，实现了用户登录注册(jwt鉴权)等基础功能，商品管理、购物车、订单管理与支付等功能。**  
  **redis缓存商品信息，减轻数据库压力，运用consul实现多模块分布式服务注册与发现，降低服务的耦合性。**  
  **基于protobuf的RPC系统定义接口和函数实现，使前后端能更快适配服务，支持多语言的开发。使用NATS消息队列实现购物车添加后的事件通知功能。运用prometheus实现项目数据指标以及多种性能指标的监控(计划实现服务容器化部署)。**  
## 项目技术栈
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
├── app(主要服务存放位置)
│   ├── cart(购物车)
│   ├── checkout(支付流程结算和检验)
│   ├── email(消息通知)
│   ├── frontend(前端)
│   ├── order(订单服务)
│   ├── payment(支付服务)
│   ├── product(商品服务)
│   └── user(用户服务)
├── common(模板函数存放)
├── db (数据库创建语句log)
├── demo(proto和thrift两种序列化协议测试)
├── deploy(docker容器化部署)
├── hello_world(测试)
├── idl(protobuf序列化脚手架)
├── rpc_gen(rpc调用关联)
├── Makefile(make项目管理工具)
├── docker-compose.yaml(docker配置文件)
├── go.work
├── go.work.sum
└── README.md
```

## 项目启动步骤
  1. 下载源码
  ```sh
    git clone https://github.com/22722679/zxh-shop.git
  ```
  2. 启动docker上部署的中间件(Mysql、Redis、consul、nats、prometheus等等)，我们使用了docker-compose 启动，需要自己配置相关文件(docker-compose.yaml)：
  ```sh
    docker-compose up
  ```
  3. 启动服务(使用make工具直接启动8个服务)
  ```sh
    make shop
  ```

## 项目作者
  ### **张小虎**  


  
## 版本控制
该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

## 鸣谢
  
- [字节跳动青训营](https://youthcamp.bytedance.com/)
- 感谢[cloudewego](https://www.cloudwego.cn/zh/)社区提供的框架支持







