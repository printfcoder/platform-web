package disk

import "database/sql"

type diskStorage struct {
	db *sql.DB
}
