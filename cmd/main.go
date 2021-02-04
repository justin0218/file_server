package main

import (
	"file_server/api"
	"file_server/pkg/etcd"
	"file_server/store"
	"fmt"
	"time"
)

func main() {
	redis := new(store.Redis)
	err := redis.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	log := new(store.Log)
	go etcd.Register()
	go api.GrpcServer()
	log.Get().Debug("server started at %v", time.Now())
	fmt.Printf("server started at %v \n", time.Now())
	select {}
}
