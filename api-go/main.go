package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gleberphant/puc-react-app/api-go/configs"
	"github.com/gleberphant/puc-react-app/api-go/intermediarios"
	"github.com/gleberphant/puc-react-app/api-go/manipuladores"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	roteador := http.NewServeMux()

	manipuladores.InjetarRotasLogin(roteador)
	manipuladores.InjetarRotasPage(roteador)
	manipuladores.InjetarRotasUsuarios(roteador)

	handler := intermediarios.ReqSizeMiddleware(
		intermediarios.LogMidleware(
			intermediarios.AuthMidleware(
				roteador,
			),
		),
	)
	handler = intermediarios.CorsMiddleware(handler)

	servidor := http.Server{
		Addr:              configs.PORTA,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Printf("\n Starting API server. Ambiente %s \n", configs.AMBIENTE)

	if err := servidor.ListenAndServe(); err != nil {
		log.Printf("\n Erro no servidor api... %v", err.Error())
	}

	log.Println("Finalizando ...")
}
