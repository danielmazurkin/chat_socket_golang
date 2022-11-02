package data

import (
	"database/sql"
	"github.com/go-redis/redis/v9"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func init() {
	f, err := os.OpenFile("log_file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Println("init started")
}

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
	db, err = sql.Open("sqlite3", "database.db")

	if err == nil {
		log.Println("Connecting with database")
		InitDatabase(db)
	} else {
		log.Println("Error with connecting ", err)
	}
	return db, err
}
