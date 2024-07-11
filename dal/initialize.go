package dal

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbSingleton *sqlx.DB

func envOrDefault(env, def string) string {
	a := os.Getenv(env)
	if len(a) == 0 {
		return def
	}

	return a
}

func GetDatabase() *sqlx.DB {
	var err error = nil
	if dbSingleton == nil {

		username := envOrDefault("DB_USERNAME", "cao")
		password := envOrDefault("DB_PASSWORD", "cao")
		host := envOrDefault("DB_HOST", "localhost")
		port := envOrDefault("DB_PORT", "5432")
		database := envOrDefault("DB_DATABASE", "cao")

		dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, database)

		dbSingleton, err = sqlx.Connect("postgres", dsn)
		if err != nil {
			panic(err)
		}
	}

	return dbSingleton
}
