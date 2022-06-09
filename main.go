package main

import (
	"fmt"
	"net/http"
)

// appdx
// https://interview-data.herokuapp.com/surveys
// https://interview-data.herokuapp.com/survey-questions
// https://interview-data.herokuapp.com/survey-responses

func main() {
	payload, err := http.Get("https://interview-data.herokuapp.com/companies")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", payload)
}
