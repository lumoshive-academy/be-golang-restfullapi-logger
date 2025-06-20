// handler/lecturer_handler.go
package handler

import (
	"fmt"
	"go-21/service"
	"net/http"
	"strconv"
)

type SubmissionHandler struct {
	Service service.Service
}

func NewSubmissionHandler(service service.Service) SubmissionHandler {
	return SubmissionHandler{
		Service: service,
	}
}

func (h *SubmissionHandler) Home(w http.ResponseWriter, r *http.Request) {
	submissions, err := h.Service.SubmissionService.GetAllSubmissions()
	if err != nil {
		http.Error(w, "Gagal mengambil data submission", http.StatusInternalServerError)
		return
	}
	fmt.Println(submissions)
}

func (h *SubmissionHandler) ShowGradeForm(w http.ResponseWriter, r *http.Request) {
	studentIDStr := r.URL.Query().Get("student_id")
	assignmentIDStr := r.URL.Query().Get("assignment_id")

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		http.Error(w, "Invalid student_id", http.StatusBadRequest)
		return
	}

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		http.Error(w, "Invalid assignment_id", http.StatusBadRequest)
		return
	}

	// Ambil data untuk ditampilkan di form
	student, err := h.Service.UserService.GetUserByID(studentID)
	if err != nil {
		http.Error(w, "Student not found", http.StatusInternalServerError)
		return
	}

	assignment, err := h.Service.AssignmentService.GetAssignmentByID(assignmentID)
	if err != nil {
		http.Error(w, "Assignment not found", http.StatusInternalServerError)
		return
	}

	data := struct {
		StudentID       int
		AssignmentID    int
		StudentName     string
		AssignmentTitle string
	}{
		StudentID:       student.ID,
		AssignmentID:    assignment.ID,
		StudentName:     student.Name,
		AssignmentTitle: assignment.Title,
	}

	fmt.Println(data)
}

func (h *SubmissionHandler) GradeSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Gagal parsing form", http.StatusBadRequest)
		return
	}

	studentID, err := strconv.Atoi(r.FormValue("student_id"))
	if err != nil {
		http.Error(w, "Invalid student_id", http.StatusBadRequest)
		return
	}

	assignmentID, err := strconv.Atoi(r.FormValue("assignment_id"))
	if err != nil {
		http.Error(w, "Invalid assignment_id", http.StatusBadRequest)
		return
	}

	gradeStr := r.FormValue("grade")
	grade, err := strconv.ParseFloat(gradeStr, 64)
	if err != nil {
		http.Error(w, "Invalid grade", http.StatusBadRequest)
		return
	}

	err = h.Service.SubmissionService.GradeSubmission(studentID, assignmentID, grade)
	if err != nil {
		http.Error(w, "Gagal memberi nilai: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/lecturer/home", http.StatusSeeOther)
}
