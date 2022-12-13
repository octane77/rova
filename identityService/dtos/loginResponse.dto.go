package dtos

import "github.com/octane77/rova/identityService/entities"

type LoginResponseDto struct {
	User  entities.User `json:"user"`
	Token string        `json:"token"`
}
