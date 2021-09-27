package main

import (
	"github.com/vitorbgouveia/go-rabbitmq/configs/env"
	"github.com/vitorbgouveia/go-rabbitmq/configs/stream"
	"github.com/vitorbgouveia/go-rabbitmq/internal/app/server"
)

func main() {
	env.Load()
	server.Init()
	stream.Start()
}
