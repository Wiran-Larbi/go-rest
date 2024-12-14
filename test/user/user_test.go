package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pts/mdes/service/user"
	"github.com/pts/mdes/types"
)

func TestUserServiceHandlers(t *testing.T) {

	t.Run("Should fail if the user payload is invalid ..", func(t *testing.T) {
		userService := &user.MockUserService{}
		handler := user.NewHandler(userService)
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "resu",
			Email:     "",
			Password:  "",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
		}

	})

	t.Run("Should correctly register the user ..", func(t *testing.T) {
		userService := &user.MockUserService{}
		handler := user.NewHandler(userService)
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "admin",
			Email:     "user.admin@gmail.com",
			Password:  "password123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, recorder.Code)
		}

	})
}
