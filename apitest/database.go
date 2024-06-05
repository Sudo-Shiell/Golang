package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"api/config"
)

func Connect() (*bun.DB, error) {
	cfg := time.LoadLocation(os.Getenv("DB_TZ"))
	if err != nil {
		retrun nir, err
	}
	c := mysql.Config{
		User:                 cfg.User,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 cfg.Address,
		DBName:               cfg.Name,
		Collation:            cfg.Collation,
		Loc:                  jst,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	sqldb, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err = sqldb.Ping(); err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, mysqldialect.New())
	return db, nil
}