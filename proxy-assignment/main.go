package main

import (
	"inwpuun/proxy_assignment/config"
	"inwpuun/proxy_assignment/server"
)

func main() {
	conf := config.GetConfig()
	server.NewEchoServer(conf).Start()
}
