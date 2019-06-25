package os

import (
	"database/sql"
	"sync"
)

var (
	once sync.Once
)

type api struct {
	db *sql.DB
}

func newAPI() *api {
	return &api{}
}

func (o *api) init() {

}
