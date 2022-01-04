package score

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
