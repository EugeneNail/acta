package http

import (
	"github.com/google/uuid"
	"github.com/samborkent/uuidv7"
	"net/http"
)

func (handler *Handler) Ping(writer http.ResponseWriter, request *http.Request) {
	message := "This is request " + uuid.UUID(uuidv7.New()).String()

	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write([]byte(message)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
