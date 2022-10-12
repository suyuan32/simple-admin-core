# Error Handling

> RPC Error

```go
status.Error(codes.Internal, result.Error.Error())
```

Just return status.Error

> Notice： codes.InvalidArgument will convert into  http.http.StatusBadRequest in api service
> It will create prompt, you can control it by api ErrorMessageMode.

> API Error

```go
errorx.NewApiError(http.StatusUnauthorized, "Please log in")
```

