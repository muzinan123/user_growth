package main

import (
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	"user_growth/conf"
	"user_growth/dbhelper"
	"user_growth/pb"
	"user_growth/ugserver"
)

func initDb() {
	// default UTC time location
	time.Local = time.UTC
	// Load global config
	conf.LoadConfigs()
	// Initialize db
	dbhelper.InitDb()
}

func main() {
	// 初始化数据库实例
	initDb()

	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//creds, err := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
	//if err != nil {
	//	log.Fatalf("credentials.NewServerTLSFromFile error=%v", err)
	//}
	//opts := []grpc.ServerOption{
	//	grpc.WriteBufferSize(1024 * 1024 * 1), // 默认32KB
	//	grpc.ReadBufferSize(1024 * 1024 * 1),  // 默认32KB
	//	grpc.KeepaliveParams(keepalive.ServerParameters{
	//		MaxConnectionIdle:     1000,             // 最多连接数
	//		MaxConnectionAge:      1 * time.Hour,    // 连接最长时间
	//		MaxConnectionAgeGrace: 10 * time.Minute, // 最长时间后延迟关闭
	//		Time:                  2 * time.Minute,  // ping间隔
	//		Timeout:               3 * time.Second,  // ping超时
	//	}),
	//	grpc.MaxConcurrentStreams(1000),
	//	grpc.ConnectionTimeout(time.Second * 1), // 连接超时
	//	grpc.Creds(creds),
	//}
	//s := grpc.NewServer(opts...)
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterUserCoinServer(s, &ugserver.UgCoinServer{})
	pb.RegisterUserGradeServer(s, &ugserver.UgGradeServer{})
	// 启动服务
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
