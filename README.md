打点监控系统
===========

本系统主要实现已lib的方式提供如下功能:
1. 心跳注册
2. 监控数据上传(默认间隔30s上传)

要求所有利用go语言实现的后端服务都嵌入该模块，做好统一的服务监控

## 后期优化

1. 提高Add的性能
2. 心跳上传暂时不考虑性能问题(net/http)

## 语言和协议支持情况

|lang | redis | http | shareMem|
|---|---|---|---|
|go|ok|ok|todo|
|c/c++|ok|ok|todo|
|python|ok|ok|todo|
|java|ok|ok|no|
|rust|ok|ok|todo|


