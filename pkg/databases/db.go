package database

import (
	"log"

	"github.com/chonlasit2000/kawaii-shop/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	// Connect
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed : %v", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenCons())
	return db
}
