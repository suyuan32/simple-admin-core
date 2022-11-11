## Use swagger

> Environment setting

[go-swagger](https://zhuanlan.zhihu.com/p/556171256?)

> [Office Document](https://goswagger.io/use/spec/meta.html)

> In the root of project run simple-admin-core/

```shell
swagger generate spec --output=./core.yml --scan-models

swagger serve --no-open -F=swagger --port 36666 core.yaml
```

use make

```shell
make gen-api

make serve-swagger
```

![pic](../../assets/swagger.png)

> Get Token
> Firstly, log in the system and press F12 to get authorization from any request

![pic](../../assets/get_token.png)

> Copy to swagger

![pic](../../assets/swagger_authority.png)

> Comment Example

We use normally use Req suffix to represent  Request， Resp to represent Response.


#### If the type name has suffix "Req" and "Info", you can omit the swagger model type declaration.

```go
type (
    // The response data of API information | API信息
    ApiInfo {
        BaseInfo

        // API path | API路径
        Path          string `json:"path"`

        // Api translation | API 多语言翻译
        Title         string `json:"title"`

        // API Description | API 描述
        Description   string `json:"description"`

        // API group | API分组
        Group         string `json:"group"`

        // API request method e.g. POST | API请求类型 如POST
        Method        string `json:"method"`
    }
}
```

generate 

```go
// The response data of API information | API信息
// swagger:model ApiInfo
type ApiInfo struct {
	BaseInfo
	// API path | API路径
	Path string `json:"path"`
	// Api translation | API 多语言翻译
	Title string `json:"title"`
	// API Description | API 描述
	Description string `json:"description"`
	// API group | API分组
	Group string `json:"group"`
	// API request method e.g. POST | API请求类型 如POST
	Method string `json:"method"`
}

```
You can also overwrite it, just add your own comment.
```go
type (
    // The response data of API information | API信息
    // swagger:response ApiInfo
    ApiInfo {
        BaseInfo

        // API path | API路径
        Path          string `json:"path"`

        // Api translation | API 多语言翻译
        Title         string `json:"title"`

        // API Description | API 描述
        Description   string `json:"description"`

        // API group | API分组
        Group         string `json:"group"`

        // API request method e.g. POST | API请求类型 如POST
        Method        string `json:"method"`
    }
}

```
generate 
```go
// The response data of API information | API信息
// swagger:response ApiInfo
type ApiInfo struct {
	BaseInfo
	// API path | API路径
	Path string `json:"path"`
	// Api translation | API 多语言翻译
	Title string `json:"title"`
	// API Description | API 描述
	Description string `json:"description"`
	// API group | API分组
	Group string `json:"group"`
	// API request method e.g. POST | API请求类型 如POST
	Method string `json:"method"`
}
```

#### If the type name has suffix "Resp"  you can omit the swagger response type declaration. Just like "Info" and "Req".

> For route, you can just add a comment for it.

api/desc/apis.api
```go
// Create or update API information | 创建或更新API
@handler createOrUpdateApi
post /api (CreateOrUpdateApiReq) returns (SimpleMsg)
```

It will generate :

```go
package api

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/api"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /api api CreateOrUpdateApi
//
// Create or update API information | 创建或更新API
//
// Create or update API information | 创建或更新API
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateApiReq
//
// Responses:
//  200: SimpleMsg
//  401: SimpleMsg
//  500: SimpleMsg

func CreateOrUpdateApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewCreateOrUpdateApiLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateApi(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

```

You can add more config in handler because it will not be over write if the file exists. 

> Notice: goctls' generating will only overwrite internal/types/* and internal/handler/routes.go. 
> If  handler and logic and so on need to regenerate, you must delete them by yourself.