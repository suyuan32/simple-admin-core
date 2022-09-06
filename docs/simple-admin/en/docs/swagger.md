## Use swagger

### Environment setting

[go-swagger](https://zhuanlan.zhihu.com/p/556171256?)

### In the root of project run
simple-admin-core/

```shell
swagger generate spec --output=./core.yml

swagger serve --no-open -F=swagger --port 36666 core.yaml
```

![pic](../../assets/swagger.png)
