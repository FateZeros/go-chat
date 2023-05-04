# 项目初始化

```bash
go mod init [projectName]

go mod tidy
```

# 项目引入 swag

[gin-swagger](https://pkg.go.dev/github.com/swaggo/gin-swagger@v1.4.3#section-readme)

```bash
go get -u github.com/swaggo/swag/cmd/swag

swag init

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

访问 swagger 地址: http://localhost:6062/swagger/index.html

## 引入 gin jwt

[gin jwt](https://pkg.go.dev/github.com/appleboy/gin-jwt/v2#section-readme)

```bash
go get github.com/appleboy/gin-jwt/v2
```

## 参考

[文档格式](https://docs.github.com/zh/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax)

[golang 文档](https://www.topgoer.com/)
