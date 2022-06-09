package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// appdx
// https://interview-data.herokuapp.com/surveys
// https://interview-data.herokuapp.com/survey-questions
// https://interview-data.herokuapp.com/survey-responses

// ex.A
// /surveys
//	{
//     "id": "17feca64-e756-4f15-beac-1dbbb293c227",
//     "company_id": "88419c6e-4c6e-4021-971a-9ba0ad76c3c5",
//     "name": "Survey A"
//   },

// ex.B
// /companies
//   {
//     "id": "88419c6e-4c6e-4021-971a-9ba0ad76c3c5",
//     "name": "JPMorgan Chase"
//   },

type Survey struct {
	ID        string
	CompanyID string //type company? feels like it def should be, but still plenty of known unknowns
	Name      string
}
type Companies []Company

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	s := getSurveys()
	fmt.Printf("in main: got surveys[%#v]\n", s)
}

func getSurveys() []Survey {
	var s []Survey
	res, err := http.Get("https://interview-data.herokuapp.com/surveys")
	if err != nil {
		panic(err)
	}

	dec := json.NewDecoder(res.Body)
	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var m Survey
		// decode an array value (Message)
		err = dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, m)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	return s
}
