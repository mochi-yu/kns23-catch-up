package handler

type Handler struct {
	Health HealthHandler
	Post   PostHandler
	User   UserHandler
}
