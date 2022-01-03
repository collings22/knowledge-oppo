package question

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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

func GetQuestionsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
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

func AddQuestionsToKnowledgeCheckCategoryHandler(w http.ResponseWriter, r *http.Request) {
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

func UpdateQuestionsToKnowledgeCheckCategoryHandler(w http.ResponseWriter, r *http.Request) {
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

func DeleteQuestionToKnowledgeCheckCategoryHandler(w http.ResponseWriter, r *http.Request) {
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
