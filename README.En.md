<div align="center">
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**English** | [中文](./README.md)
---
[![Go-Zero](https://img.shields.io/badge/Go--Zero-v1.7.4-brightgreen.svg)](https://go-zero.dev/)
[![Ent](https://img.shields.io/badge/Ent-v0.14.1-blue.svg)](https://entgo.io/)
[![Casbin](https://img.shields.io/badge/Casbin-v2.99.0-orange.svg)](https://github.com/casbin/casbin)
[![Release](https://img.shields.io/badge/Release-v1.6.0-green.svg)](https://github.com/suyuan32/simple-admin-core/releases)
[![License: Apache2.0](https://img.shields.io/badge/License-Apache2.0-yellow.svg)](https://opensource.org/licenses/MIT)
![公众号](https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E5%87%A0%E9%A2%97%E9%85%A5-blue)
![注意](https://img.shields.io/badge/%E6%B3%A8%E6%84%8F-%E5%85%B3%E6%B3%A8%E5%85%AC%E4%BC%97%E5%8F%B7%E5%8A%A0%E5%85%A5%E5%BE%AE%E4%BF%A1%E7%BE%A4-blue)

## Introduction

**Simple Admin** is a powerful microservice framework for large management systems. It is based on **Go Zero** and
supports several advanced features. It provides a complete set of functionalities for **user management, permission
management, role management, menu management, log management, and configuration management**. Additionally, it offers
features like **multilingual support**. Simple Admin is suitable for both small and large enterprises to quickly build
distributed backend management systems.

Here are some key points about Simple Admin:

- **Development and Deployment Tools**: Simple Admin comes with robust development and deployment tools. It is
  well-suited for high-concurrency, highly reliable, and complex business scenarios. The project supports easy one-click
  upgrades and provides comprehensive documentation, videos, and examples to help developers get started quickly.

- **Free Basic Modules**: The official release includes **six free basic modules**, covering 80% of common business
  requirements.

- **Module Store**: Simple Admin also offers a **module store**, where you can conveniently purchase and use additional
  modules to enhance your system.

<br>
The 6 free modules are all free for commercial use and are suitable for developer learning, internal use of enterprises, personal project use, etc. Everyone is welcome to use and feedback problems, and we will continue to update and maintain them.

<br>

> ✨ We accept outsourcing. Please contact us via email if you have any outsourcing needs.

## Official Tutorial

### [Simple Admin](https://www.youtube.com/@yuansu5197)

## New upgraded UI (Vben5), greatly optimized experience! [Visit](https://github.com/suyuan32/simple-admin-vben5-ui)

### Online Preview

#### [Free Edition Online Preview](https://vben5-preview.ryansu.tech/)

- Account **admin**
- Password **simple-admin**

#### [Multi-tenant Edition Online Preview](https://tenant-preview.ryansu.tech/)

- Administrator Tenant Account
  - Enterprise: **admin**
  - Account: **admin**
  - Password: **simple-admin**

- Tenant Account
  - Enterprise: **测试企业**
  - Account: **admin**
  - Password: **simple-admin**

> Read-only, cannot be modified or registered

## Feature

- **State of The Art Development**：Use latest back-end technology development such as ent, go-zero, casbin
- **Fully support go-swagger**: Write comment in api file and generate swagger doc easily
- **Error handling**: Handle error messages via one module
- **International**：Support different languages show in the front-end via put locale path in the message
- **Service Discover**: Use k8s endpoints to do service discovery and load-balance
- **Authority** Manage authority via Casbin, based on RBAC
- **Code Generation**: Built-in three-terminal Web, API, RPC code generation
- **Multiple extensions**: Provides a variety of extensions and has a very simple access function
- **Other** Builtin concurrency control, adaptive circuit breaker, adaptive load shedding, auto-trigger, auto recover
- **ORM**: Powerful `Ent` supported

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
| Backend UI  | Backend UI          | [Backend UI](https://github.com/suyuan32/simple-admin-vben5-ui)           |
| FMS         | File Management     | [File](https://github.com/suyuan32/simple-admin-file)                     |
| Job         | Scheduled Task      | [Job](https://github.com/suyuan32/simple-admin-job)                       |
| MMS         | Member Management   | [Member](https://github.com/suyuan32/simple-admin-member-api)             |
| MCMS        | Message Center      | [Message Center](https://github.com/suyuan32/simple-admin-message-center) |

## Member Exclusive Modules

| Module Name          | Module Introduction               |
|----------------------|-----------------------------------|
| CMS                  | Content Management Module         |
| Simple-Uni           | Mini Program Development Scaffold |
| Simple-Nuxt          | PC Web End Development Scaffold   |
| Core Data Permission | Core data permission version      |

# Community Modules

[Click to view](https://github.com/suyuan32/awesome-simple-admin-module)

# Notice

1. It is forbidden to use Simple Admin to develop websites and applications that violate local laws and regulations
2. Simple Admin does not assume any legal responsibility for websites and applications developed using Simple Admin
3. It is forbidden to resell free or paid module source code

## Project Planning Progress

[RoadMap](https://github.com/suyuan32/simple-admin-core/issues/63)

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

[Apache2.0 © Ryan-2022](./LICENSE)
