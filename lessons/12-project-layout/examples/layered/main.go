package main

import (
	"log/slog"
	"net/http"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/handler"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/middleware"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/repository"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/service"
)

func main() {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	slog.Info("layered server on :8083")
	http.ListenAndServe(":8083", middleware.Logging(mux))
}
