package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"wangzheng/brain/distx/config"
	"wangzheng/brain/distx/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	proto.UnimplementedHelloServer
	proto.UnimplementedInsertServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Hello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &proto.HelloResponse{Reply: "hello: " + in.GetName()}, nil
}

func (s *server) Insert(ctx context.Context, in *proto.InsertRequest) (*proto.InsertResponse, error) {
	config := config.NewConfig()
	dsn := config.Mysql.DB2DSN
	// 连接到数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	// 执行插入操作
	res, err := db.Exec("INSERT INTO tb1 (name) VALUES (?)", in.Name)
	if err != nil {
		log.Fatalf("failed to insert data: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("failed to get last insert id: %v", err)
	}
	// 返回插入结果
	return &proto.InsertResponse{
		Id: id,
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
