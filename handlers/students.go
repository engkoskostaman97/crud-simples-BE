package handlers

import (
	"encoding/json"
	"net/http"
	dto "pelajar/dto/result"
	studentsdto "pelajar/dto/students"
	"pelajar/models"
	"pelajar/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handler struct {
	StudentRepository repositories.StudentRepository
}

// Create `path_file` Global variable here ...
var PATH_FILE = "http://localhost:5000/uploads/"

func HandlerStudent(StudentRepository repositories.StudentRepository) *handler {
	return &handler{StudentRepository}
}

func (h *handler) FindStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.StudentRepository.FindStudents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range users {
		users[i].Avatar = PATH_FILE + p.Avatar
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	student, err := h.StudentRepository.GetStudent(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student.Avatar = PATH_FILE + student.Avatar

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(student)}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	request := studentsdto.CreateStudentRequest{
		Avatar: r.FormValue("avatar"),
		Name:   r.FormValue("name"),
		Gender: r.FormValue("gender"),
		Dob:    r.FormValue("dob"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student := models.Student{
		Name:   request.Name,
		UserID: userId,
		Avatar: filename,
		Gender: request.Gender,
		Dob:    request.Dob,
	}

	student, err = h.StudentRepository.CreateStudent(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student, _ = h.StudentRepository.GetStudent(student.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: student}
	json.NewEncoder(w).Encode(response)
}
func (h *handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get student id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get image filename
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	request := studentsdto.UpdateStudentRequest{
		Name:   r.FormValue("name"),
		Gender: r.FormValue("gender"),
		Dob:    r.FormValue("dob"),
		UserID: userId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student, _ := h.StudentRepository.GetStudent(id)

	student.Name = request.Name
	student.Gender = request.Gender
	student.Dob = request.Dob

	if filename != "false" {
		student.Avatar = filename
	}

	student, err = h.StudentRepository.UpdateStudent(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: student}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.StudentRepository.GetStudent(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.StudentRepository.DeleteStudent(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.Student) studentsdto.StudentResponse {
	return studentsdto.StudentResponse{
		ID:     u.ID,
		Name:   u.Name,
		Avatar: u.Avatar,
		Gender: u.Gender,
		Dob:    u.Dob,
	}
}
