# SimpleBank
這是一個使用Go, Gin, sqlc完成的簡單銀行，使用者可以進行註冊登入和存提款以及轉移

## Getting Start
### Install
```
git pull https://github.com/wsy0414/simplebank-go.git
```

### Start Service
使用local執行可依照下列步驟
1. 準備好PostgreSQL,這邊可以使用docker處理
```
make postgres
```
2. 啟動Gin
```
make server
```
---
使用docker執行可依照下列步驟

1. 
```
make docker
```

### Swagger Doc
有使用Swagger作為API文件
```
http://localhost:8080/swagger/index.html
```