package postgres

import (
	"context"
	"database/sql"
	"log"
	"runtime"
	"unsplash_analog/config"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresqlDatabase struct {
	conn *pgxpool.Pool
}

var ctx context.Context = context.Background()

var DB PostgresqlDatabase

// Connect to database
func InitDB(url string, migrate_version uint) error {
	var err error

	DB.conn, err = pgxpool.Connect(context.Background(), url)
	if err != nil {
		return err
	}
	log.Println("Connected to database")
	// down if set in config
	if config.Conf.DOWN_OLD_DB_EVERYTIME {
		if err = dropDatabase(url); err != nil {
			return err
		}
	}
	// migrate
	err = migrateDB(url, migrate_version)
	if err != nil {
		return err
	}

	return DB.conn.Ping(ctx)
}

// Drop database
func dropDatabase(url string) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	log.Println("Database migration...")
	// Connect
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	var m *migrate.Migrate = new(migrate.Migrate)
	m, err = migrate.NewWithDatabaseInstance(
		"file://postgres/migrations",
		"postgres", driver)
	if err != nil || m == nil {
		return err
	}
	defer m.Close()
	// Drop
	if err = m.Drop(); err.Error() != "no change" {
		return err
	}
	return nil
}

// Migrate database
func migrateDB(url string, version uint) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	log.Println("Database migration...")
	// Connect
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	var m *migrate.Migrate = new(migrate.Migrate)
	m, err = migrate.NewWithDatabaseInstance(
		"file://postgres/migrations",
		"postgres", driver)
	if err != nil || m == nil {
		return err
	}
	defer m.Close()
	// Get version
	if version == 0 {
		// Up to latest version
		if err = m.Up(); err.Error() != "no change" {
			return err
		}
		return nil
	}
	// Migrate to fixed version
	if err = m.Migrate(version); err.Error() != "no change" {
		return err
	}
	return nil
}

// Log errors
func error_logging(err error) {
	if err != nil {
		pc := make([]uintptr, 10)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		// fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		log.Printf("[Postgres] error on %s: %s", frame.Function, err)
	}
}
