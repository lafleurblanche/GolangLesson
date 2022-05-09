package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	resp, _ := client.R().Get("http://localhost:8080/api/hello")
	fmt.Printf("Get resp = %+v\n", resp)
}
