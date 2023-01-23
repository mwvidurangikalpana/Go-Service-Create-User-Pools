package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/GO-SERVICE-CREATE-USER-POOLS/router"
)

func main() {
	fmt.Println("Creating User Pool API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8088", r))
	fmt.Println("Listening at port 8088 ...")
}
