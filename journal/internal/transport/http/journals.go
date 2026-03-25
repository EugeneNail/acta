package http

import (
	"net/http"

	"github.com/EugeneNail/acta/journal/internal/transport/http/middleware"
	"github.com/google/uuid"
)

type JournalsResponse struct {
	UserUuid uuid.UUID `json:"userUuid"`
}

// GetJournals handles journal list requests.
func (handler *Handler) GetJournals(request *http.Request) (int, any) {
	userUUID := middleware.UserUUID(request.Context())

	return http.StatusOK, JournalsResponse{
		UserUuid: userUUID,
	}
}
