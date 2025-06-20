package handler

import (
	"encoding/json"
	"go-21/dto"
	"go-21/service"
	"go-21/utils"
	"net/http"
)

type AuthHandler struct {
	Service service.Service
}

func NewAuthHandler(service service.Service) AuthHandler {
	return AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var LoginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&LoginRequest); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// fmt.Printf("data %s", LoginRequest.Email)

	user, err := h.Service.AuthService.Login(LoginRequest.Email, LoginRequest.Password)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error())
		return
	}

	// // Simpan user_id dan role ke cookie
	// http.SetCookie(w, &http.Cookie{
	// 	Name:  "user_id",
	// 	Value: strconv.Itoa(user.ID),
	// 	Path:  "/",
	// })

	// http.SetCookie(w, &http.Cookie{
	// 	Name:  "user_role",
	// 	Value: user.Role,
	// 	Path:  "/",
	// })

	utils.ResponseSuccess(w, http.StatusOK, "Login success", user)
}
