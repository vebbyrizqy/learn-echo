package utils

import "github.com/google/uuid"

var Sessions = map[string]map[string]interface{}{}

func GenerateSessionID() string {

	return uuid.NewString()
}