# Error Handling

### RPC Error

```go
status.Error(codes.Internal, result.Error.Error())
```

Just return status.Error

### API Error

```go
errorx.NewApiError(http.StatusUnauthorized, "Please log in")
```

