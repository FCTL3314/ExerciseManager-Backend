package main

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/token_util"
	"fmt"
)

func main() {
	cfg, err := bootstrap.NewConfig()
	if err != nil {
		panic(err)
	}

	token, err := token_util.CreateAccessToken(&domain.User{ID: 1}, cfg.JWTSecret, 12)
	if err != nil {
		panic(err)
	}

	isAuthorized, err := token_util.IsAuthorized(token, cfg.JWTSecret)
	tokenID, err := token_util.ExtractIDFromToken(token, cfg.JWTSecret)

	fmt.Println(token)
	fmt.Println(isAuthorized)
	fmt.Println(tokenID)
}
