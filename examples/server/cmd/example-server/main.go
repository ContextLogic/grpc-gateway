package main

import (
	"flag"

	"github.com/ContextLogic/grpc-gateway/examples/server"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := server.Run(); err != nil {
		glog.Fatal(err)
	}
}
