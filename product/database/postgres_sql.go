package database

import (
	"fmt"
	"time"

	"github.com/AdiPP/e-commerce/product/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresSqlDatabase struct {}

var (
	host = "localhost"
	user = "developer"
	pasword = "developer"
	dbname = "acme_product"
	port = "5432"
	sslmode="disable"
	dbConn *gorm.DB
)

func NewPostgresSqlDatabase() Database {
	CreateDBConnection()
	AutoMigrateDB()

	return &PostgresSqlDatabase{}
}

func CreateDBConnection() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		pasword,
		dbname,
		port,
		sslmode,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	dbConn = db

	return err
}

func CheckDBConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()

	if err != nil {
		return dbConn, err
	}

	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}

	return dbConn, nil
}

func AutoMigrateDB() error {
	db, err := CheckDBConnection()

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&entity.Category{}, &entity.Product{})

	return err
}

func (d *PostgresSqlDatabase) GetDBConnection() (*gorm.DB, error) {
	db, err := CheckDBConnection()

	if err != nil {
		return db, err
	}

	return db, nil
}