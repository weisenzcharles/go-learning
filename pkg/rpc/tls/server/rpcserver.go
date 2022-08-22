package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Cal int

type Result struct {
	Num, Ans int
}

func main() {
	rpc.Register(new(Cal))

	// 服务端对客户端鉴权配置
	// cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	// certPool := x509.NewCertPool()
	// certBytes, _ := ioutil.ReadFile("../client/client.crt")
	// certPool.AppendCertsFromPEM(certBytes)
	// config := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// 	ClientAuth:   tls.RequireAndVerifyClientCert,
	// 	ClientCAs:    certPool,
	// }

	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}

}
