package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Starting the server on :8080")

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
}