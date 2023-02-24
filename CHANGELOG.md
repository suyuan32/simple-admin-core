
<a name="v0.2.5"></a>
## [v0.2.5](https://github.com/suyuan32/simple-admin-core/compare/v0.2.4...v0.2.5) (2023-02-23)

### Chore

* update tool version
* update uuid package version
* update all in one docker compose

### Docs

* update change log

### Feat

* validator parameter in parse function
* add ldflags to reduce the size of binary file
* Tracing Analysis

### Fix

* remove member translation
* change page default order to desc
* user list roleIds request params

### Refactor

* remove member in core

### Pull Requests

* Merge pull request [#72](https://github.com/suyuan32/simple-admin-core/issues/72) from suyuan32/refator-split-member


<a name="v0.2.4"></a>
## [v0.2.4](https://github.com/suyuan32/simple-admin-core/compare/v0.2.4-beta...v0.2.4) (2023-02-13)

### Chore

* update tool version to v0.2.0

### Feat

* multiple position support
* add enabled support in rpc
* multiple role supports

### Fix

* remove current active menu param
* optimize casbin
* optimize api comment


<a name="v0.2.4-beta"></a>
## [v0.2.4-beta](https://github.com/suyuan32/simple-admin-core/compare/v0.2.3...v0.2.4-beta) (2023-02-09)

### Docs

* update CHANGELOG.md

### Fix

* optimize some bugs

### Refactor

* optimize interface

### Wip

* refactor create or update
* refactor create or update
* refactor create or update

### Pull Requests

* Merge pull request [#68](https://github.com/suyuan32/simple-admin-core/issues/68) from suyuan32/refator-createUpdate


<a name="v0.2.3"></a>
## [v0.2.3](https://github.com/suyuan32/simple-admin-core/compare/v0.2.2...v0.2.3) (2023-02-06)

### Chore

* update tool version and readme

### Feat

* member log in and register
* member rank management
* member init codes  and fix menu create optional bug
* member management
* post management in rpc and api
* post management for ent
* department for users
* home path for user
* dept generating codes
* ent department schema

### Fix

* department status bugs
* menu authority bugs
* status code generation bug

### Refactor

* rename post to position
* rename orderNo to sort

### Wip

* dept generating codes

### Pull Requests

* Merge pull request [#66](https://github.com/suyuan32/simple-admin-core/issues/66) from suyuan32/feat-post
* Merge pull request [#65](https://github.com/suyuan32/simple-admin-core/issues/65) from suyuan32/feat-dept


<a name="v0.2.2"></a>
## [v0.2.2](https://github.com/suyuan32/simple-admin-core/compare/v0.2.1...v0.2.2) (2023-01-29)

### Chore

* update tool version

### Doc

* update readme

### Docs

* update readme

### Feat

* split proto files

### Fix

* update makefile
* repeated bug
* repeated bug
* repeated bug
* DatabaseConf struct Type tag
* update github setting

### Pull Requests

* Merge pull request [#61](https://github.com/suyuan32/simple-admin-core/issues/61) from suyuan32/feat-proto-split
* Merge pull request [#58](https://github.com/suyuan32/simple-admin-core/issues/58) from ch3nnn/fix/config-database


<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/suyuan32/simple-admin-core/compare/v0.2.1-beta...v0.2.1) (2023-01-10)

### Doc

* update readme


<a name="v0.2.1-beta"></a>
## [v0.2.1-beta](https://github.com/suyuan32/simple-admin-core/compare/v0.2.0...v0.2.1-beta) (2023-01-09)

### Chore

* update tool version

### Docs

* update changelog

### Feat

* change user and token id to uuid
* uuid mixin

### Fix

* ErrorCtx and OkJsonCtx
* uuid mixin


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/suyuan32/simple-admin-core/compare/v0.2.0-beta...v0.2.0) (2022-12-29)

### Chore

* update go zero version to 1.4.3

### Docs

* update readme

### Fix

* allow to alter table structure
* translator test
* dashboard menu type
* status validator

### Refactor

* move docs outside the project


<a name="v0.2.0-beta"></a>
## [v0.2.0-beta](https://github.com/suyuan32/simple-admin-core/compare/v0.1.9...v0.2.0-beta) (2022-12-17)

### Docs

* update api_example.md
* update ent doc
* update gorm doc
* update CHANGELOG.md

### Fix

* postgres dsn
* pagination template bug [#51](https://github.com/suyuan32/simple-admin-core/issues/51)
* postgres dsn

### Pull Requests

* Merge pull request [#52](https://github.com/suyuan32/simple-admin-core/issues/52) from suyuan32/fix-pagination


<a name="v0.1.9"></a>
## [v0.1.9](https://github.com/suyuan32/simple-admin-core/compare/list...v0.1.9) (2022-12-15)

### Chore

* update tool to v0.1.3
* update tools

### Docs

* group doc
* group doc
* update env_setting.md
* update quickcmd.md
* update CHANGELOG.md

### Feat

* group rpc logic
* gitlab file generating
* batch delete user

### Refactor

* optimize database driver name

### Pull Requests

* Merge pull request [#49](https://github.com/suyuan32/simple-admin-core/issues/49) from suyuan32/group-logic


<a name="list"></a>
## [list](https://github.com/suyuan32/simple-admin-core/compare/v0.1.8...list) (2022-12-10)


<a name="v0.1.8"></a>
## [v0.1.8](https://github.com/suyuan32/simple-admin-core/compare/v0.1.7...v0.1.8) (2022-12-10)

### Docs

* update all in one and simple admin tools docs
* all in one tutorial
* all in one tutorial
* all in one tutorial
* add discussion
* update sidebar
* update readme
* update web code generation tutorial
* update quick_develop_example.md
* update readme and change log
* update statistic code
* update api example parameter

### Feat

* ent no cache driver
* all in one docker-compose file
* add batch delete for token management

### Fix

* force logging out i18n


<a name="v0.1.7"></a>
## [v0.1.7](https://github.com/suyuan32/simple-admin-core/compare/v0.1.6...v0.1.7) (2022-12-03)

### Docs

* update quick_develop_example.md
* api and rpc code generation doc
* feat rpc ent logic tutorials
* update ent doc
* update change log

### Fix

* update tool to v0.1.0
* initialize code
* init api code
* optimize request url
* optimize request url
* job service_context.go

### Pull Requests

* Merge pull request [#46](https://github.com/suyuan32/simple-admin-core/issues/46) from suyuan32/fix-reuse-pkg


<a name="v0.1.6"></a>
## [v0.1.6](https://github.com/suyuan32/simple-admin-core/compare/v0.1.5...v0.1.6) (2022-11-18)

### Docs

* update swagger version
* update file manager doc
* update README.md
* update CHANGELOG.md

### Fix

* block token bug
* validator example

### Perf

* optimize database config

### Refactor

* move errorx.msg to i18n.var
* move middleware to pkg for common usage
* move middleware to pkg for common usage

### Pull Requests

* Merge pull request [#45](https://github.com/suyuan32/simple-admin-core/issues/45) from suyuan32/fix-reuse-pkg


<a name="v0.1.5"></a>
## [v0.1.5](https://github.com/suyuan32/simple-admin-core/compare/v0.1.4...v0.1.5) (2022-11-13)

### Docs

* update error_handling.md

### Fix

* user validator
* update user
* empty translation
* move menu translation to server
* init database error

### Refactor

* response data
* response data

### Wip

* error code optimize
* error code optimize

### Pull Requests

* Merge pull request [#43](https://github.com/suyuan32/simple-admin-core/issues/43) from suyuan32/fix-error-code


<a name="v0.1.4"></a>
## [v0.1.4](https://github.com/suyuan32/simple-admin-core/compare/v0.1.4-beta...v0.1.4) (2022-11-11)

### Chore

* update simple admin tool to 0.0.8

### Docs

* ent
* update Chinese doc for ent
* update Chinese doc for ent

### Feat

* error msg i18n
* cache config

### Fix

* optimize api definition
* use server translate api description

### Wip

* update trans


<a name="v0.1.4-beta"></a>
## [v0.1.4-beta](https://github.com/suyuan32/simple-admin-core/compare/v0.1.3...v0.1.4-beta) (2022-11-07)

### Docs

* update ent doc
* env setting and global vars
* update quick cmd and ent doc
* update read me
* update change log

### Feat

* ent support

### Fix

* remove auto migrate
* duplicate err

### Wip

* ent support
* ent support

### Pull Requests

* Merge pull request [#40](https://github.com/suyuan32/simple-admin-core/issues/40) from suyuan32/ent


<a name="v0.1.3"></a>
## [v0.1.3](https://github.com/suyuan32/simple-admin-core/compare/v0.1.2...v0.1.3) (2022-11-02)

### Chore

* update goctls
* update gorm version

### Docs

* update swagger doc
* update changelog

### Feat

* docker ignore

### Fix

* DictionaryDetail 表结构 [#38](https://github.com/suyuan32/simple-admin-core/issues/38)
* blocking in cron
* api model index
* http status

### Refactor

* optimize project structure
* model path and use simple-models in cron
* rename api_desc directory
* logmessage and message

### Pull Requests

* Merge pull request [#37](https://github.com/suyuan32/simple-admin-core/issues/37) from MUHM/fix/status


<a name="v0.1.2"></a>
## [v0.1.2](https://github.com/suyuan32/simple-admin-core/compare/v0.1.1...v0.1.2) (2022-10-28)

### Chore

* update simple admin tool

### Docs

* update readme
* job schedule doc

### Feat

* job dockerfile
* rocket mq deploy file
* rocket mq deploy file

### Fix

* add optional tag for userprofile request
* add optional tag for userprofile request
* add optional tag for page request
* CreateAt,UpdateAt,DeleteAt 命名不统一 [#29](https://github.com/suyuan32/simple-admin-core/issues/29)
* swagger response type definition wrong [#27](https://github.com/suyuan32/simple-admin-core/issues/27)
* change log message into lowercase

### Perf

* optimizer api generation

### Refactor

* change api file name into snake format
* change rpc file name into snake format

### Wip

* cron and mq

### Pull Requests

* Merge pull request [#34](https://github.com/suyuan32/simple-admin-core/issues/34) from suyuan32/feat-rckmq
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-core/issues/31) from RogueCultivators/fix/list


<a name="v0.1.1"></a>
## [v0.1.1](https://github.com/suyuan32/simple-admin-core/compare/v0.1.0...v0.1.1) (2022-10-21)

### Docs

* jwt black list

### Feat

* force log out
* api service for token management
* rpc service for token management
* jwt management

### Fix

* error judgement in database


<a name="v0.1.0"></a>
## [v0.1.0](https://github.com/suyuan32/simple-admin-core/compare/v0.0.9...v0.1.0) (2022-10-17)

### Docs

* log collecting
* update k8s-deploy
* update prometheus doc
* prometheus deployment
* update file manager deploy doc
* update env setting
* update file manager doc

### Feat

* filebeat deploy
* storing log into persistence volume
* prometheus deployment and fix bugs in other deployment
* prometheus support

### Fix

* update change log

### Pull Requests

* Merge pull request [#20](https://github.com/suyuan32/simple-admin-core/issues/20) from suyuan32/feat-log-collection
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-core/issues/19) from suyuan32/feat-log-collection
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-core/issues/18) from suyuan32/feat-monitor


<a name="v0.0.9"></a>
## [v0.0.9](https://github.com/suyuan32/simple-admin-core/compare/v0.0.8.1...v0.0.9) (2022-10-12)

### Docs

* update english doc
* update chinese doc
* update k8s deployment english doc
* update k8s deployment chinese doc
* k8s deploy
* update index readme

### Feat

* k8s deploy config
* k8s config
* gitlab ci/cd

### Fix

* update go mod

### Revert

* api config
* origin config

### Pull Requests

* Merge pull request [#16](https://github.com/suyuan32/simple-admin-core/issues/16) from suyuan32/feat-k8s


<a name="v0.0.8.1"></a>
## [v0.0.8.1](https://github.com/suyuan32/simple-admin-core/compare/v0.0.8...v0.0.8.1) (2022-10-09)

### Docs

* update readme about k8s support
* changelog and readme

### Fix

* initialize url
* delete provider config when update


<a name="v0.0.8"></a>
## [v0.0.8](https://github.com/suyuan32/simple-admin-core/compare/v0.0.7...v0.0.8) (2022-10-08)

### Docs

* fix nginx config
* add oauth doc
* update change log

### Feat

* oauth api service
* oauth rpc service

### Fix

* oauth initialize bugs and callback bugs
* add oauth menu insert in initialization code
* change create account rules when use oauth log in
* google oauth


<a name="v0.0.7"></a>
## [v0.0.7](https://github.com/suyuan32/simple-admin-core/compare/v0.0.7-beta...v0.0.7) (2022-09-30)

### Fix

* authority tree generation and init database code
* bugs in menu tree generation


<a name="v0.0.7-beta"></a>
## [v0.0.7-beta](https://github.com/suyuan32/simple-admin-core/compare/v0.0.6...v0.0.7-beta) (2022-09-29)

### Docs

* add badges
* read only notice
* deploy in docker
* add consul url
* add consul url
* rpc example service config
* call rpc service
* docker
* docker

### Feat

* dictionary
* dictionary
* online preview
* online preview
* error handling doc
* global variable, GORM and authorization doc
* example doc
* discord link and qq link
* API and RPC example doc
* update changelog

### Fix

* docker compose and rpc doc
* update tool version


<a name="v0.0.6"></a>
## [v0.0.6](https://github.com/suyuan32/simple-admin-core/compare/v0.0.5...v0.0.6) (2022-09-23)

### Feat

* use consul to do service discover and get configuration feat: dockerfile feat: consul doc

### Fix

* update change log


<a name="v0.0.5"></a>
## [v0.0.5](https://github.com/suyuan32/simple-admin-core/compare/v0.0.4...v0.0.5) (2022-09-21)

### Feat

* validator doc
* add translation test
* validator definition fix: swagger doc about conditions
* validator

### Fix

* update simple admin tools version
* add ui build doc
* menu meta setting logic
* use default logger and fix some bugs
* update readme
* update doc
* update readme


<a name="v0.0.4"></a>
## [v0.0.4](https://github.com/suyuan32/simple-admin-core/compare/v0.0.2...v0.0.4) (2022-09-13)

### Feat

* add bilibili link
* web setting doc
* file manager doc
* swagger doc
* example doc

### Fix

* require properties in api doc
* add recommend in doc
* fix the bugs of post request in swagger 修复go-swagger中的post请求参数
* tip for doc
* tip for init
* preview url
* update tool version
* simple admin tool version
* adjust imports for new goctl

### Refactor

* menu properties fix: bug in mysql group


<a name="v0.0.2"></a>
## [v0.0.2](https://github.com/suyuan32/simple-admin-core/compare/v0.0.1...v0.0.2) (2022-09-05)

### Feat

* add video link
* add init url
* doc
* add logger for all logic fix: some bugs in config
* user profile

### Fix

* readme
* optimize readme
* optimize config file
* add user profile init

### Pull Requests

* Merge pull request [#2](https://github.com/suyuan32/simple-admin-core/issues/2) from suyuan32/dev


<a name="v0.0.1"></a>
## v0.0.1 (2022-08-26)

### Feat

* swagger doc
* part swagger
* change password
* initialize database api

### Fix

* some bugs on menu and model
* some bugs on user and captcha
* add error msg
* init database context exceed timeout
* optimize error messages management

### Refactor

* optimize project error handler and initialize
* reconstruct project

### Pull Requests

* Merge pull request [#1](https://github.com/suyuan32/simple-admin-core/issues/1) from suyuan32/dev

