# 错误处理

> RPC 错误

```go
status.Error(codes.Internal, result.Error.Error())
```

直接 return status.Error

> 注意： codes.InvalidArgument 会转化为前端 http.http.StatusBadRequest , 会产生弹窗， 通过 api ErrorMessageMode 控制

> API 错误

```go
errorx.NewApiError(http.StatusUnauthorized, "Please log in")
```
