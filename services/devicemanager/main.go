package main

import (
	"github.com/PandoCloud/pando-cloud/pkg/server"
)

func main() {
	// init server
	err := server.Init("devicemanager")
	if err != nil {
		server.Log.Fatal(err)
		return
	}

	// register a rpc service
	/*
		r, err := NewRegistry()
		if err != nil {
			server.Log.Fatal(err)
			return
		}
		err = server.RegisterRPCHandler(r)
		if err != nil {
			server.Log.Errorf("Register RPC service Error: %s", err)
			return
		}
	*/

	// start to run
	err = server.Run()
	if err != nil {
		server.Log.Fatal(err)
	}
}
