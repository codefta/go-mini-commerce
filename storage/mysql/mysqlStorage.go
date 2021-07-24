package mysql_storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlStorage struct {
	db *sql.DB
}

type MySqlStorageConfigs struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func NewMyqlStorage(configs MySqlStorageConfigs) (*MysqlStorage, error) {
	dsn := fmt.Sprintf("%s:%v@tcp(%v:%v)/%s?parseTime=true", configs.DBUser, configs.DBPass, configs.DBHost, configs.DBPort, configs.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection due to: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db due to: %v", err)
	}

	return &MysqlStorage{db: db}, nil
}