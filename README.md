<div align="center">
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**中文** | [English](./README.En.md) 
---
[![Go-Zero](https://img.shields.io/badge/Go--Zero-v1.9.4-brightgreen.svg)](https://go-zero.dev/)
[![Ent](https://img.shields.io/badge/Ent-v0.14.5-blue.svg)](https://entgo.io/)
[![Casbin](https://img.shields.io/badge/Casbin-v3.8.1-orange.svg)](https://github.com/casbin/casbin)
[![Release](https://img.shields.io/badge/Release-v1.8.0-green.svg)](https://github.com/suyuan32/simple-admin-core/releases)
[![License: Apache2.0](https://img.shields.io/badge/License-Apache2.0-yellow.svg)](https://opensource.org/licenses/MIT)
![公众号](https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E5%87%A0%E9%A2%97%E9%85%A5-blue)
![注意](https://img.shields.io/badge/%E6%B3%A8%E6%84%8F-%E5%85%B3%E6%B3%A8%E5%85%AC%E4%BC%97%E5%8F%B7%E5%8A%A0%E5%85%A5%E5%BE%AE%E4%BF%A1%E7%BE%A4-blue)

## 简介

Simple Admin 是一个强大的、易扩展的后台管理系统，基于 Go-Zero、Vben Admin、Ent、Casbin
等开源项目构建，提供了完整的用户管理、权限管理、角色管理、菜单管理、日志管理、配置管理等功能，支持多语言等特性，适用于小型或大型企业快速搭建分布式后台管理系统。 <br><br>
Simple Admin 有完善的开发部署工具， 十分适合高并发、高可靠、复杂的业务场景，项目可以方便地一键升级，提供了完善的文档视频和示例，让开发者可以快速上手，快速开发。官方提供了
6 大免费基础模块，可以满足 80 % 的常用业务需求。同时提供了模块商店，可以方便的购买和使用更多的模块。<br><br>
6 大免费模块均可免费商用，适合开发者学习、企业内部使用、个人项目使用等，欢迎大家使用和反馈问题，我们会持续更新和维护。

> 若您拟将本项目用于商业用途（包括所有产生收益的商业行为），请严格遵守 Apache 2.0
> 开源协议，同时完整保留作者技术支持声明。您需确保项目相关的版权声明信息（含文案、日志及代码中内嵌的版权声明）全部保留，该类信息仅为合规性文案，不会对您的业务功能及运营产生任何影响。若您需剔除相关版权声明或用于商业用途，应先行[购买相应授权](https://simple-admin-official.simple-hub.top/commercial)。

## 特性

- **最新技术栈**：使用 ent, Casbin, kafka 等前沿技术开发
- **完全支持go-swagger**: 直接在api文件内编写注释即可直接生成swagger文档
- **统一的错误处理**: 整个系统拥有国际化的统一错误处理
- **国际化**：内置完善的国际化方案
- **服务注册发现**: 完善的服务注册发现机制，原生支持K8s
- **权限**: 内置完善的动态路由权限生成方案, 集成RBAC权限控制
- **代码生成**: 内置三端 Web, API, RPC 代码生成
- **多种扩展**: 提供多种扩展，同时具有非常简单的接入功能
- **其他**: 流量控制， ES服务
- **ORM**: 强大的 Ent 框架支持

## 支持功能

- 用户管理：管理系统用户数据
- 部门管理：管理所属部门
- 岗位管理：配置系统用户所属担任职务
- 菜单管理：配置系统菜单，树形展示
- 角色管理：管理角色权限，支持多角色
- 字典管理：维护数据字典，方便前端使用
- 接口文档：根据业务代码自动生成相关的api接口文档
- 代码生成：自动生成 CRUD 代码，快捷生成自定义逻辑
- 令牌管理：管理 token 状态，支持拉黑 token

### 在线预览

#### [免费版在线预览](https://vben5-preview.ryansu.tech/)

- 账号 **admin**
- 密码 **simple-admin**

#### [多租户版在线预览](https://tenant-preview.ryansu.tech/)

- 管理员租户账号
    - 企业： **admin**
    - 账号: **admin**
    - 密码: **simple-admin**

- 租户账号
    - 企业: **测试企业**
    - 账号: **admin**
    - 密码: **simple-admin**

> 只读，不可修改和注册

## 免费的官方模块

| 模块名称       | 模块介绍 | 模块地址                                                                      |
|------------|------|---------------------------------------------------------------------------|
| Core       | 核心模块 | [Core](https://github.com/suyuan32/simple-admin-core)                     |
| Backend UI | 后端界面 | [Backend UI](https://github.com/suyuan32/simple-admin-vben5-ui)           |
| FMS        | 文件管理 | [File](https://github.com/suyuan32/simple-admin-file)                     |
| Job        | 定时任务 | [Job](https://github.com/suyuan32/simple-admin-job)                       |
| MMS        | 会员管理 | [Member](https://github.com/suyuan32/simple-admin-member-api)             |
| MCMS       | 消息中心 | [Message Center](https://github.com/suyuan32/simple-admin-message-center) |


# 社区模块

[点击查看](https://github.com/suyuan32/awesome-simple-admin-module)


## 文档

### [Simple Admin 中文文档](https://doc.ryansu.tech/zh)

## Stars

[![Star History Chart](https://api.star-history.com/svg?repos=suyuan32/simple-admin-core&type=Date)](https://github.com/suyuan32/simple-admin-core)

## 维护者

[@Ryan Su](https://github.com/suyuan32)

## License

[Apache2.0 © Ryan-2022](./LICENSE)
