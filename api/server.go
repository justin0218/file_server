package api

import (
	"file_server/api/proto"
	"file_server/store"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type FileSvr struct {
}

func GrpcServer() {
	conf := new(store.Config)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", conf.Get().Etcd.Key))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	svr := grpc.NewServer(opts...)
	proto.RegisterFileServer(svr, &FileSvr{})
	err = svr.Serve(lis)
	if err != nil {
		panic(err)
	}
}
