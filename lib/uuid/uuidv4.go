package uuid

import "github.com/google/uuid"

func GenUuidv4() string {
	return uuid.New().String()
}
