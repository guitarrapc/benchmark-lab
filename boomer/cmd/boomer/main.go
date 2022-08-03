package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/myzhan/boomer"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func hello() {

	host := os.Getenv("SERVER_ENDPOINT")
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		host = "localhost:50051"
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// レスポンスタイムを計測するために、リクエストを投げる前の時刻を取得する
	start := time.Now()
	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "boomerし太郎"})

	if err != nil {
		// リクエストに失敗した場合はRecordFailureを呼びます
		boomer.RecordFailure(
			"grpc",     // リクエストの種別
			"SayHello", // rpc名
			time.Since(start).Nanoseconds()/int64(time.Millisecond), // レスポンスタイム
			fmt.Sprintf("failed: %+v", err),                         // エラー理由
		)
	} else {
		// リクエストに成功した場合はRecordSuccessを呼びます
		boomer.RecordSuccess(
			"grpc",     // リクエストの種別
			"SayHello", // rpc名
			time.Since(start).Nanoseconds()/int64(time.Millisecond), // レスポンスタイム
			int64(len(response.String())),                           // レスポンスサイズ
		)
	}
}

func main() {
	task := &boomer.Task{
		Name:   "sample",
		Weight: 10,
		Fn:     hello,
	}

	boomer.Run(task)
}
