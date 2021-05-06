package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"path"
	"strings"
	"time"
	"xorm.io/core"
	"xorm.io/xorm"
	"yama.io/yamaIterativeE/internal/conf"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Engine represents a XORM engine or session.
type Engine interface {
	Delete(interface{}) (int64, error)
	Exec(...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	ID(interface{}) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Sql(string, ...interface{}) *xorm.Session
	Table(interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
}

var (
	x            *xorm.Engine
	legacyTables []interface{}
	HasEngine    bool
)

func init() {
	legacyTables = append(legacyTables,
		new(User))
	gonicNames := []string{"SSL"}
	for _, name := range gonicNames {
		core.LintGonicMapper[name] = true
	}
	Proxy = TransactionProxy{}
}

func getEngine() (*xorm.Engine, error) {

	Param := "?"
	if strings.Contains(conf.Database.Name, Param) {
		Param = "&"
	}

	connStr := ""
	switch conf.Database.Type {

	case "sqlite3":
		if err := os.MkdirAll(path.Dir(conf.Database.Path), os.ModePerm); err != nil {
			return nil, fmt.Errorf("create directories: %v", err)
		}
		conf.UseSQLite3 = true
		connStr = "file:" + conf.Database.Path + "?cache=shared&mode=rwc"

	default:
		return nil, fmt.Errorf("unknown database type: %s", conf.Database.Type)
	}
	return xorm.NewEngine(conf.Database.Type, connStr)
}

func SetEngine() (*gorm.DB, error) {
	var err error
	x, err = getEngine()
	if err != nil {
		return nil, fmt.Errorf("connect to database: %v", err)
	}

	x.SetMapper(core.GonicMapper{})

	x.SetMaxOpenConns(conf.Database.MaxOpenConns)
	x.SetMaxIdleConns(conf.Database.MaxIdleConns)
	x.SetConnMaxLifetime(time.Second)

	x.ShowSQL(true)

	return Init()
}

func NewEngine() (err error) {
	if _, err = SetEngine(); err != nil {
		return err
	}

	return nil
}
