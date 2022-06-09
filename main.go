package main

import (
	"fmt"
	"net/http"
)

func main() {
	payload, err := http.Get("https://interview-data.herokuapp.com/companies")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", payload)
}
