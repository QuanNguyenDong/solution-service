package main

import (
	"fmt"
	"net/http"

	"github.com/QuanNguyenDong/solution-service/internal/router"
)

func main() {
	router := router.New()

	port := "8080"
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}
