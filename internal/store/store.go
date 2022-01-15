package store

import (
	"database/sql"
	"github.com/bulatok/task/internal/models"
	"github.com/bulatok/task/internal/store/SQLdb"
)

type Store interface{
	Open() error
	Close() error
	DeleteByID(int) (error)
	AddUser(models.User) error
	GetUsers() (string, error)
}
func NewPSQL(dbLink string) (*SQLdb.PostgresDB){
	db := &SQLdb.PostgresDB{DB: &sql.DB{}, Link: dbLink}
	return db
}