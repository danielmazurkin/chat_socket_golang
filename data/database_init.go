package data

import (
	"ChatSocket/logger"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(db *sql.DB) {
	user := new(UserStruct)
	user.CreateTable(db)
	message := new(MessageStruct)
	message.CreateTable(db)
}

func InitRedisStorage() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func OpenDatabase() (db *sql.DB, err error) {

	host := GetEnvFallback("POSTGRES_HOST", "localhost")
	port := 5433
	user := GetEnvFallback("POSTGRES_USER", "chat_user")
	password := GetEnvFallback("POSTGRES_PASSWORD", "password")
	dbname := GetEnvFallback("POSTGRES_DB", "chat_user")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	logger.Log.Println("PSQL INFO: ", psqlInfo)
	db, err = sql.Open("postgres", psqlInfo)

	if err == nil {
		logger.Log.Println("Connecting with database")
		InitDatabase(db)
	} else {
		logger.Log.Println("Error with connecting ", err)
	}
	return db, err
}
