package pw_dao

import (
	"database/sql"
	pw_mysql "github.com/Andreson/go-portifolio/persistence-wrapper/connection/mysql"
	pq_metadata "github.com/Andreson/go-portifolio/persistence-wrapper/objectmetadata"
)



func Save(dto interface{}){

	insertSql := pq_metadata.BuildInsertQuery(dto)

	db :=pw_mysql.DbConn()
	insertQuery, err := db.Query(insertSql)

	if err != nil {
        panic(err.Error())
	}

	defer insertQuery.Close()

}


func Query(query string, callback func(*sql.Rows) ) {
	db :=pw_mysql.DbConn()
	stmt, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for stmt.Next() {
		callback(stmt)
	}
	defer stmt.Close()
	return
}

