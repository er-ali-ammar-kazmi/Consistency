package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

type middleware struct {
	logger  *os.File
	handler http.Handler
}

func (mdlwr middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	if r.Method == http.MethodPost && strings.Contains(r.URL.Path, "login") {
	} else if r.Method == http.MethodGet {
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			fmt.Fprintf(w, "No validation mechanism provided, jwt_token missing!")
			return
		}

		tokenString := cookie.Value
		claims, err := ValidateToken(tokenString)
		if err != nil {
			fmt.Fprintf(w, "Validation mechanism failed, wrong jwt token!")
			return
		}

		email := claims["user_email"]
		mdlwr.logger.WriteString(fmt.Sprintf("%v : %s[%s] - %s\n", now, r.Method, email, r.URL.Path))
	} else {
		fmt.Fprintf(w, "either wrong method type or wrong URI, please check!")
		return
	}
	rateLimiter := RateLimiter(r)
	if rateLimiter.Allow() {
		mdlwr.handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "Application/Json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Too many requests!"})
		return
	}
}

func NewMiddleWare(mx *http.ServeMux) *middleware {

	file, err := os.Create("logger.log")
	if err != nil {
		log.Println("error forming a connection with log file: ", err.Error())
	}

	return &middleware{logger: file,
		handler: mx}
}

func NewToken(user User) (any, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.Id,
		"user_email": user.Email,
		"expiry":     time.Now().Add(time.Hour * 2).Unix(),
	})

	secret := os.Getenv("SECRET")
	auth, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func ValidateToken(tokenString string) (map[string]any, error) {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s", "Signing method not verified")
		} else {
			secret := []byte(os.Getenv("SECRET"))
			err := method.Verify(strings.Join(strings.Split(tokenString, ".")[:2], "."), t.Signature, secret)
			if err != nil {
				return nil, err
			}
			return secret, nil
		}
	})
	if err != nil {
		return nil, fmt.Errorf("validation failed: %s", err.Error())
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("%s", "error Occured at jwt token validation")
}

func RateLimiter(r *http.Request) rate.Limiter {
	limit := rate.Limit(2)
	burst := 10
	ipMap := sync.Map{}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println(err)
	}
	limiter, exists := ipMap.Load(host)
	if !exists {
		limiter = rate.NewLimiter(limit, burst)
		ipMap.Store(host, limiter)
	}
	ratelimiter, _ := limiter.(rate.Limiter)
	return ratelimiter
}
