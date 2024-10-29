package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	baseURL := "http://localhost:8080/lease"

	if len(os.Args) < 2 {
		fmt.Println("Please provide arguments.")
		return
	}

	taskId := os.Args[1]

	params := url.Values{}
	params.Add("id", taskId)
	params.Add("duration", "20s")

	// Faz uma solicitação HTTP para criar um lease
	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Lease solicitado")
}
