package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/health-check", HealthCheck).Methods("GET")

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/knowledge-check/{category}", KnowledgeCheckHandler).Methods(http.MethodGet)
	r.HandleFunc("/knowledge-check/{category}/score", KnowledgeCheckScoreHandler).Methods(http.MethodPost)
	r.Use(mux.CORSMethodMiddleware(r))

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))

}

func KnowledgeCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tmp := []question{}
	for _, item := range questions {
		if item.Category == params["category"] {
			tmp = append(tmp, item)
		}
	}
	json.NewEncoder(w).Encode(&tmp)
}

func KnowledgeCheckScoreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	decoder := json.NewDecoder(r.Body)
	var answers []answer
	err := decoder.Decode(&answers)

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	score := 0
	for _, ans := range answers {
		for _, cAns := range correctAnswers {
			if ans.ID == cAns.ID && cAns.Category == params["category"] && ans.Answer == cAns.Answer {
				score++
			}
		}
	}

	pct := (float64(score) / float64(len(answers))) * 100

	json.NewEncoder(w).Encode(&pct)
}

type question struct {
	ID       string   `json:"id"`
	Label    string   `json:"label"`
	Author   string   `json:"author"`
	Category string   `json:"category"`
	Options  []string `json:"options"`
}

var questions = []question{
	{
		ID:       "1",
		Label:    "Who didnt score in 2004 Carling Cuo Final?",
		Author:   "Dan Coltrane",
		Category: "Boro",
		Options:  []string{"Zenden", "Job", "Southgate"},
	},
	{
		ID:       "2",
		Label:    "Who did score in 2004 Carling Cuo Final?",
		Author:   "Dan Coltrane",
		Category: "Boro",
		Options:  []string{"Zenden", "Juninho", "Southgate"},
	},
	{
		ID:       "3",
		Label:    "Which ex-Barca goalkeeper played for Boro?",
		Author:   "Dan Coltrane",
		Category: "Boro",
		Options:  []string{"Turnbull", "Jones", "Valdes"},
	},
	{
		ID:       "4",
		Label:    "Who managed Boro to Euros Final?",
		Author:   "Dan Coltrane",
		Category: "Boro",
		Options:  []string{"Wilder", "Warnock", "McClaren"},
	},
	{
		ID:       "5",
		Label:    "Who played Thor in the MCU?",
		Author:   "Kyle Coltrane",
		Category: "Marvel",
		Options:  []string{"Tony Stark", "Chris Hemsworth", "Barry Southgate"},
	},
	{
		ID:       "6",
		Label:    "Which Spiderman actor made the first appearance?",
		Author:   "Dan Coltrane",
		Category: "Marvel",
		Options:  []string{"Tom", "Andrew", "Toby"},
	},
}

type answer struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Answer   string `json:"answer"`
}

var correctAnswers = []answer{
	{
		ID:       "1",
		Category: "Boro",
		Answer:   "Southgate",
	},
	{
		ID:       "2",
		Category: "Boro",
		Answer:   "Zenden",
	},
	{
		ID:       "3",
		Category: "Boro",
		Answer:   "Valdes",
	},
	{
		ID:       "4",
		Category: "Boro",
		Answer:   "McClaren",
	},
	{
		ID:       "5",
		Category: "Marvel",
		Answer:   "Chris Hemsworth",
	},
	{
		ID:       "6",
		Category: "Marvel",
		Answer:   "Toby",
	},
}
