## 前端设置

### 主要修改 env 文件


> .env

```text
# port
VITE_PORT = 3100

# spa-title
VITE_GLOB_APP_TITLE = Simple Admin

# spa shortname
VITE_GLOB_APP_SHORT_NAME = Simple_Admin

```

可以修改调试端口和系统名称


> .env.development

```text
# Whether to open mock
VITE_USE_MOCK = false

# public path
VITE_PUBLIC_PATH = /

# Cross-domain proxy, you can configure multiple
# Please note that no line breaks
VITE_PROXY = [["/sys-api","http://localhost:8500"],["/file-manager","http://localhost:9102"]]

VITE_BUILD_COMPRESS = 'none'

# Delete console
VITE_DROP_CONSOLE = false

# Basic interface address SPA
VITE_GLOB_API_URL=

# File upload address， optional
VITE_GLOB_UPLOAD_URL=/upload

# File store address
VITE_FILE_STORE_URL=http://localhost:8080

# Interface prefix
VITE_GLOB_API_URL_PREFIX=

```

主要修改 VITE_PROXY 用于实现跨域请求不同地址，
> 注意必须写在同一行
VITE_PROXY = [["/sys-api","http://localhost:8500"],["/file-manager","http://localhost:9102"]]