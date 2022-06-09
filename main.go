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
	ID       string `json:"id"`
	SurveyID string `json:"survey_id"` //!!
	Prompt   string `json:"prompt"`
}

type Response struct {
	ID         string `json:"id"`
	QuestionID string `json:"question_id"` // !!
	EmployeeID string `json:"employee_id"` //   !!!
	Score      int    `json:"score"`
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

// from either FS if ran prior, or web if FS cache data files arent present or otherwise err on os.Open
func hydrateServiceCache() {

	const fscachedir = `./data/`

	surveycache := fscachedir + "surveys"
	//const surveysURL etc could be handy wiring for that network/REST API/JSON+HTTP wiring. sounds like messing config. what to move to compile-time, init-time, run-time, is always a fun game of semantics.
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
	responsecache := fscachedir + "survey-responses"
	//const surveysURL etc could be handy wiring for that network/REST API/JSON+HTTP wiring. sounds like messing config. what to move to compile-time, init-time, run-time, is always a fun game of semantics.
	rdat, err := os.Open(responsecache)
	if err != nil {
		// if we do encounter an error, we can assume its... file unfound? :D happy-pathing here for now, adding ~syslog verbosity levels to point that out but defering any wiring heft given ~externalities.
		fmt.Println("INFO: couldnt load from cached files on fs, os.Open(%s)", responsecache)
		fmt.Printf("err when trying to load from FS cache: %s\n", err)

		response, err := http.Get("https://interview-data.herokuapp.com/surveys")
		if err != nil {
			panic(err)
		}
		responseStream := response.Body
		r = getResponses(responseStream)
		m, err := json.Marshal(s) //this feels like it could be so much better each second more i spend looking...
		os.WriteFile(responsecache, m, 0644)
		if err != nil {
			panic(err) //good memes from this idiomatic ~repetition
		}
	} else {
		responseStream := rdat
		fmt.Printf("INFO: loading from cached file %s\n", responsecache)
		r = getResponses(responseStream)
	}

	questioncache := fscachedir + "survey-questions"
	//const surveysURL etc could be handy wiring for that network/REST API/JSON+HTTP wiring. sounds like messing config. what to move to compile-time, init-time, run-time, is always a fun game of semantics.
	qdat, err := os.Open(questioncache)
	if err != nil {
		// if we do encounter an error, we can assume its... file unfound? :D happy-pathing here for now, adding ~syslog verbosity levels to point that out but defering any wiring heft given ~externalities.
		fmt.Println("INFO: couldnt load from cached files on fs, os.Open(%s)", questioncache)
		fmt.Printf("err when trying to load from FS cache: %s\n", err)

		response, err := http.Get("https://interview-data.herokuapp.com/survey-questions")
		if err != nil {
			panic(err)
		}
		questionStream := response.Body
		q = getQuestions(questionStream)
		m, err := json.Marshal(s) //this feels like it could be so much better each second more i spend looking...
		os.WriteFile(questioncache, m, 0644)
		if err != nil {
			panic(err) //good memes from this idiomatic ~repetition
		}
	} else {
		questionStream := qdat
		fmt.Printf("INFO: loading from cached file %s\n", questioncache)
		q = getQuestions(questionStream)
	}
}

// omniscient gopher? let alone, a bit WET
// too much power spider man... would be nice to externalize from main duties though. I want to check fs first anyways maybe? plenty of orthagonal cases which may be at-first unrelated but... you get the picture.
func init() {

	fmt.Printf("DEBUG: priming...")
	defer func(t time.Time) {
		fmt.Printf("DEBUG: primed (%v)\n", time.Now().Sub(t)) //time.Now()?
	}(time.Now())

	hydrateServiceCache()

}

func main() {
	fmt.Println("surveys in! get em while theyre hot :-)")
	fmt.Printf("%d surveys found\n", len(s))
	fmt.Printf("%d questions found\n", len(q))
	fmt.Printf("%d responses found\n", len(r))

	//this is the one we care about, wrt state-of-things in the livecoding session
	const someSurveyID = "17feca64-e756-4f15-beac-1dbbb293c227"

	// this funcmain() gophers ~internal state/context:
	var (
		qs []Question
		rs []Response

		//ii.
		activememory = make(map[string][]Response) // mapping questionID -> slice(list) of type Response
	)

	for _, v := range s {
		if v.ID == someSurveyID {
			fmt.Println("found it!")
			break
		} else {
			fmt.Printf("n")
		}
	}

	for _, question := range q {
		// given global questions q, for each question, is it one we care about?
		if question.SurveyID == someSurveyID {
			// yes, this one is
			qs = append(qs, question)
			activememory[question.ID] = make([]Response, 0)
			// this is heinous... easy low-hanging fruit though. would be great to benchmark.
			for _, response := range r {
				if response.QuestionID == question.ID {
					rs = append(rs, response)
					activememory[question.ID] = append(activememory[question.ID], response)
				}
			}
		}
	}

	fmt.Printf("local gophers ctx minima for summary on someSurveyID(%s) r[%d] q[%d]\n", someSurveyID, len(rs), len(qs))

	for k, v := range activememory {
		fmt.Printf("questionID[%v] ResponseCount[%v]\n", k, len(v))
		var localsum float64
		for _, v := range activememory[k] {
			localsum += float64(v.Score)
			// fmt.Printf("%v :: %s\n", v.Score, GetRating(v.Score).pp())
		}
		fmt.Printf("localavg[%v]\n", localsum/float64(len(activememory[k])))
	}

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
