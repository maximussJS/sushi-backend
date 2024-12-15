package utils

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.NewString()
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
