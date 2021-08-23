package buildlist

import (
	"database/sql"

	"github.com/Nonne46/Builds-List/internal/app/store/sqlstore"
	_ "github.com/mattn/go-sqlite3"
)

func Start() error {
	db, err := newDB("Builds.db")
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store, db)

	return srv.router.Run(":8080")
}

func newDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
