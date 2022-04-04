package database

import "gorm.io/gorm"

type Database interface {
	GetDBConnection() (*gorm.DB, error)
}