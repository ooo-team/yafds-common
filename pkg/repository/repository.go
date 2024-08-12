package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	common "github.com/ooo-team/yafds-common/pkg"
)

func GetDB() *sql.DB {

	host, err := common.LoadEnvVar("dbHost")
	if err != nil {
		panic(err.Error())
	}

	port_str, err := common.LoadEnvVar("dbPort")
	if err != nil {
		panic(err.Error())
	}

	port, err := strconv.Atoi(port_str)
	if err != nil {
		panic("cannot convert string dbPort to int")
	}
	user, err := common.LoadEnvVar("dbUser")
	if err != nil {
		panic(err.Error())
	}

	password, err := common.LoadEnvVar("dbPassword")
	if err != nil {
		panic(err.Error())
	}

	dbname, err := common.LoadEnvVar("dbName")
	if err != nil {
		panic(err.Error())
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}

	return db
}
