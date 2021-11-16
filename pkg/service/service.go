package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mjudeikis/backend-service/pkg/models"
	"github.com/mjudeikis/backend-service/pkg/utils/recover"
	"go.uber.org/zap"
)

var _ Service = &ApiService{}

type Service interface {
	Run(ctx context.Context, stop <-chan struct{}, done chan<- struct{}) error
}

type ApiService struct {
	log    *zap.Logger
	server *http.Server
	router *mux.Router

	store map[string]models.UserInternal // our unique store index is email
}

func New(log *zap.Logger) (*ApiService, error) {

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	})

	s := &ApiService{
		log:    log,
		router: r,
		store:  make(map[string]models.UserInternal),
	}

	s.router.HandleFunc("/user", s.registerUser).Methods("POST")

	s.server = &http.Server{
		Addr: ":8080",
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"Content-Type"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}),
		)(s.router),
	}

	return s, nil
}

func (s *ApiService) Run(ctx context.Context, stop <-chan struct{}, done chan<- struct{}) error {
	s.log.Info("Starting API Service")

	go func() {
		defer recover.Panic(s.log)

		<-stop
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err := s.server.Shutdown(ctx)
		if err != nil {
			s.log.Error("api shutdown error", zap.Error(err))
		}
		defer close(done)
	}()

	s.log.Info("Server will now listen", zap.String("url", ":8080"))
	return s.server.ListenAndServe()
}
