package main

import (
	"fmt"
	"strings"
)

func validateCredentials(username, password string) bool {
	// Validasi panjang username dan password
	if len(username) < 6 || len(password) < 8 {
		return false
	}
	// Validasi apakah password mengandung username
	if strings.Contains(password, username) {
		return false
	}
	return true
}

func main() {
	username := "user123"
	password := "password123"
	if validateCredentials(username, password) {
		fmt.Println("Username dan password valid")
	} else {
		fmt.Println("Username atau password tidak valid")
	}
}
