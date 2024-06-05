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
	
}