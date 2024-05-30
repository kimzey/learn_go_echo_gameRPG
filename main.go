package main

import (
	"github.com/kimzey/iskeai-shop/config"
	"github.com/kimzey/iskeai-shop/databases"
	"github.com/kimzey/iskeai-shop/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	server := server.NewEchoServer(conf,db.ConnectionGetting())

	server.Start()


}