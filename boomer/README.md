## boomer

Benchmark with Boomer, A better load generator for locust, written in golang.

> see: https://github.com/myzhan/boomer

## Getting started

based on article https://qiita.com/shka0909/items/ea0ec3ddaecb3dfa8239

Create workspace and move, open VSCode.
```
mkdir boomber
cd boomer
code .
```

Initialize Go, write your code.
```
go mod init github.com/guitarrapc/benchmark-lab
go mod tidy
# update to latest boomer, 1.6.0 not accept locust 2.0. see https://github.com/myzhan/boomer/issues/168
go get -u github.com/myzhan/boomer@master
go mod tidy
```

Build server to confirm code pass CI.

```
go build -o server.exe ./cmd/server/main.go
go build -o boomer.exe ./cmd/boomer/main.go
```

You are ready to run.Use docker compose to launch Locust, server and boomer. Make sure there are boomer connected to locust log `Boomer is connected to master(tcp://master:5557) press Ctrl+c to quit.` and Locust accepted log `1d0b9ec1ddb3/INFO/locust.runners: Client '7493afc4ceca_6c1e6cc0977f49e8a025335e4690486f' reported as ready. Currently 1 clients ready to swarm.`.

```
docker compose up --build
```

![image](https://user-images.githubusercontent.com/3856350/182552392-0ef3027d-4948-44fe-babe-f7221536e988.png)

Open Locust and run your benchmark.

```
open http://localhost:8089
```


![image](https://user-images.githubusercontent.com/3856350/182551847-703f3767-4510-4904-8bd5-92e37d50b6be.png)
![image](https://user-images.githubusercontent.com/3856350/182551984-4a260f97-acfa-453c-860b-78c05758b576.png)
