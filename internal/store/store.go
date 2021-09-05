package store

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/ppal31/mygo/internal/seeds"
	"github.com/ppal31/mygo/internal/store/database"
	"github.com/ppal31/mygo/internal/store/database/migrate"
	"time"
)

func Connect(driver, datasource string, seed bool) (*database.DataStore, error) {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		return nil, err
	}
	dbx := sqlx.NewDb(db, driver)
	if err := pingDatabase(dbx); err != nil {
		return nil, err
	}
	if err := setupDatabase(dbx); err != nil {
		return nil, err
	}
	if err := seedDatabase(dbx, seed); err != nil {
		return nil, err
	}

	return database.New(dbx), nil
}

func seedDatabase(dbx *sqlx.DB, seed bool) error {
	if !seed {
		return nil
	}
	return seeds.Execute(dbx)
}

// helper function to ping the database with backoff to ensure
// a connection can be established before we proceed with the
// database setup and migration.
func pingDatabase(db *sqlx.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}

// helper function to setup the databsae by performing automated
// database migration steps.
func setupDatabase(db *sqlx.DB) error {
	return migrate.Migrate(db)
}
