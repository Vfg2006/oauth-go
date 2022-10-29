package main

import (
	"context"
	"fmt"

	"github.com/vfg2006/oauth-go/authenticator"
	"github.com/vfg2006/oauth-go/config"
	"github.com/vfg2006/oauth-go/server"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("INIT - SERVER")
	config, err := config.NewConfigFromFile()
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server, err := server.New(
		config,
		authenticator.New(),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := server.Run(ctx); err != nil {
		logrus.Error(err)
	}
}
