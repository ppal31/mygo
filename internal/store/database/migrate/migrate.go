package migrate

import (
	"github.com/jmoiron/sqlx"
	"github.com/ppal31/mygo/internal/store/database/migrate/sqlite"
)

// Migrate performs the database migration.
func Migrate(db *sqlx.DB) error {
	switch db.DriverName() {
	default:
		return sqlite.Migrate(db.DB)
	}
}
