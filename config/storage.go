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

// Config function connect to database
func Postgres(env models.Environment) *repository.Queries {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.DatabaseHost, env.DatabasePort, env.DatabaseUser, env.DatabasePass, env.DatabaseName,
		))
	if err != nil {
		panic(err)
	}

	return repository.New(db)
}
