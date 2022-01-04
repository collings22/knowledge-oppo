package routes

import (
	"github.com/collings22/knowledge-oppo/controllers/health"
	"github.com/collings22/knowledge-oppo/controllers/question"
	"github.com/collings22/knowledge-oppo/controllers/score"
	"github.com/gorilla/mux"
)

func RoutesHandler(r *mux.Router) {
	r.HandleFunc("/health-check", health.HealthCheck).Methods("GET")

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/knowledge-check/{category}", question.GetQuestionsByCategoryHandler).Methods("GET")
	r.HandleFunc("/knowledge-check/{category}/questions", question.AddQuestionsToKnowledgeCheckCategoryHandler).Methods("POST")
	r.HandleFunc("/knowledge-check/{category}/questions", question.UpdateQuestionsToKnowledgeCheckCategoryHandler).Methods("PUT")
	r.HandleFunc("/knowledge-check/{category}/questions/{id}", question.DeleteQuestionToKnowledgeCheckCategoryHandler).Methods("DELETE")

	r.HandleFunc("/knowledge-check/{category}/score", score.KnowledgeCheckScoreHandler).Methods("POST")
}
