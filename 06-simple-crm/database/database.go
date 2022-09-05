package database

import (
	"github.com/jinzhu/gorm"
	// to get the exact support for SQLite database
	// gorm contains language(dialect), from there sqlite will be used
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// gorm is the golang ORM which helps golang to interact with the database
	DBconn *gorm.DB

)