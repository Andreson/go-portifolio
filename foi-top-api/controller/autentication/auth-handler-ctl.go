package aut_controller


import (
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top/controller"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	aut_service "github.com/Andreson/go-portifolio/foi-top/service/autentication"
	"log"
	"net/http"
)

func  Init(){

	http.HandleFunc("/error",UnauthorizedAutentication)
	http.HandleFunc("/login",Login)
}

func Login(w http.ResponseWriter, r *http.Request) {
	 login, err :=toDto(r)
	w.Header().Set("Content-Type", "application/json")
	if err!=nil {
		w.Write( controller.GetSuccessMsg("Erro ao receber parametros, verificar os parametros enviados e tente novamente ") )
		w.WriteHeader(http.StatusBadRequest)
	}

	if  token , err :=aut_service.GenerateToken(login) ; err==nil {
		login.TokenJWT = token
		w.Write(controller.GetSuccessBody(login))
	} else {
		w.Write( controller.GetSuccessMsg("Erro ao validar autenticaçao, por favor tente mais tarde") )
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UnauthorizedAutentication(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}

func FilterHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jwtToken := r.Header.Get("Bearer")

		if jwtToken=="" {
			w.Write( controller.GetSuccessMsg("AUTENTICAÇAO INVALIDA, ACESSO NAO AUTORIZADO"))
		//	http.Redirect(w, r, "/error", http.StatusUnauthorized)
			return
		}
		 isValid, err :=aut_service.ValidatedToken(domain.LoginDTO{TokenJWT: jwtToken})

		 if err!=nil {
			 log.Println("token nao informado : ",err)
			//erro ao validar token
			 w.Write( controller.GetSuccessMsg("Erro ao validar autenticaçao, token invalido!") )
			 w.WriteHeader(http.StatusInternalServerError)

			 return
		 }
		 if !isValid {

			 w.Write( controller.GetSuccessMsg("AUTENTICAÇAO INVALIDA, ACESSO NAO AUTORIZADO"))
			 w.WriteHeader(http.StatusUnauthorized)

			 return
		 }

		log.Println("Before")
		h.ServeHTTP(w, r) // call original
		log.Println("After")
	})
}


func  toDto(r *http.Request) (domain.LoginDTO,error) {
	decoder := json.NewDecoder(r.Body)
	var data domain.LoginDTO
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Print("Request login ",data)
	return data, nil
}

