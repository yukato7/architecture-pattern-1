package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/yutify/gengo-api/config"
)

const (
	chargeLogTable = "charge_logs"
	userTable      = "users"
)

type Client struct {
	*sql.DB
}

func New(config *config.DB) (*Client, error) {
	conf := &mysql.Config{
		User:                 config.User,
		Passwd:               config.Password,
		Addr:                 config.Host + config.Port,
		Net:                  "tcp",
		DBName:               config.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &Client{db}, nil
}

func withTransaction(ctx context.Context, db *sql.DB, txFunc func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil) //Todo add isolationLevel
	if err != nil {
		return fmt.Errorf("failed to begin transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println(err)
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
