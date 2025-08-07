package utils
import"github.com/google/uuid"

func CreateNewID() string {
	return uuid.New().String()
}