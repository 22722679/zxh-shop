# 电商商城项目
基于Go语言开发的用于线上购物的电商商城。

## 目录
  - [网站地址](#网站地址)
  - [简介](#简介)
  - [项目技术栈](#项目技术栈)
  - [基础环境](#基础环境)
  - [文件目录](#文件目录)
  - [实现功能](#实现功能)
  - [项目启动步骤](#项目启动步骤)
  - [项目作者](项目作者)
  - [版本控制](版本控制)

## 网站地址
  该链接访问到目前网站：[www.zxiaohuselfshop.com](http://8.153.71.196:8080)

## 简介
  基于Go语言的hertz、kitex框架开发的电商商城项目，实现了用户登录注册(jwt鉴权)等基础功能，商品管理、购物车、订单管理与支付等功能。\n
  redis缓存商品信息，减轻数据库压力，运用consul实现多模块分布式服务注册与发现，降低服务的耦合性。\n
  基于protobuf的RPC系统定义接口和函数实现，使前后端能更快适配服务，支持多语言的开发。\n
  使用NATS消息队列实现购物车添加后的事件通知功能。运用prometheus实现项目数据指标以及多种性能指标的监控。
