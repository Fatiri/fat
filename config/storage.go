/*
Package driver notes
* Service only configuration to database, using postgres
*/
package config

import (
	"database/sql"
	"fmt"

	"github.com/FAT/models"
	"github.com/FAT/repository"

	_ "github.com/lib/pq"
)

type Storage interface {
	Postgres() (*repository.Queries, error)
}

type StorageCtx struct {
	env *models.Environment
}

func NewStorage(env *models.Environment) Storage {
	return &StorageCtx{
		env: env,
	}
}

// Config function connect to database
func (s *StorageCtx) Postgres() (*repository.Queries, error) {
	fmt.Println(s.env)
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			s.env.DatabaseHost, s.env.DatabasePort, s.env.DatabaseUser, s.env.DatabasePass, s.env.DatabaseName,
		))
	if err != nil {
		return nil, err
	}

	return repository.New(db), nil
}
