package main

import (
	"log"
	"os"
	"toggler/configs"
	"toggler/routes"

	"net/http"
)

func main() {
	pid := os.Getpid()

	log.Printf("Just ran subprocess %d, exiting\n", pid)

	prepareDB()
	prepareServer()
}

func prepareServer() {
	serverPort := configs.EnvPort()
	hostname := configs.EnvHostname()

	router := routes.DoRouting()

	serverAddress := hostname + ":" + serverPort

	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  configs.DefaultReadTimeout(),
		WriteTimeout: configs.DefaultWriteTimeout(),
		Handler:      router,
	}

	log.Println("Starting server on " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func prepareDB() {
	if configs.ShouldConnectDB() {
		configs.ConnectToDB()
	}
}
