package db

import (
	"database/sql"
	"fmt"
	"github.com/micro/go-log"
)

type db struct {
	Mysql Mysql `json："mysql"`
}

// Mysql mySQL 配置
type Mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func initPG() {

	log.Logf("[initPG] init postgreSQL")

	var err error

	pgDB, err = sql.Open("postgres", parseConnectStr())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	pgDB.SetMaxOpenConns(pgConfig.MaxOpenConnection)
	pgDB.SetMaxIdleConns(pgConfig.MaxIdleConnection)

	if err = pgDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Logf("[initPG] pg connected")
}

func parseConnectStr() string {

	str := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?", pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.DBName)

	log.Logf("[initPG] pg connected %s", str)

	str = fmt.Sprintf("%ssslmode=%s", str, pgConfig.SSLMode)

	if pgConfig.SSLMode != "disable" {

		if pgConfig.SSLCert != "" {
			str += "&sslcert=" + pgConfig.SSLCert
		}

		if pgConfig.SSLKey != "" {
			str += "&sslkey=" + pgConfig.SSLKey
		}

		if pgConfig.SSLRootCert != "" {
			str += "&sslrootcert=" + pgConfig.SSLRootCert
		}
	} else {

	}

	return str
}
