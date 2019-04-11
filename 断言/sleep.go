package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

//
//var period = flag.Duration("period", 5*time.Second, "sleep period")
//
//func main2() {
//	flag.Parse()
//	// time.Sleep(1 * time.Second)
//	fmt.Printf("Sleeping for %v...", *period)
//	time.Sleep(*period)
//
//	fmt.Println("=========")
//}
type HelloService struct{}

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		fmt.Println("accept")
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
