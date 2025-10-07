<div align="center"> 
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin Vben5 UI</h1>
</div>

**中文** | [English](./README.En.md)

## 简介

Simple Admin UI 是基于 vue-vben-admin v5 二次开发的为 Simple Admin 专门开发的后台管理界面，基于 Vue3 和 TypeScript 开发， 提供后台错误统一处理，国际化等功能，本项目完全免费，可用于学习和商用

## 中文文档地址 [点击查看](https://doc.vben.pro/)

### 在线预览

#### [免费版在线预览](https://preview.ryansu.tech/)

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

## 特性

- **最新技术栈**：基于 Vue3/vite5 等最新技术开发
- **TypeScript**: 采用 Typescript 的语言
- **主题**：可配置的主题
- **国际化**：内置完善的国际化方案
- **Mock 数据** 内置 Mock 数据测试方案
- **权限** 支持动态路由权限
- **组件** 二次封装了多个常用的组件
- **Remeda**: 使用 remeda 作为数据处理工具

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

## 项目规划进度

[RoadMap](https://github.com/suyuan32/simple-admin-core/issues/63)

## 预览

![pic](https://i.postimg.cc/qqPNR02x/register-zh-cn.png) ![pic](https://i.postimg.cc/PxczkCr6/dashboard-zh-cn.png)

[更多预览](https://suyuan32.github.io/simple-admin-core/#/simple-admin/zh-cn/docs/screenshot)

## 文档

[文档](https://doc.vben.pro/)

- ant-design-vue [地址](https://antdv.com/components/overview)

## 准备

- [node](http://nodejs.org/) 和 [git](https://git-scm.com/) -项目开发环境
- [Vite](https://vitejs.dev/) - 熟悉 vite 特性
- [Vue3](https://v3.vuejs.org/) - 熟悉 Vue 基础语法
- [TypeScript](https://www.typescriptlang.org/) - 熟悉`TypeScript`基本语法
- [Es6+](http://es6.ruanyifeng.com/) - 熟悉 es6 基本语法
- [Vue-Router-Next](https://next.router.vuejs.org/) - 熟悉 vue-router 基本使用
- [Ant-Design-Vue](https://2x.antdv.com/docs/vue/introduce-cn/) - ui 基本使用
- [Mock.js](https://github.com/nuysoft/Mock) - mockjs 基本语法

## 快速开始

[快速开始文档](https://doc.ryansu.pro/zh/guide/basic-config/env_setting.html)

## 更新日志

[CHANGELOG](./CHANGELOG.md)

## 项目地址

- [Simple-Admin-ui](https://github.com/suyuan32/Simple-Admin-ui)
- [Simple-Admin](https://github.com/suyuan32/Simple-Admin)

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/suyuan32/Simple-Admin-ui/issues/new/choose) 或者提交一个 Pull Request。

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

## 浏览器支持

本地开发推荐使用`Chrome 80+` 浏览器

支持现代浏览器, 不支持 IE

| [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png" alt=" Edge" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>IE | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png" alt=" Edge" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Edge | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/firefox/firefox_48x48.png" alt="Firefox" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Firefox | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/chrome/chrome_48x48.png" alt="Chrome" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Chrome | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/safari/safari_48x48.png" alt="Safari" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Safari |
| :-: | :-: | :-: | :-: | :-: |
| not support | last 2 versions | last 2 versions | last 2 versions | last 2 versions |

## 相关仓库

如果这些插件对你有帮助，可以给一个 star 支持下

- [vite-plugin-mock](https://github.com/anncwb/vite-plugin-mock) - 用于本地及开发环境数据 mock
- [vite-plugin-html](https://github.com/anncwb/vite-plugin-html) - 用于 html 模版转换及压缩
- [vite-plugin-style-import](https://github.com/anncwb/vite-plugin-style-import) - 用于组件库样式按需引入
- [vite-plugin-theme](https://github.com/anncwb/vite-plugin-theme) - 用于在线切换主题色等颜色相关配置
- [vite-plugin-imagemin](https://github.com/anncwb/vite-plugin-imagemin) - 用于打包压缩图片资源
- [vite-plugin-compression](https://github.com/anncwb/vite-plugin-compression) - 用于打包输出.gz|.brotil 文件
- [vite-plugin-svg-icons](https://github.com/anncwb/vite-plugin-svg-icons) - 用于快速生成 svg 雪碧图

## 维护者

[@Ryan Su](https://github.com/suyuan32)

## License

[Apache2.0 © Ryan-present](.originallicense/LICENSE)
