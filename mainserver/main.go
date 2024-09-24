package main

import (
	"log"
	"net"
	"time"
	"user_growth/conf"
	"user_growth/dbhelper"
	"user_growth/pb"
	"user_growth/ugserver"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
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
	// Initialize database instance
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
	//	grpc.WriteBufferSize(1024 * 1024 * 1), // Default 32KB
	//	grpc.ReadBufferSize(1024 * 1024 * 1),  // Default 32KB
	//	grpc.KeepaliveParams(keepalive.ServerParameters{
	//		MaxConnectionIdle:     1000,             // Maximum number of connections
	//		MaxConnectionAge:      1 * time.Hour,    // Maximum connection duration
	//		MaxConnectionAgeGrace: 10 * time.Minute, // Grace period for delayed closure after max duration
	//		Time:                  2 * time.Minute,  // Ping interval
	//		Timeout:               3 * time.Second,  // Ping timeout
	//	}),
	//	grpc.MaxConcurrentStreams(1000),
	//	grpc.ConnectionTimeout(time.Second * 1), // Connection timeout
	//	grpc.Creds(creds),
	//}
	//s := grpc.NewServer(opts...)
	s := grpc.NewServer()
	// Register services
	pb.RegisterUserCoinServer(s, &ugserver.UgCoinServer{})
	pb.RegisterUserGradeServer(s, &ugserver.UgGradeServer{})
	// Start the server
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
