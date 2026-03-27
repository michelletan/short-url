package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register endpoint")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login endpoint")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout endpoint")
}

// GET /me
func Me(w http.ResponseWriter, r *http.Request) {
	// TODO: get user from context (set by middleware)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User is authenticated",
	})
}
