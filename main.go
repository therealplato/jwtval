package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: jwtval <token> <signing key>")
	}
	tokenString := os.Args[1]
	signKey := []byte(os.Args[2])
	token, err := jwt.Parse(tokenString, makeKeyFunc(signKey))
	if err != nil {
		log.Fatal(err)
	}
	bb, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		fmt.Println(token)
		os.Exit(0)
	}
	fmt.Println(string(bb))
}

func makeKeyFunc(signKey []byte) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return signKey, nil
	}
}
