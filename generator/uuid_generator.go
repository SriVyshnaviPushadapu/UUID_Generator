package generator

import "github.com/google/uuid"

func GenerateUUID() string{
	return uuid.New().String()
}

func GenerateCustomUUID() string{
	return "Custom" + uuid.New().String()
}