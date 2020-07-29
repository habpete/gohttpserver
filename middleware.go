package main

import (
	"net/http"
)

// func RequireAuthentification (next http.Handler) http.Handler {
// 	return http.HandlerFunc(func (w http.ResponseWriter, r* http.Request) {})
// }

func RequireAuthentification (next http.Handler) {}