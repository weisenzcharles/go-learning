package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	// 如果客户端不需要对服务端鉴权，可以设置 InsecureSkipVerify:true，即可跳过对服务端的鉴权
	// config := &tls.Config{
	// 	InsecureSkipVerify: true,
	// }
	// conn, _ := tls.Dial("tcp", "localhost:1234", config)
	// defer conn.Close()
	// client := rpc.NewClient(conn)

	// var result Result
	// if err := client.Call("Cal.Square", 12, &result); err != nil {
	// 	log.Fatal("Failed to call Cal.Square. ", err)
	// }

	// log.Printf("%d^2 = %d", result.Num, result.Ans)

	// 服务端对客户端鉴权配置
	// cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")
	// certPool := x509.NewCertPool()
	// certBytes, _ := ioutil.ReadFile("../server/server.crt")
	// certPool.AppendCertsFromPEM(certBytes)
	// config := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// 	RootCAs:      certPool,
	// }

	certPool := x509.NewCertPool()
	// 如果需要对服务器端鉴权，那么需要将服务端的证书添加到信任证书池中
	certBytes, err := ioutil.ReadFile("../server/server.crt")
	if err != nil {
		log.Fatal("Failed to read server.crt")
	}
	certPool.AppendCertsFromPEM(certBytes)

	config := &tls.Config{
		RootCAs: certPool,
	}

	conn, _ := tls.Dial("tcp", "localhost:1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	var result Result
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
