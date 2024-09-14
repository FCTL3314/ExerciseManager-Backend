package main

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"fmt"
)

func main() {
	cfg, err := bootstrap.NewConfig()
	if err != nil {
		panic(err)
	}

	token, err := tokenutil.CreateAccessToken(&domain.User{ID: 5}, cfg.JWTSecret, 12)
	if err != nil {
		panic(err)
	}

	isAuthorized, err := tokenutil.IsAuthorized(token, cfg.JWTSecret)
	tokenID, err := tokenutil.ExtractIDFromToken(token, cfg.JWTSecret)

	fmt.Println(token)
	fmt.Println(isAuthorized)
	fmt.Println(tokenID)
}
