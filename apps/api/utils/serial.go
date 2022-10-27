package utils

import "github.com/google/uuid"

func CreateSerial(prefix string) string {
	return prefix + "-" + uuid.NewString()
}
