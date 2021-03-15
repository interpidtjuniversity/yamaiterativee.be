package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"strings"
	"time"
	"yama.io/yamaIterativeE/internal/conf"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func openDB(opts conf.DatabaseOpts) (*gorm.DB, error) {
	dsn, err := parseDSN(opts)
	if err != nil {
		return nil, errors.Wrap(err, "parse DSN")
	}

	return gorm.Open(opts.Type, dsn)
}

func parseDSN(opts conf.DatabaseOpts) (dsn string, err error) {
	// In case the database name contains "?" with some parameters
	concate := "?"
	if strings.Contains(opts.Name, concate) {
		concate = "&"
	}
	switch opts.Type {
	case "sqlite3":
		dsn = "file:" + opts.Path + "?cache=shared&mode=rwc"
	default:
		return "", errors.Errorf("unrecognized dialect: %s", opts.Type)
	}

	return dsn, nil
}

func Init() (*gorm.DB, error) {
	db, err := openDB(conf.Database)
	if err != nil {
		return nil, errors.Wrap(err, "open database")
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(conf.Database.MaxOpenConns)
	db.DB().SetMaxIdleConns(conf.Database.MaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Minute)

	if !conf.IsProdMode() {
		db = db.LogMode(true)
	}

	switch conf.Database.Type {
	case "sqlite3":
		conf.UseSQLite3 = true
	}

	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(time.Microsecond)
	}

	return db, db.DB().Ping()
}