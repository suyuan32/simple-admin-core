# 错误处理

### RPC 错误

```go
status.Error(codes.Internal, result.Error.Error())
```

直接 return status.Error

### API 错误

```go
errorx.NewApiError(http.StatusUnauthorized, "Please log in")
```

