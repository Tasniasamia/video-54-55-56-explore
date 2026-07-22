package db

import (
	"log"
	"mains/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"fmt"
)

func GetConnectionString(cfg config.Dbconfig) string {
	
	connectStr:= fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", cfg.Dbuser, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname);

	if(!cfg.EnableSSLMode){
		connectStr += " sslmode=disable"
	}
	return connectStr
}



func NewConnection(cfg config.Dbconfig) *sqlx.DB {
	connStr := GetConnectionString(cfg)
	db, err := sqlx.Connect("postgres", connStr)
	log.Println(db);
	log.Println(err)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	return db
}