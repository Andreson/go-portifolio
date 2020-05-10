package user_controller

import (
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top/controller"
	aut_controller "github.com/Andreson/go-portifolio/foi-top/controller/autentication"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	user_service "github.com/Andreson/go-portifolio/foi-top/service/user"
	"log"
	"net/http"
	"strconv"
	"strings"
)


func Init(){
	http.HandleFunc("/user",Save)
	findUserh :=http.HandlerFunc(FindById)
	http.Handle("/user/",aut_controller.FilterHandler(findUserh))
}


func FindById(w  http.ResponseWriter, r *http.Request) {

	tempIdParam := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err :=strconv.Atoi(tempIdParam)
	controller.DefaultCallBody( func()interface{}{
		return user_service.FindById(domain.UserDTO{Id:int64(  id )})
	} ,w,err)
}

func Save(w  http.ResponseWriter, r *http.Request){

	dto, err :=toDto(r)
	controller.DefaultCall( func(){
		user_service.Save(dto)
	} ,w,err)
}

func  toDto(r *http.Request) (domain.UserDTO,error) {
	decoder := json.NewDecoder(r.Body)
	var data domain.UserDTO
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Print("Request ",data)
	return data, nil
}
