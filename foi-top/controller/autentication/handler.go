package aut_controller

import (
	"github.com/Andreson/go-portifolio/foi-top/controller"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	aut_service "github.com/Andreson/go-portifolio/foi-top/service/autentication"
	"log"
	"net/http"
)

func  Init(){

	http.HandleFunc("/error",UnauthorizedAutentication)
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
			w.WriteHeader(http.StatusUnauthorized)

			http.Redirect(w, r, "/error", http.StatusUnauthorized)
			return
		}
		 isValid, err :=aut_service.ValidatedToken(domain.Login{TokenJWT:jwtToken})

		 if err!=nil {
			 log.Println("token nao informado")
			//erro ao validar token
			 w.Write( controller.GetSuccessMsg("Erro ao validar autenticaçao, por favor tente mais tarde") )
			 w.WriteHeader(http.StatusInternalServerError)
			 h.ServeHTTP(w, r) // call original
			 return
		 }
		 if !isValid {

			 w.Write( controller.GetSuccessMsg("AUTENTICAÇAO INVALIDA, ACESSO NAO AUTORIZADO"))
			 w.WriteHeader(http.StatusUnauthorized)
			 h.ServeHTTP(w, r) // call original
			 return
		 }



		log.Println("Before")
		h.ServeHTTP(w, r) // call original
		log.Println("After")
	})
}

