package http

type Handler struct{}

// NewHandler constructs the HTTP handler set.
func NewHandler() *Handler {
	return &Handler{}
}
