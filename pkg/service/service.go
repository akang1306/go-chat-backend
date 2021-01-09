package service

type database interface {
}

// Handler provides the interface to handle different controller requests
type Handler struct {
	db database
}

func NewService() Handler {
	return Handler{}
}
