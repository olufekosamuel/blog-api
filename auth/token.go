package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(email string, id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := make(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 8760).Unix() //Token expires after 1 year

	token.Claims = claims
	return token.SignedString([]byte("nZr4u7x!z%C*F-JaNdRgUkXp2s5v8y/B?D(G+KbPeShVmYq3t6w9z$C&F)H@McQf"))
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("nZr4u7x!z%C*F-JaNdRgUkXp2s5v8y/B?D(G+KbPeShVmYq3t6w9z$C&F)H@McQf"), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenId(r *http.Request) (float64, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("nZr4u7x!z%C*F-JaNdRgUkXp2s5v8y/B?D(G+KbPeShVmYq3t6w9z$C&F)H@McQf"), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["id"].(float64), nil
	}
	return 0, nil
}

func ExtractTokenEmail(r *http.Request) (string, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("nZr4u7x!z%C*F-JaNdRgUkXp2s5v8y/B?D(G+KbPeShVmYq3t6w9z$C&F)H@McQf"), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["email"].(string), nil
	}
	return "", nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	_, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	//fmt.Println(string(b))
}
