# FAQ

> Q: What is the default account?

A: account: admin   password: simple-admin

> Q: Why the account I register cannot log in?

A: The use register's default role is member. He does not have authorization to log in. \
You should modify the role's authority or change the user's role.

> Q: How to solve cross domain problems？

A: Modify api/core.go

```go
server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
```

Change  * to your own IP or domain, default is * which means allow all IPs or domain。