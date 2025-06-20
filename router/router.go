package router

import (
	"go-21/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	// r.Use(r.Middlewares()....Logger)

	r.Post("/login", handler.AuthHandler.Login)

	// r.Get("/student/home", handler.AssignmentHandler.ListAssignments)
	// r.Get("/student/submit", assignmentHandler.ShowSubmitForm)
	// r.Post("/student/submit", assignmentHandler.SubmitAssignment)

	// r.Get("/lecturer/home", submissionHandler.Home)
	// r.Get("/lecturer/grade-form", submissionHandler.ShowGradeForm)
	// r.Post("/lecturer/grade", submissionHandler.GradeSubmission)
	return r
}
