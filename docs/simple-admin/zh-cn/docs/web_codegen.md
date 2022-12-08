# Vben Web 端代码生成

> 首先确认你安装了以下软件:
- simple-admin-tool (goctls) v0.1.1-beta +


## 创建 example 代码

```shell
goctls frontend vben --apiFile=/home/ryan/GolandProjects/simple-admin-example-api/desc/student.api --o=./ --folderName=example --prefix=example-api --subFolder=
```

### 参数介绍

| 参数         | 介绍         | 使用方法                                                               |
|------------|------------|--------------------------------------------------------------------|
| apiFile    | api文件的绝对路径 | 填入api文件的绝对路径，如上面的 student.api                                      |
| o          | 输出路径       | 输入 simple admin backend ui 目录                                      |
| folderName | 文件夹名称      | core服务是 sys, 示例项目是 example                                         |
| prefix     | 请求前缀       | 请求前缀用于请求转发，如sys-api, 示例项目为example-api, 需要修改env.development,添加proxy |
| subFolder  | 子目录        | 用于在views下创建子目录，如sys有user,token等子目录                                 |


> 执行命令后会生成下面的代码

- `src/api/example/student.ts src/api/example/model/student.ts`    API声明和请求代码
- `src/locales/lang/en/example.ts src/locales/lang/en/example.ts`  国际化代码
- `src/views/example/*` 视图代码

> 生成代码后还需要做的事

- 修改 env.development 和 deploy/default.conf 添加新的服务地址
- 修改国际化代码，优化中文翻译
- 修改views 中的 *.data.ts 完善中文翻译，行列以及提交表格中的字段名需要自行翻译添加到国际化代码中使用 `locales/lang/example.ts`

> 示例地址 https://github.com/suyuan32/simple-admin-backend-ui/tree/example-code-gen