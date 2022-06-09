package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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

type Question struct {
	ID       string
	SurveyID string //!!
	Prompt   string
}

type Response struct {
	ID         string
	QuestionID string // !!
	EmployeeID string //   !!!
	Score      int
}

type Companies []Company

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	s []Survey
	q []Question
	r []Response
)

// omniscient gopher? let alone, a bit WET
// too much power spider man... would be nice to externalize from main duties though. I want to check fs first anyways maybe? plenty of orthagonal cases which may be at-first unrelated but... you get the picture.
func init() {
	fmt.Printf("DEBUG: priming...")
	defer func(t time.Time) {
		fmt.Printf("DEBUG: primed (%v)\n", time.Now().Sub(t)) //time.Now()?
	}(time.Now())

	const fscachedir = `./data/`

	surveycache := fscachedir + "surveys"
	dat, err := os.Open(surveycache)
	if err != nil {
		// if we do encounter an error, we can assume its... file unfound? :D happy-pathing here for now, adding ~syslog verbosity levels to point that out but defering any wiring heft given ~externalities.
		fmt.Println(`INFO: couldnt load from cached files on fs, os.Open("./data/surveys")`)
		fmt.Printf("err when trying to load from FS cache: %s\n", err)

		response, err := http.Get("https://interview-data.herokuapp.com/surveys")
		if err != nil {
			panic(err)
		}
		surveyStream := response.Body
		s = getSurveys(surveyStream)
		m, err := json.Marshal(s) //this feels like it could be so much better each second more i spend looking...
		os.WriteFile("data/surveys", m, 0644)
		if err != nil {
			panic(err) //good memes from this idiomatic ~repetition
		}
	} else {
		surveyStream := dat
		fmt.Println(`INFO: loading from cached file ./data/surveys`)
		s = getSurveys(surveyStream)

	}
	// fmt.Printf("in init: got surveys[%#v]\n", s)
	// res, err := http.Get("https://interview-data.herokuapp.com/survey-questions")
	// if err != nil {
	// 	panic(err)
	// }
	// questionStream := res.Body
	// q = getQuestions(questionStream)
	// fmt.Printf("q[%#v]\n", q)

	// res, err := http.Get("https://interview-data.herokuapp.com/survey-responses")
	// if err != nil {
	// 	panic(err)
	// }
	// responseStream := res.Body
	// r = getResponses(responseStream)
	// fmt.Printf("r[%#v]\n", r)
}

func main() {
	fmt.Println("surveys in! get em while theyre hot :-)")
	fmt.Printf("%d surveys found\n", len(s))
}

func getSurveys(in io.Reader) []Survey {
	var (
		err error
		s   []Survey

		dec = json.NewDecoder(in)
	)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var tmp Survey
		// decode an array value (Message)
		err = dec.Decode(&tmp)
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, tmp)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return s
}

func getQuestions(in io.Reader) []Question {
	var (
		err error
		q   []Question

		dec = json.NewDecoder(in)
	)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var tmp Question
		// decode an array value (Message)
		err = dec.Decode(&tmp)
		if err != nil {
			log.Fatal(err)
		}
		q = append(q, tmp)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return q
}

func getResponses(in io.Reader) []Response {
	var (
		err error
		r   []Response

		dec = json.NewDecoder(in)
	)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var tmp Response
		// decode an array value (Message)
		err = dec.Decode(&tmp)
		if err != nil {
			log.Fatal(err)
		}
		r = append(r, tmp)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return r
}
