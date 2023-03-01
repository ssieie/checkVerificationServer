package middleware

import (
	X "checkVerification/utils"
	"net/http"
	"regexp"
)

func Cross(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "uuid")
		w.Header().Set("content-type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		match, _ := regexp.MatchString(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`, request.Header.Get("uuid"))

		if !match {
			_, _ = writer.Write(X.JSON(X.Z{
				"code":    http.StatusUnauthorized,
				"message": "Invalid request",
			}))
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func HttpHandler(next http.Handler) http.Handler {
	return Cross(Verify(next))
}
