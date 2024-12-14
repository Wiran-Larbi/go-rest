package auth

import (
	"testing"

	"github.com/pts/mdes/service/auth"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPwd, err := auth.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if len(hashedPwd) == 0 {
		t.Errorf("Hashed password is empty !")
	}

	if hashedPwd == "password" {
		t.Error("Hashed password is not hashed !")
	}
}

func TestComparePassword(t *testing.T) {
	password := "password"
	hashedPwd, err := auth.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !auth.ComparePassword(hashedPwd, password) {
		t.Error("Password does not match hashed password !")
	}

	if auth.ComparePassword(hashedPwd, "wrongpassword") {
		t.Error("Password matches wrong password !")
	}
}
