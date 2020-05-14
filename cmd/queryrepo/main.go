package main

import (
	"github.com/alexjch/queryrepo/internal/cmd/args"
	"github.com/alexjch/queryrepo/internal/cmd/service"
)

func main() {
	var serviceArgs *args.Args
	serviceArgs = args.ParseArgs()
	srv := &service.Service{}
	srv.Start(serviceArgs.Port, serviceArgs.RepoUrl)
}
