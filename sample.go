package main

import (
	"flag"
	"gosample/gorpc"

	"github.com/golang/glog"
)

func main() {

	flag.Parse()
	defer glog.Flush()

	//go_chain.RUN()
	//gorpc.Server()
	gorpc.Client()

}
