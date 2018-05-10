package gorpc

import (
	"helloworld"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	//Address is ip:port listening by server
	Address = "localhost:8081"
)

//Client send message to server
func Client() {

	//connect to server
	conn, _ := grpc.Dial(Address, grpc.WithInsecure())
	defer conn.Close()

	message := helloworld.HelloRequest{Name: "hello, world!"}
	client := helloworld.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalln("Get response from server failed.")
	}

	log.Println(res.Name)
}
