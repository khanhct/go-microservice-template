package main

import (
	"os"

	"github.com/spf13/viper"

	"casorder/cmd"
	"casorder/utils/mgrpc"
)

func main() {
	cmd.Initialize()

	gs := mgrpc.New(viper.GetString("grpc.host"), viper.GetString("grpc.port"))
	if err := gs.Start(); err != nil {
		os.Exit(1)
	}
}
