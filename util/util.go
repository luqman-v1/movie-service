package util

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

func Env(path ...string) (err error) {
	err = godotenv.Load(path...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return
}

func ItoS(input int) string {
	return strconv.Itoa(input)
}

func StoI(input string) int32 {
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Println("error convert string to int", err)
	}
	return int32(i)
}
