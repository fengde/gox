```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
go mod download
```
