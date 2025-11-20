package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type middleware struct {
	logger  *os.File
	handler http.Handler
}

func (mdlwr middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	cookie, err := r.Cookie("auth")
	if err != nil {
		fmt.Fprintf(w, "No validation mechanism provided, jwt token missing!")
	}
	tokenString := cookie.Raw
	claims, err := ValidateToken(tokenString)
	if err != nil {
		fmt.Fprintf(w, "Validation mechanism failed, wrong jwt token!")
	}
	email := claims["user_email"]
	mdlwr.logger.WriteString(fmt.Sprintf("%v : %s[%s] - %s\n", now, r.Method, email, r.URL))
	mdlwr.handler.ServeHTTP(w, r)
}

func NewMiddleWare(mx *http.ServeMux) *middleware {

	file, err := os.Create("logger.log")
	if err != nil {
		log.Println("error forming a connection with log file: ", err.Error())
	}

	return &middleware{logger: file,
		handler: mx}
}

func NewToken(user User) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.Id,
		"user_email": user.Email,
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	})

	secret := os.Getenv("SECRET")
	auth, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(auth)
	return auth
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
