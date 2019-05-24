package main

import (
	"fmt"
	"net"
	"os"

	"github.com/deps-cloud/rds/api"
	"github.com/deps-cloud/rds/pkg/config"
	"github.com/deps-cloud/rds/pkg/remotes"
	"github.com/deps-cloud/rds/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func main() {
	configPath := "${HOME}/.rds/config.yaml"
	port := 8090

	cmd := &cobra.Command{
		Use:   "rds",
		Short: "rds runs the repository discovery service.",
		Run: func(cmd *cobra.Command, args []string) {
			address := fmt.Sprintf(":%d", port)

			cfg, err := config.Load(os.ExpandEnv(configPath))
			if err != nil {
				panic(err)
			}

			remote, err := remotes.ParseConfig(cfg)
			if err != nil {
				panic(err)
			}

			listener, err := net.Listen("tcp", address)
			if err != nil {
				panic(err)
			}

			impl := service.NewServer(remote)

			server := grpc.NewServer()
			api.RegisterRepositoryDiscoveryServiceServer(server, impl)

			logrus.Infof("[main] starting gRPC on %s", address)
			if err := server.Serve(listener); err != nil {
				panic(err)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&configPath, "config", configPath, "The path to the config file")
	flags.IntVar(&port, "port", port, "The port to run on")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
