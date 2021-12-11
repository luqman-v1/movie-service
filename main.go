package main

import (
	"movie-service/repo/mysql"
	"movie-service/repo/redis"
	"movie-service/transport/grpc"
	"movie-service/util"
)

func main() {
	_ = util.Env()
	mysql.RunMigrate()
	mysql.Conn()
	redis.Conn()
	grpc.Start()
}
