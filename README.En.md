<div align="center">
<img src="https://s1.imagehub.cc/images/2022/08/12/logo_512.png" width="300px" height="300px"/>
<h1>Simple Admin</h1>
</div>

**English** | [中文](./README.md)

## Introduction

Simple Admin is a powerful microservice framework for basic management. 
It is based on go-zero and supports several advanced features. 
It can help you to develop a microservice back-end management core in a short time.

## Feature

- **State of The Art Development**：Use latest back-end technology development such as gorm, go-zero, casbin
- **Fully support go-swagger**: Write comment in api file and generate swagger doc easily
- **Error handling**: Handle error messages via one module
- **International**：support different languages show in the front-end via put locale path in the message 
- **Service Discover** builtin service discovery, load balancing
- **Authority** Manage authority via casbin
- **Other** builtin concurrency control, adaptive circuit breaker, adaptive load shedding, auto-trigger, auto recover

## Current Progress

| Module             | Status   |
|--------------------|----------|
| Login and Register | finished |
| Menu Management    | finished |
| Role Management    | finished |
| Role Authority     | finished |
| User Management    | finished |
| Operation log      | finished |



## Documentation

Native document will come soon.

- go-zero
  [Document](https://go-zero.dev/)
- ant-design-vue [Document](https://antdv.com/components/overview)

## Preparation
- [Golang](http://go.dev/) and [git](https://git-scm.com/) - Project development environment
- [Mysql](https://www.mysql.com/) - Familiar with mysql database
- [GORM](https://gorm.io/) - Familiar with GORM apis
- [Casbin](https://casbin.org/) - Familiar with Casbin apis

## Install and use

- Get the project code

```bash
git clone https://github.com/suyuan32/Simple-Admin.git
```

- Installation dependencies

```bash
cd Simple-Admin/core

go mod tidy

```

- run

```bash
# run core api
cd api 
go run core.go -f etc/core.yaml

# run core rpc
cd rpc
go run core.go -f etc/core.yaml
```

- build

```bash
go build -o core core.go
```

## Change Log

[CHANGELOG](./CHANGELOG.zh_CN.md)

## Project

- [Simple-Admin-ui](https://github.com/suyuan32/Simple-Admin-ui)
- [Simple-Admin](https://github.com/suyuan32/Simple-Admin)

## How to contribute

You are very welcome to join！[Raise an issue](https://github.com/suyuan32/Simple-Admin/issues/new/choose) Or submit a Pull Request。

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

## Maintainer

[@Ryan Su](https://github.com/suyuan32)

## License

[MIT © Ryan-2022](./LICENSE)
