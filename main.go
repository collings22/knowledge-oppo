package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	questionService "github.com/collings22/knowledge-oppo/question"

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

	routesHandler(r)

	r.Use(mux.CORSMethodMiddleware(r))

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))

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

func routesHandler(r *mux.Router) {
	r.HandleFunc("/health-check", HealthCheck).Methods("GET")

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/knowledge-check/{category}", questionService.GetQuestionsByCategoryHandler).Methods("GET")
	r.HandleFunc("/knowledge-check/{category}/questions", questionService.AddQuestionsToKnowledgeCheckCategoryHandler).Methods("POST")
	r.HandleFunc("/knowledge-check/{category}/questions", questionService.UpdateQuestionsToKnowledgeCheckCategoryHandler).Methods("PUT")
	r.HandleFunc("/knowledge-check/{category}/questions/{id}", questionService.DeleteQuestionToKnowledgeCheckCategoryHandler).Methods("DELETE")

	r.HandleFunc("/knowledge-check/{category}/score", KnowledgeCheckScoreHandler).Methods("POST")
}
