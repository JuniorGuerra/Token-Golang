package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for i := 0; i < 1000; i++ {
		v1 := GenerateToken("NombreAleatorio")
		v2 := GenerateToken("NombreAleatorio")

		if v1 == v2 {
			fmt.Println("token: ", v1)
			fmt.Println("token: ", v2)
		}
	}
}

func GenerateToken(name string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(name), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store: ", string(hash))
	return base64.StdEncoding.EncodeToString(hash)
}
