<div align="center">
<img src="https://s1.imagehub.cc/images/2022/08/12/logo_512.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**中文** | [English](./README.En.md)

## Introduction

Simple Admin 是一个开箱即用的分布式微服务后端管理系统，基于go-zero开发，提供丰富的功能如服务发现，权限管理等。
该框架可以帮助您快速开发具有RPC服务的后台管理系统。

## 特性

- **最新技术栈**：使用 gorm, casbin, kafka 等前沿技术开发
- **完全支持go-swagger**: 直接在api文件内编写注释即可直接生成swagger文档
- **统一的错误处理**: 整个系统拥有国际化的统一错误处理
- **国际化**：内置完善的国际化方案
- **服务注册发现** 完善的服务注册发现机制
- **权限** 内置完善的动态路由权限生成方案, 集成RBAC权限控制
- **其他** 流量控制， ES服务

## 当前进度

| 功能   | 进度  |
|------|-----|
| 登录注册 | 已完成 |
| 菜单管理 | 已完成 |
| 角色管理 | 已完成 |
| 角色权限 | 已完成 |
| 用户管理 | 已完成 |
| 操作日志 | 已完成 |
| 服务注册发现 | 已完成 |
| 配置中心 | 已完成 |


## 预览

![pic](https://s1.imagehub.cc/images/2022/09/15/-2022-09-05-21-49-00.png)
![pic](https://s1.imagehub.cc/images/2022/09/15/register_zh_cn.png)
![pic](https://s1.imagehub.cc/images/2022/09/15/add_example_api_authority.png)

[更多预览](https://suyuan32.github.io/simple-admin-core/#/simple-admin/zh-cn/docs/screenshot)


## 文档

[文档](https://suyuan32.github.io/simple-admin-core/)

或者直接本地运行
```shell
cd docs
docsify serve .
```

- go-zero
  [Document](https://go-zero.dev/)
- ant-design-vue [Document](https://antdv.com/components/overview)

## 准备

- [Golang](http://go.dev/) and [git](https://git-scm.com/) - Go 语言
- [Mysql](https://www.mysql.com/) - Mysql数据库
- [GORM](https://gorm.io/) - GORM 数据库ORM组件
- [Casbin](https://casbin.org/) - 权限管理
- [Go-swagger](https://goswagger.io/) - Go-swagger 文档生成调试
- [Consul](https://www.consul.io/docs) - Consul

## 安装使用

- 获取项目代码

```bash
git clone https://github.com/suyuan32/simple-admin-core.git
```

- 安装依赖

```bash
cd simple-admin-core/

go mod tidy
```
- 编辑配置文件 api/etc/core.yaml  rpc/etc/core.yaml
- 
- 运行

```bash
# run core api
cd api 
go run core.go -f etc/core.yaml

# run core rpc
cd rpc
go run core.go -f etc/core.yaml
```

- 打包


```bash
go build -o core core.go
```


## 更新日志

[CHANGELOG](./CHANGELOG.md)

## 项目地址

- [Simple-Admin-ui](https://github.com/suyuan32/Simple-Admin-ui)
- [Simple-Admin](https://github.com/suyuan32/Simple-Admin)

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/suyuan32/Simple-Admin/issues/new/choose) 或者提交一个 Pull Request。

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
## 维护者

[@Ryan Su](https://github.com/suyuan32)

## License

[MIT © Ryan-2022](./LICENSE)
