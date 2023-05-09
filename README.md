<div align="center">
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**中文** | [English](./README.En.md) 
---
[![Go-Zero](https://img.shields.io/badge/Go--Zero-v1.5.2-brightgreen.svg)](https://go-zero.dev/)
[![Vben Admin](https://img.shields.io/badge/Vben%20Admin-v2.10.0-yellow.svg)](https://vvbin.cn/doc-next/)
[![Ent](https://img.shields.io/badge/Ent-v0.12.3-blue.svg)](https://entgo.io/)
[![Casbin](https://img.shields.io/badge/Casbin-v2.69.0-orange.svg)](https://github.com/casbin/casbin)
[![Release](https://img.shields.io/badge/Release-v1.0.4-green.svg)](https://github.com/suyuan32/simple-admin-core/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![SimpleAdmin](https://dcbadge.vercel.app/api/server/NDED5p2hJk)](https://discord.gg/NDED5p2hJk)
![QQ群](https://img.shields.io/badge/QQ%E7%BE%A4-801043319-blue)

## 简介

Simple Admin 是一个开箱即用的分布式微服务后端管理系统，基于go-zero开发，为开发中大型后台提供了丰富的功能，支持三端代码生成。
官方自带多种扩展，助力中小企业快速上云，快速迭代。适合用于微服务学习和商用，开源免费。

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

## 支持功能

- 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
- 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
- 岗位管理：配置系统用户所属担任职务。
- 菜单管理：配置系统菜单，操作权限，按钮权限标识，接口权限等。
- 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
- 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
- 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
- 会员管理：管理注册会员信息
- 接口文档：根据业务代码自动生成相关的api接口文档。
- 代码生成：根据数据表结构生成对应的增删改查相对应业务
- 服务监控：查看一些服务器的基本信息

## 项目规划进度

[RoadMap](https://github.com/suyuan32/simple-admin-core/issues/63)

## 预览

### 在线预览
[在线预览](http://101.132.124.135/#/dashboard)
账号 admin
密码 simple-admin
#### 只读，不可修改和注册

![pic](https://i.postimg.cc/qqPNR02x/register-zh-cn.png)
![pic](https://i.postimg.cc/PxczkCr6/dashboard-zh-cn.png)

## 文档

[Simple Admin 文档](https://doc.ryansu.pro)


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

## 快速开始

[快速开始文档](https://doc.ryansu.pro/zh/guide/basic-config/env_setting.html)

## 更新日志

[CHANGELOG](./CHANGELOG.md)

## 相关项目

- [Simple Admin](https://github.com/suyuan32/simple-admin-core)
- [Simple Admin 后端界面](https://github.com/suyuan32/simple-admin-backend-ui)

## 可选组件

- [文件管理](https://github.com/suyuan32/simple-admin-file)
- [定时任务](https://github.com/suyuan32/simple-admin-job)
- [会员管理](https://github.com/suyuan32/simple-admin-member-api)

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

## 交流群
> QQ 801043319 

>[Discord](https://discord.gg/NDED5p2hJk)

> [论坛](https://github.com/suyuan32/simple-admin-core/discussions)

> 微信群

<div align="center">
<img src="https://doc.ryansu.pro/assets/contact.png" width="250px" height="320px"/>
</div>

## Stars

[![Star History Chart](https://api.star-history.com/svg?repos=suyuan32/simple-admin-core&type=Date)](https://github.com/suyuan32/simple-admin-core)

## 维护者

[@Ryan Su](https://github.com/suyuan32)

## License

[MIT © Ryan-2022](./LICENSE)
