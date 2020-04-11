package main

import (
	"database/sql"
	"fmt"
	pw_dao "github.com/Andreson/go-portifolio/persistence-wrapper/dao"

)

func main() {

	books := make([]Book,0,0);

	callback := func(stmt *sql.Rows ){
			var id,userId int
			var ok bool
			var title string

			stmt.Scan(&id,&userId,&title,&ok)

		books = append(books,Book{Id:id,UserId:userId, Title:title, Ok:ok})
	}

	pw_dao.Query("select * from book",callback)

	fmt.Println("books : ",books)

}


type Book struct {
	Id     int
	UserId int
	Title  string
	Ok     bool
}

// func teste() {
// 	query := pq_metadata.BuildSelectQuery(Pessoa{Name: "Thiago", Idade: 31, Endereco: "av piassnaguabá"})

// 	insert := pq_metadata.BuildInsertQuery(Pessoa{Name: "Thiago", Idade: 31, Endereco: "av piassnaguabá"})

// 	fmt.Printf("Insert query : %s \nSelect query: %s \n", query, insert)
// }

// type Pessoa struct {
// 	Name     string
// 	Idade    int
// 	Endereco string
// }
