package main

import (
	"os"

	"github.com/chonlasit2000/kawaii-shop/config"
	"github.com/chonlasit2000/kawaii-shop/modules/servers"
	database "github.com/chonlasit2000/kawaii-shop/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := database.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
