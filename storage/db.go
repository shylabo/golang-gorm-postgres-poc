package storage

import (
	"log"

	"github.com/shylabo/golang-postgres-poc/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	dsn := config.GetPostgresConnectionString()
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
