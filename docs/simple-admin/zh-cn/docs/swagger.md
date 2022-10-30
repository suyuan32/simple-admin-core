## 使用swagger

> 环境配置

[go-swagger](https://zhuanlan.zhihu.com/p/556171256?)

> [官方文档](https://goswagger.io/use/spec/meta.html)

> 在项目根目录运行 simple-admin-core/

```shell
swagger generate spec --output=./core.yml --scan-models

swagger serve --no-open -F=swagger --port 36666 core.yaml
```

![pic](../../assets/swagger.png)

> 获取 Token 
> 
> 先登录系统，在任意请求中复制 authorization

![pic](../../assets/get_token.png)

> 粘贴到 swagger

![pic](../../assets/swagger_authority.png)


> 注解规范

通常对于请求参数我们使用 Req 即 Request 的缩写， 返回值 Resp 即 Response 的缩写

```text
// 对于 post 请求体我们使用 swagger:model, 如

// Get API list request params | API列表请求参数
// swagger:model ApiListReq
ApiListReq {
    PageInfo
    // API path | API路径
    // Max length: 100
    Path          string `json:"path,optional" validate:"omitempty,max=100"`
    
    // API Description | API 描述
    // Max length: 50
    Description   string `json:"description,optional" validate:"omitempty,max=50"`
    
    // API group | API分组
    // Max length: 10
    Group         string `json:"group,optional" validate:"omitempty,alphanum,max=10"`
    
    // API request method e.g. POST | API请求类型 如POST
    // Max length: 4
    Method        string `json:"method,optional" validate:"omitempty,uppercase,max=4"`
}

// 对于 get query 请求， 可以使用 swagger:parameters 注解，也可使用model

// swagger:parameters listBars addBars
type BarSliceParam struct {
    // a BarSlice has bars which are strings
    //
    // min items: 3
    // max items: 10
    // unique: true
    // items.minItems: 4
    // items.maxItems: 9
    // items.items.minItems: 5
    // items.items.maxItems: 8
    // items.items.items.minLength: 3
    // items.items.items.maxLength: 10
    // items.items.items.pattern: \w+
    // collection format: pipe
    // in: query
    // example: [[["bar_000"]]]
    BarSlice [][][]string `json:"bar_slice"`
}

// 对于 response 来说， 使用 swagger:response 注解

// The response data of API information | API信息
    // swagger:response ApiInfo
    ApiInfo {
        // ID
        Id            uint64 `json:"id"`

        CreatedAt      int64  `json:"createdAt"`

        // API path | API路径
        Path          string `json:"path"`

        // API Description | API 描述
        Description   string `json:"description"`

        // API group | API分组
        Group         string `json:"group"`

        // API request method e.g. POST | API请求类型 如POST
        Method        string `json:"method"`
    }


// 对于 route 来说， 只需要一行注解, 如 api/api_desc/apis.api

// Create or update API information | 创建或更新API
@handler createOrUpdateApi
post /api (CreateOrUpdateApiReq) returns (SimpleMsg)

// 自动生成 api/internal/handler/api/create_or_update_api_handler.go

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
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

```

> 注意 goctls 的生成只会覆盖 internal/types/* 和 internal/handler/routes.go， 如果 handler 需要重新生成需要手动删除再生成