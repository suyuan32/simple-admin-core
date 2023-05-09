<div align="center">
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**English** | [中文](./README.md)
---
[![Go-Zero](https://img.shields.io/badge/Go--Zero-v1.5.2-brightgreen.svg)](https://go-zero.dev/)
[![Vben Admin](https://img.shields.io/badge/Vben%20Admin-v2.10.0-yellow.svg)](https://vvbin.cn/doc-next/)
[![Ent](https://img.shields.io/badge/Ent-v0.12.3-blue.svg)](https://entgo.io/)
[![Casbin](https://img.shields.io/badge/Casbin-v2.69.0-orange.svg)](https://github.com/casbin/casbin)
[![Release](https://img.shields.io/badge/Release-v1.0.4-green.svg)](https://github.com/suyuan32/simple-admin-core/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![SimpleAdmin](https://dcbadge.vercel.app/api/server/NDED5p2hJk)](https://discord.gg/NDED5p2hJk)
![QQ群](https://img.shields.io/badge/QQ%E7%BE%A4-801043319-blue)

## Introduction

Simple Admin is an out-of-the-box distributed microservice back-end management system developed based on go-zero, which
provides rich functions for the development of medium and large back-ends, supports three-terminal code generation.
The official comes with a variety of extensions to help small and medium-sized enterprises quickly go to the cloud and
iterate quickly. Suitable for microservice learning and commercial use, open source and free.

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

## Support functions

- User management: The user is the system operator, and this function mainly completes the system user configuration.
- Department management: Configure the system organization (company, department, group), and the tree structure shows the support data permissions.
- Position management: configure the positions that system users belong to.
- Menu management: configure system menu, operation authority, button authority identification, interface authority, etc.
- Role management: role menu permission assignment, role setting, data range permission division by organization.
- Dictionary management: maintain some relatively fixed data frequently used in the system.
- Operation log: system normal operation log record and query; system abnormal information log record and query.
- Member management: manage registered member information
- Interface documents: Automatically generate relevant API interface documents based on business codes.
- Code generation: Generate corresponding additions, deletions, modifications, and queries based on the data table structure
- Service monitoring: View some basic information about servers

## Project Planning Progress

[RoadMap](https://github.com/suyuan32/simple-admin-core/issues/63)

## Preview

### Online preview
[Online Preview](http://101.132.124.135/#/dashboard)
Account:   admin 
Password:  simple-admin
#### Read Only, cannot register and modify

![pic](https://i.postimg.cc/qqPNR02x/register-zh-cn.png)
![pic](https://i.postimg.cc/PxczkCr6/dashboard-zh-cn.png)

## Documentation

[Simple Admin Document](https://doc.ryansu.pro)

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

## Quick Start

[Quick Start Document](https://doc.ryansu.pro/en/guide/basic-config/env_setting.html)

## Change Log

[CHANGELOG](./CHANGELOG.zh_CN.md)

## Relative Project

- [Simple Admin](https://github.com/suyuan32/simple-admin-core)
- [Simple Admin Backend UI](https://github.com/suyuan32/simple-admin-backend-ui)

## Optional Components

- [File Management](https://github.com/suyuan32/simple-admin-file)
- [Scheduled Task](https://github.com/suyuan32/simple-admin-job)
- [Member Management](https://github.com/suyuan32/simple-admin-member-api)


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

# Community

> [Discard](https://discord.gg/NDED5p2hJk)

> [Discussion](https://github.com/suyuan32/simple-admin-core/discussions)

## Stars

[![Star History Chart](https://api.star-history.com/svg?repos=suyuan32/simple-admin-core&type=Date)](https://github.com/suyuan32/simple-admin-core)


## Maintainer

[@Ryan Su](https://github.com/suyuan32)

## License

[MIT © Ryan-2022](./LICENSE)
