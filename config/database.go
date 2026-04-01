package config

import (
	"fmt"

	"github.com/kataras/golog"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Postgres struct {
	host     string
	user     string
	password string
	dbname   string
	port     uint16
	TimeZone string
}

type SQLite struct {
	databasePath string
}

func SetupSQLite(s SQLite) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(s.databasePath), &gorm.Config{})

	if err != nil {
		golog.Fatalf("Database can not be created... %s", err)
	}

	return db
}

func SetupPostgres(p Postgres) *gorm.DB {
	dbString := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", p.user, p.password, p.dbname, p.port, p.TimeZone)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		golog.Fatalf("Connection can not be created... %s", err)
	}

	return db
}

func SetupDB() {
	db := SetupSQLite(SQLite{databasePath: "./db/sqlite.sqlite3"})
	DB = db
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
