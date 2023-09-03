package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBHandler struct {
	db *gorm.DB
}

func NewDBHandler(dsn string) (*DBHandler, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "wblzero_db.",
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}

	return &DBHandler{db: db}, nil
}

func (dbh *DBHandler) GetDB() *gorm.DB {
	return dbh.db
}

func (dbh *DBHandler) Close() {
	if dbh.db != nil {
		sqlDB, err := dbh.db.DB()
		if err != nil {
			log.Printf("Error getting database connection: %s", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %s", err)
		}
	}
}
