package internal

import (
	"database/sql"
	"fmt"

	"github.com/BlackChaosNL/Orchestra/cmd/api/models"
	"github.com/BlackChaosNL/Orchestra/config"
	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Postgres struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	TimeZone string
}

type SQLite struct {
	databasePath string
}

func setupSQLite(s SQLite) *gorm.DB {
	_, dir, file_name := GetLatestSqleanRelease()
	const ext_name = "sqlite_ext"
	sql.Register(ext_name,
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				fmt.Sprintf("%s/%s", dir, file_name),
			},
		},
	)

	conn, _ := sql.Open(ext_name, s.databasePath)
	db, _ := gorm.Open(sqlite.Dialector{
		DriverName: ext_name,
		DSN:        s.databasePath,
		Conn:       conn,
	}, &gorm.Config{
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})

	return db
}

func setupPostgres(p Postgres) *gorm.DB {
	dbString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", p.user, p.password, p.dbname, p.port, p.TimeZone)
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	return db
}

func SetupDB() *gorm.DB {
	isPostgresEnabled := config.GetBoolFromEnv("ORCHESTRA_API_POSTGRES_ENABLED", false)
	if isPostgresEnabled {
		return setupPostgres(Postgres{
			host:     config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_HOST", "postgres"),
			user:     config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_USER", "postgres"),
			password: config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_PASSWORD", "postgres"),
			dbname:   config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_DBNAME", "postgres"),
			port:     config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_PORT", "5432"),
			TimeZone: config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_TIMEZONE", "Etc/UTC"),
		})
	} else {
		return setupSQLite(SQLite{databasePath: config.GetStrFromEnv("ORCHESTRA_API_SQLITE_DATABASE_PATH", "./sqlite.sqlite3")})
	}
}

func LoadTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.Membership{})
	db.AutoMigrate(&models.Setting{})

	db.SetupJoinTable(&models.User{}, "Groups", &models.Membership{})
}
