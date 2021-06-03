package httphandler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type Handler struct {
	Users map[string]string
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/public", h.Public)
	mux.HandleFunc("/authorize", h.Authorize)

	mux.Handle("/restricted", h.Restrict(func(tc *TokenClaim) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Welcome,", tc.Name)
		})
	}))

	mux.ServeHTTP(rw, req)
}

func (h Handler) Public(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is public endpoint"))
}

func (h Handler) Authorize(w http.ResponseWriter, req *http.Request) {
	// restrictions...
	if req.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, `{"message":"Method Not Allowed"}`)
		return
	}

	userIncoming := req.FormValue("user")
	passIncoming := req.FormValue("pass")

	// check values
	for user, pass := range h.Users {
		if user == userIncoming && pass == passIncoming {
			tokenString, err := createToken(user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Sorry, error while Signing Token!")
				// log.Printf("Token Signing error: %v\n", err)
				return
			}

			// w.Header().Set("Content-Type", "application/jwt")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{ "token": "%s" }`, tokenString)
			return
		}
	}
	// giving up
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w, "i don't know you")
	return
}

func (h Handler) Restrict(tch func(tc *TokenClaim) http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Authorization") == "" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		token, err := request.ParseFromRequestWithClaims(req, request.OAuth2Extractor, &TokenClaim{}, verifyTokenWithKey)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(rw, "Invalid toke1n:", err)
			return
		}

		tch(token.Claims.(*TokenClaim)).ServeHTTP(rw, req)
	})
}

func verifyTokenWithKey(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}

func createToken(user string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	token.Claims = &TokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"root",
		user,
	}

	return token.SignedString(signKey)
}
