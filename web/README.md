<div align="center">
<img src="https://i.postimg.cc/nh8mVKkf/logo.png" width="300px" height="300px"/>
<h1>Simple Admin Vben5 UI</h1>
</div>

**English** | [中文](./README.zh-CN.md)

## Introduction

Simple Admin Vben5 UI is a modern UI for Simple Admin. It is based on vue-vben-admin and supports several advanced features. It can help you developing a distributed backend management system in a short time.

## Document [Click here](https://doc.vben.pro/)

### Online Preview

#### [Free Edition Online Preview](https://preview.ryansu.tech/)

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

- **Latest technology stack**: Developed based on the latest technologies such as Vue3/vite5
- **TypeScript**: Uses Typescript language
- **Theme**: Configurable themes
- **Internationalization**: Built-in complete internationalization solution
- **Mock data** Built-in Mock data testing solution
- **Permission** Supports dynamic routing permissions
- **Component** Secondary encapsulation of multiple commonly used components
- **Remeda**: Use remeda as a data processing tool

## Support functions

- User management: manage system user data
- Department management: manage the department to which you belong
- Post management: configure the positions held by system users
- Menu management: configure system menus, tree display
- Role management: manage role permissions, support multiple roles
- Dictionary management: maintain data dictionary, convenient for front-end use
- Interface document: automatically generate related api interface documents according to business code
- Code generation: automatically generate CRUD code, quickly generate custom logic
- Token management: manage token status, support blacklisting token

## Project Planning Progress

[RoadMap](https://github.com/suyuan32/simple-admin-core/issues/63)

### Preview

![pic](https://i.postimg.cc/qqPNR02x/register-zh-cn.png) ![pic](https://i.postimg.cc/PxczkCr6/dashboard-zh-cn.png)

[More](https://suyuan32.github.io/simple-admin-core/#/simple-admin/zh-cn/docs/screenshot)

## Documentation

[Simple Admin Documentation](https://doc.vben.pro/)

- ant-design-vue [Document](https://antdv.com/components/overview)

## Preparation

- [node](http://nodejs.org/) and [git](https://git-scm.com/) - Project development environment
- [Vite](https://vitejs.dev/) - Familiar with vite features
- [Vue3](https://v3.vuejs.org/) - Familiar with Vue basic syntax
- [TypeScript](https://www.typescriptlang.org/) - Familiar with the basic syntax of `TypeScript`
- [Es6+](http://es6.ruanyifeng.com/) - Familiar with es6 basic syntax
- [Vue-Router-Next](https://next.router.vuejs.org/) - Familiar with the basic use of vue-router
- [Ant-Design-Vue](https://2x.antdv.com/docs/vue/introduce-cn/) - ui basic use
- [Mock.js](https://github.com/nuysoft/Mock) - mockjs basic syntax

## Quick Start

[Quick Start Document](https://doc.ryansu.pro/en/guide/basic-config/env_setting.html)

## Change Log

[CHANGELOG](./CHANGELOG.md)

## Project

- [Simple-Admin-ui](https://github.com/suyuan32/Simple-Admin-ui)
- [Simple-Admin](https://github.com/suyuan32/Simple-Admin)

## How to contribute

You are very welcome to join！[Raise an issue](https://github.com/suyuan32/Simple-Admin-ui/issues/new/choose) Or submit a Pull Request。

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

## Related warehouse

If these plugins are helpful to you, you can give a star support

- [vite-plugin-mock](https://github.com/anncwb/vite-plugin-mock) - Used for local and development environment data mock
- [vite-plugin-html](https://github.com/anncwb/vite-plugin-html) - Used for html template conversion and compression
- [vite-plugin-style-import](https://github.com/anncwb/vite-plugin-style-import) - Used for component library style introduction on demand
- [vite-plugin-theme](https://github.com/anncwb/vite-plugin-theme) - Used for online switching of theme colors and other color-related configurations
- [vite-plugin-imagemin](https://github.com/anncwb/vite-plugin-imagemin) - Used to pack compressed image resources
- [vite-plugin-compression](https://github.com/anncwb/vite-plugin-compression) - Used to pack input .gz|.brotil files
- [vite-plugin-svg-icons](https://github.com/anncwb/vite-plugin-svg-icons) - Used to quickly generate svg sprite

## Browser support

The `Chrome 80+` browser is recommended for local development

Support modern browsers, not IE

| [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png" alt=" Edge" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>IE | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png" alt=" Edge" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Edge | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/firefox/firefox_48x48.png" alt="Firefox" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Firefox | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/chrome/chrome_48x48.png" alt="Chrome" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Chrome | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/safari/safari_48x48.png" alt="Safari" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)</br>Safari |
| :-: | :-: | :-: | :-: | :-: |
| not support | last 2 versions | last 2 versions | last 2 versions | last 2 versions |

## Maintainer

[@Ryan Su](https://github.com/suyuan32)

## License

[Apache2.0 © Ryan-2022](.originallicense/LICENSE)
