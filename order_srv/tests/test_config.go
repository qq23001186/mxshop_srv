package tests

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"nd/order_srv/proto"
)

var (
	TargetAddr  = "127.0.0.1:50060"
	OrderClient proto.OrderClient
	Conn        *grpc.ClientConn
)

func Init() {
	var err error
	Conn, err = grpc.Dial(TargetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	OrderClient = proto.NewOrderClient(Conn)
}
