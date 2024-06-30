package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"user_growth/pb"
)

func main() {
	// 连接到服务
	addr := flag.String("addr", "localhost:80", "the address to connect to")
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithWriteBufferSize(1024 * 1024 * 1), // 默认32KB
		grpc.WithReadBufferSize(1024 * 1024 * 1),  // 默认32KB,
	}
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 请求服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 新建客户端
	cCoin := pb.NewUserCoinClient(conn)
	cGrade := pb.NewUserGradeClient(conn)
	// 测试1：UserCoinServer.ListTasks
	r1, err1 := cCoin.ListTasks(ctx, &pb.ListTasksRequest{})
	if err1 != nil {
		log.Printf("cCoin.ListTasks error=%v\n", err1)
	} else {
		log.Printf("cCoin.ListTasks: %+v\n", r1.GetDatalist())
	}
	// 测试2：UserGradeServer.ListGrades
	r2, err2 := cGrade.ListGrades(ctx, &pb.ListGradesRequest{})
	if err2 != nil {
		log.Printf("cGrade.ListGrades error=%v\n", err2)
	} else {
		log.Printf("cGrade.ListGrades: %+v\n", r2.GetDatalist())
	}
}
