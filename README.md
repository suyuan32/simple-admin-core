<div align="center">
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**中文** | [English](./README.En.md) 
---
[![Go-Zero](https://img.shields.io/badge/Go--Zero-v1.6.5-brightgreen.svg)](https://go-zero.dev/)
[![Vben Admin](https://img.shields.io/badge/Vben%20Admin-v2.11.4-yellow.svg)](https://vvbin.cn/doc-next/)
[![Ent](https://img.shields.io/badge/Ent-v0.13.1-blue.svg)](https://entgo.io/)
[![Casbin](https://img.shields.io/badge/Casbin-v2.89.0-orange.svg)](https://github.com/casbin/casbin)
[![Release](https://img.shields.io/badge/Release-v1.4.2-green.svg)](https://github.com/suyuan32/simple-admin-core/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![公众号](https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E5%87%A0%E9%A2%97%E9%85%A5-blue)
![注意](https://img.shields.io/badge/%E6%B3%A8%E6%84%8F-%E5%85%B3%E6%B3%A8%E5%85%AC%E4%BC%97%E5%8F%B7%E5%8A%A0%E5%85%A5%E5%BE%AE%E4%BF%A1%E7%BE%A4-blue)

## 简介

Simple Admin 是一个强大的、易扩展的后台管理系统，基于 Go-Zero、Vben Admin、Ent、Casbin
等开源项目构建，提供了完整的用户管理、权限管理、角色管理、菜单管理、日志管理、配置管理等功能，支持多语言等特性，适用于小型或大型企业快速搭建分布式后台管理系统。 <br><br>
Simple Admin 有完善的开发部署工具， 十分适合高并发、高可靠、复杂的业务场景，项目可以方便地一键升级，提供了完善的文档视频和示例，让开发者可以快速上手，快速开发。官方提供了
6 大免费基础模块，可以满足 80 % 的常用业务需求。同时提供了模块商店，可以方便的购买和使用更多的模块。<br><br>
6大免费模块均可免费商用，适合开发者学习、企业内部使用、个人项目使用等，欢迎大家使用和反馈问题，我们会持续更新和维护。

## [Goctls](https://github.com/suyuan32/goctls)

基于 go zero 的加强版工具，针对 simple admin 提供了大量优化，具有大量额外的代码生成功能，全面支持ent，轻松实现三端代码生成，使开发变得简单。

## [Doge](https://github.com/suyuan32/doge)

Doge 是 Simple Admin 的模块下载部署的命令行工具，提供模块源码下载，模块 docker , k8s 部署，服务器维护等功能。用户可以上传自己的付费模块获取收益，现已收录
10 + 模块。

> [模块商店](https://doge.ryansu.tech/store/index)

## 相关教程

> [Bilibili 视频教程](https://space.bilibili.com/9872669/channel/series) \
> 关注微信公众号 - 几颗酥 获取更多教程

#### [Ent 中文文档](https://suyuan32.github.io/ent-chinese-doc/#/zh-cn/getting-started)

## 特性

- **最新技术栈**：使用 ent, casbin, kafka 等前沿技术开发
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

## 永久免费的官方模块

| 模块名称       | 模块介绍 | 模块地址                                                                      |
|------------|------|---------------------------------------------------------------------------|
| Core       | 核心模块 | [Core](https://github.com/suyuan32/simple-admin-core)                     |
| Backend UI | 后端界面 | [Backend UI](https://github.com/suyuan32/simple-admin-backend-ui)         |
| FMS        | 文件管理 | [File](https://github.com/suyuan32/simple-admin-file)                     |
| Job        | 定时任务 | [Job](https://github.com/suyuan32/simple-admin-job)                       |
| MMS        | 会员管理 | [Member](https://github.com/suyuan32/simple-admin-member-api)             |
| MCMS       | 消息中心 | [Message Center](https://github.com/suyuan32/simple-admin-message-center) |

## 会员专属的模块

| 模块名称        | 模块介绍        |
|-------------|-------------| 
| CMS         | 内容管理模块      |
| Simple-Uni  | 小程序开发脚手架    |
| Simple-Nuxt | PC 网页端开发脚手架 |

# 社区模块

[点击查看](https://github.com/suyuan32/awesome-simple-admin-module)

## 项目规划进度

[RoadMap](https://github.com/suyuan32/simple-admin-core/issues/63)

## 预览

### 在线预览

[在线预览](http://101.132.124.135:8080/)
账号 admin
密码 simple-admin
#### 只读，不可修改和注册

![pic](https://i.postimg.cc/qqPNR02x/register-zh-cn.png)
![pic](https://i.postimg.cc/PxczkCr6/dashboard-zh-cn.png)

## 文档

### [Simple Admin 中文文档](https://doc.ryansu.tech/zh)


- go-zero
  [Document](https://go-zero.dev/)
- ant-design-vue [Document](https://antdv.com/components/overview)

## 准备

- [Golang](http://go.dev/) and [git](https://git-scm.com/) - Go 语言
- [Ent](https://entgo.io/docs/getting-started) - Ent
- [Mysql](https://www.mysql.com/) - Mysql数据库
- [GORM](https://gorm.io/) - GORM 数据库ORM组件
- [Casbin](https://casbin.org/) - 权限管理
- [Go-swagger](https://goswagger.io/) - Go-swagger 文档生成调试

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/suyuan32/simple-admin-core/issues/new) 或者提交一个 Pull Request。

**Pull Request:**

1. Fork 代码!
2. 创建自己的分支: `git checkout -b feat/xxxx`
3. 提交你的修改: `git commit -am 'feat(function): add xxxxx'`
4. 推送您的分支: `git push origin feat/xxxx`
5. 提交`pull request`

## Git 贡献提交规范

- 参考 [vue](https://github.com/vuejs/vue/blob/dev/.github/COMMIT_CONVENTION.md) 规范 ([Angular](https://github.com/conventional-changelog/conventional-changelog/tree/master/packages/conventional-changelog-angular))

    - `feat` 增加新功能
    - `fix` 修复问题/BUG
    - `style` 代码风格相关无影响运行结果的
    - `perf` 优化/性能提升
    - `refactor` 重构
    - `revert` 撤销修改
    - `test` 测试相关
    - `docs` 文档/注释
    - `chore` 依赖更新/脚手架配置修改等
    - `workflow` 工作流改进
    - `ci` 持续集成
    - `types` 类型定义文件更改
    - `wip` 开发中

## 更新日志

[CHANGELOG](./CHANGELOG.md)

## 交流群

> [Discord](https://discord.gg/NDED5p2hJk)

> [论坛](https://github.com/suyuan32/simple-admin-core/discussions)

> 关注公众号 《几颗酥》 加入微信群

## Stars

[![Star History Chart](https://api.star-history.com/svg?repos=suyuan32/simple-admin-core&type=Date)](https://github.com/suyuan32/simple-admin-core)

## 维护者

[@Ryan Su](https://github.com/suyuan32)

## License

[MIT © Ryan-2022](./LICENSE)
