package models

import (
	"fmt"
	"database/sql"
	"log"
	"../consts"
	_"github.com/lib/pq"
)

var db *sql.DB

func simpleQuery(sql string) (error) {
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func InitDB(user, password, dbname string)  {
	// Connection to db
	var err error
	connectionString := fmt.Sprintf(
		consts.InitDbString,
		user,
		dbname,
		password,
	)
	db, err = sql.Open("postgres", connectionString)

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	// Initialization db if db is absent
	db.Query(consts.CreateUsers)
	db.Query(consts.CreateTournaments)
	db.Query(consts.CreateTournamentsUsers)

	fmt.Println("Conntected...")
}

func Reset()  {
	db.Query(consts.DeleteTournamentsUsers)
	db.Query(consts.DeleteTournaments)
	db.Query(consts.DeleteUsers)
}
