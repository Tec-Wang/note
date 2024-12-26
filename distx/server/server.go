package main

import (
	"context"
	"wangzheng/brain/distx/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		ctx := context.Background()
		//调用 grpc port 50052
		conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := proto.NewHelloClient(conn)
		res, err := client.Hello(ctx, &proto.HelloRequest{
			Name: "wangzheng",
		})

		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"message": res.Reply,
		})

	})

	r.GET("/insert", func(c *gin.Context) {
		ctx := context.Background()
		//调用 grpc port 50052
		conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := proto.NewInsertClient(conn)
		res, err := client.Insert(ctx, &proto.InsertRequest{
			Name: "wangzheng",
		})

		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"message": res.Id,
		})

	})

	r.Run(":8080")
}
