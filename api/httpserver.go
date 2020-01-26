package main

import (
	"fmt"
	"log"
	"net/http"

	"./controller"
	"./model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Server Stargted Successfully")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
