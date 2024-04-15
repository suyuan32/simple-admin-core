<div align="center">
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**English** | [中文](./README.md)
---
[![Go-Zero](https://img.shields.io/badge/Go--Zero-v1.6.4-brightgreen.svg)](https://go-zero.dev/)
[![Vben Admin](https://img.shields.io/badge/Vben%20Admin-v2.10.1-yellow.svg)](https://vvbin.cn/doc-next/)
[![Ent](https://img.shields.io/badge/Ent-v0.13.1-blue.svg)](https://entgo.io/)
[![Casbin](https://img.shields.io/badge/Casbin-v2.86.1-orange.svg)](https://github.com/casbin/casbin)
[![Release](https://img.shields.io/badge/Release-v1.4.0-green.svg)](https://github.com/suyuan32/simple-admin-core/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![SimpleAdmin](https://dcbadge.vercel.app/api/server/NDED5p2hJk)](https://discord.gg/NDED5p2hJk)
![公众号](https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E5%87%A0%E9%A2%97%E9%85%A5-blue)
![注意](https://img.shields.io/badge/%E6%B3%A8%E6%84%8F-%E5%85%B3%E6%B3%A8%E5%85%AC%E4%BC%97%E5%8F%B7%E5%8A%A0%E5%85%A5%E5%BE%AE%E4%BF%A1%E7%BE%A4-blue)

## Introduction

Simple Admin is an out-of-the-box distributed microservice back-end management system developed based on go-zero, which
provides rich functions for the development of medium and large back-ends, supports three-terminal code generation.
The official comes with a variety of extensions to help small and medium-sized enterprises quickly go to the cloud and
iterate quickly. Suitable for microservice learning and commercial use, open source and free.

## Official Tutorial

### [Simple Admin](https://www.youtube.com/@yuansu5197)

## [Goctls](https://github.com/suyuan32/goctls)

Based on the enhanced version of go zero, it provides a lot of optimizations for simple admin, has a lot of additional
code generation functions,
fully supports ent, and easily realizes three-terminal code generation, making development easier.

## [Doge](https://github.com/suyuan32/doge)

Doge is Simple Admin's command line tool for module download and deployment. It provides module source code download,
module docker, k8s deployment, server maintenance and other functions. Users can upload their own paid modules to earn
revenue, and now 10+ modules are included.

> [Module Store](https://doge.ryansu.tech/store/index)

## Feature

- **State of The Art Development**：Use latest back-end technology development such as ent, go-zero, casbin
- **Fully support go-swagger**: Write comment in api file and generate swagger doc easily
- **Error handling**: Handle error messages via one module
- **International**：support different languages show in the front-end via put locale path in the message
- **Service Discover**: use k8s endpoints to do service discovery and load-balance
- **Authority** Manage authority via Casbin, based on RBAC
- **Code Generation**: Built-in three-terminal Web, API, RPC code generation
- **Multiple extensions**: Provides a variety of extensions and has a very simple access function
- **Other** builtin concurrency control, adaptive circuit breaker, adaptive load shedding, auto-trigger, auto recover

## Supported functions

- User management: manage system user data
- Department management: manage the department to which you belong
- Post management: configure the positions held by system users
- Menu management: configure system menus, tree display
- Role management: manage role permissions, support multiple roles
- Dictionary management: maintain data dictionary, convenient for front-end use
- Interface document: automatically generate related api interface documents according to business code
- Code generation: automatically generate CRUD code, quickly generate custom logic
- Token management: manage token status, support blacklisting token

## Permanently Free Official Modules

| Module Name | Module Introduction | Module Address                                                            |
|-------------|---------------------|---------------------------------------------------------------------------|
| Core        | Core Module         | [Core](https://github.com/suyuan32/simple-admin-core)                     |
| Backend UI  | Backend Interface   | [Backend UI](https://github.com/suyuan32/simple-admin-backend-ui)         |
| FMS         | File Management     | [File](https://github.com/suyuan32/simple-admin-file)                     |
| Job         | Scheduled Task      | [Job](https://github.com/suyuan32/simple-admin-job)                       |
| MMS         | Member Management   | [Member](https://github.com/suyuan32/simple-admin-member-api)             |
| MCMS        | Message Center      | [Message Center](https://github.com/suyuan32/simple-admin-message-center) |

## Member Exclusive Modules

| Module Name | Module Introduction               |
|-------------|-----------------------------------| 
| CMS         | Content Management Module         |
| Simple-Uni  | Mini Program Development Scaffold |
| Simple-Nuxt | PC Web End Development Scaffold   |

# Community Modules

[Click to view](https://github.com/suyuan32/awesome-simple-admin-module)

## Project Planning Progress

[RoadMap](https://github.com/suyuan32/simple-admin-core/issues/63)

## Preview

### Online preview

[Online Preview](http://101.132.124.135:8080/)
Account:   admin 
Password:  simple-admin
#### Read Only, cannot register and modify

![pic](https://i.postimg.cc/qqPNR02x/register-zh-cn.png)
![pic](https://i.postimg.cc/PxczkCr6/dashboard-zh-cn.png)

## Documentation

### [Simple Admin Document](https://doc.ryansu.tech)

- go-zero
  [Document](https://go-zero.dev/)
- ant-design-vue [Document](https://antdv.com/components/overview)

## Preparation
- [Golang](http://go.dev/) and [git](https://git-scm.com/) - Project development environment
- [Ent](https://entgo.io/docs/getting-started) - Ent
- [Mysql](https://www.mysql.com/) - Familiar with mysql database
- [GORM](https://gorm.io/) - Familiar with GORM apis
- [Casbin](https://casbin.org/) - Familiar with Casbin apis
- [Go-swagger](https://goswagger.io/) - Go-swagger document generation


## How to contribute

You are very welcome to join！[Raise an issue](https://github.com/suyuan32/simple-admin-core/issues/new) Or submit a Pull Request。

**Pull Request:**

1. Fork code!
2. Create your own branch: `git checkout -b feat/xxxx`
3. Submit your changes: `git commit -am 'feat(function): add xxxxx'`
4. Push your branch: `git push origin feat/xxxx`
5. submit`pull request`

## Git Contribution submission specification

- reference [vue](https://github.com/vuejs/vue/blob/dev/.github/COMMIT_CONVENTION.md) specification ([Angular](https://github.com/conventional-changelog/conventional-changelog/tree/master/packages/conventional-changelog-angular))

    - `feat` Add new features
    - `fix` Fix the problem/BUG
    - `style` The code style is related and does not affect the running result
    - `perf` Optimization/performance improvement
    - `refactor` Refactor
    - `revert` Undo edit
    - `test` Test related
    - `docs` Documentation/notes
    - `chore` Dependency update/scaffolding configuration modification etc.
    - `workflow` Workflow improvements
    - `ci` Continuous integration
    - `types` Type definition file changes
    - `wip` In development

## Change Log

[CHANGELOG](./CHANGELOG.md)

# Community

> [Discard](https://discord.gg/NDED5p2hJk)

> [Discussion](https://github.com/suyuan32/simple-admin-core/discussions)

## Stars

[![Star History Chart](https://api.star-history.com/svg?repos=suyuan32/simple-admin-core&type=Date)](https://github.com/suyuan32/simple-admin-core)


## Maintainer

[@Ryan Su](https://github.com/suyuan32)

## License

[MIT © Ryan-2022](./LICENSE)
