<h1 align="center">🎊🥂 Welcome to kv-services 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0.0.1-blue.svg?cacheSeconds=2592000" />

[comment]: <> (  <a href="https://github.com/kissvivi/kv-services/blob/main/LICENSE" target="_blank">)

[comment]: <> (    <img alt="License: Apache License" src="https://img.shields.io/badge/License-Apache License-yellow.svg" />)

[comment]: <> (  </a>)

[//]: # (  <a href="https://twitter.com/jobsvivi" target="_blank">)

[//]: # (    <img alt="Twitter: jobsvivi" src="https://img.shields.io/twitter/follow/jobsvivi.svg?style=social" />)

[//]: # (  </a>)
</p>

> 一个可扩展的golang快速开发脚手架内置RBAC权限模型，可直接上手开发其他业务，可改造性强，可单体开发可微服务开发

## 📅 功能计划（规划）
结构优化

## 开发日志
* 使用了泛型解决最基础的增删改查
* 此项目由v-services与kv-iot开发实践中总结而来，后面会有持续优化

### ✨ [kv-services 文档地址/问题记录](http://doc.kv-iot.cn/)

### 前端开源地址
* 暂无，开发中

## 🪄 Install 如何运行

### 开发方式运行
```sh
go mod tidy
go mod vendor

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

go build -o auth cmd/auth/main.go
```

### docker方式运行
打包docker镜像 待更新


运行服务
```sh
// 直接运行可执行文件
```

## 📝项目结构理念
### 服务划分
根据每个大的功能划分服务
- 授权以及用户服务（auth）

### 服务内结构划分
- data层 -> 数据操作层
- endpoint层 -> 数据暴露层
- service层 -> 业务逻辑层


### 关于我们
* 

### 基于此项目开发
kv-flow任务流框架

## Author

👤 **jobs_vivi**

* Twitter: [@jobsvivi](https://twitter.com/jobsvivi)
* Github: [@kissvivi](https://github.com/kissvivi)

## Show your support

Give a ⭐️ if this project helped you!

## Thanks 感谢赞助
<a href="https://jb.gg/OpenSourceSupport">
<img  src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" width="10%">
</a>

## 📝 License

[comment]: <> (Copyright © 2022 [jobs_vivi]&#40;https://github.com/kissvivi&#41;.<br />)

[comment]: <> (This project is [Apache License]&#40;https://github.com/kissvivi/kv-services/blob/main/LICENSE&#41; licensed.)

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_