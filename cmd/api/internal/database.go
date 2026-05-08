package internal

import (
	"fmt"

	"github.com/BlackChaosNL/Orchestra/cmd/api/models"
	"github.com/BlackChaosNL/Orchestra/config"
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
	port     string
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
	dbString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", p.user, p.password, p.dbname, p.port, p.TimeZone)
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
	isPostgresEnabled := config.GetBoolFromEnv("ORCHESTRA_API_POSTGRES_ENABLED", false)
	if isPostgresEnabled {
		DB = SetupPostgres(Postgres{
			host:     config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_HOST", "postgres"),
			user:     config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_USER", "postgres"),
			password: config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_PASSWORD", "postgres"),
			dbname:   config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_DBNAME", "postgres"),
			port:     config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_PORT", "5432"),
			TimeZone: config.GetStrFromEnv("ORCHESTRA_API_POSTGRES_TIMEZONE", "Etc/UTC"),
		})
	} else {
		DB = SetupSQLite(SQLite{databasePath: config.GetStrFromEnv("ORCHESTRA_API_SQLITE_DATABASE_PATH", "./db/sqlite.sqlite3")})
	}
}

func LoadTables() {
	Migrate(
		&models.User{},
		&models.Group{},
		&models.Membership{},
		&models.Setting{},
	)
	DB.SetupJoinTable(&models.User{}, "Groups", &models.Membership{})
}

func Migrate(tables ...any) error {
	return DB.AutoMigrate(tables...)
}
