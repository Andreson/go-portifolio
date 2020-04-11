package pw_mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbStringConnection string
)

const dbDriver = "mysql"

func init() {
	fmt.Println(" \n-- CONFIGURAÇÃO MYSQL  -- \n")
	if stringConnection, err := buildStringConnection(); err != nil {
		fmt.Println("" +
			"Erro ao inicializar banco de dados, verifique se as variaveis de ambiente foram declaradas : \n" +
			"\n  PW_MYSQL_USERNAME " +
			"\n  PW_MYSQL_PASSWORD " +
			"\n  PW_MYSQL_DATABASE_NAME " +
			"\n PW_MYSQL_PORT  " +
			"\n " +
			"\nE tente novamente. \n")
		panic(errors.New("Falha ao inicializar conexao com bando de dados"))
	} else {
		dbStringConnection = stringConnection
	}
}

func buildStringConnection() (strConnection string, erro error) {

	if strConnection = os.Getenv("PW_MYSQL_STR_CON"); strConnection != "" {
		return
	} else if dbName := os.Getenv("PW_MYSQL_DATABASE_NAME"); dbName != "" {

		var dbUser, dbHost, dbPort string

		if dbUser = os.Getenv("PW_MYSQL_USERNAME"); dbUser == "" {
			dbUser = "root"
		}
		if dbPort = os.Getenv("PW_MYSQL_PORT"); dbPort == "" {
			dbPort = "3306"
		}
		if dbHost = os.Getenv("PW_MYSQL_HOST"); dbHost == "" {
			dbHost = "localhost"
		}
		dbPass := os.Getenv("PW_MYSQL_PASSWORD")
		//	strConnection := "root:123@tcp(172.17.0.3:3306)/funcionarios"
		if dbPass == "" {
			strConnection = fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
		} else {
			strConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		}

	} else {
		fmt.Println("Erro ao inicializar confis")
		erro = errors.New("Falha ai inicializar conexao com bando de dados")
	}
	return
}

func DbConn() (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbStringConnection)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(" -- Erro ao pingar conexão com banco de dados", err)
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}
