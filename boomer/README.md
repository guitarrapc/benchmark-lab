## boomer

Benchmark with Boomer, A better load generator for locust, written in golang.

> see: https://github.com/myzhan/boomer

## Getting started

based on article https://qiita.com/shka0909/items/ea0ec3ddaecb3dfa8239

```
go mod init github.com/guitarrapc/benchmark-lab
go mod tidy
# update to latest boomer. 1.6.0 not accept locust 2.0
go get -u github.com/myzhan/boomer@master
go mod tidy

go build -o server.exe ./cmd/server/main.go
go build -o boomer.exe ./cmd/boomer/main.go

docker compose up --build
```
