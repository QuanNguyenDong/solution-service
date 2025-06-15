package main

import (
	"fmt"
	"net/http"

	"github.com/QuanNguyenDong/solution-service/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := router.New()

	port := "8080"
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}
