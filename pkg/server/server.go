package server

import (
	"context"
	"github.com/auctionee/front/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type AuthServer struct {
	Ctx               context.Context
	port              string
	connectionTimeout time.Duration
}

func NewAuthServer(port int) *AuthServer {
	ctx := logger.NewCtxWithLogger()
	return &AuthServer{
		port: ":" + strconv.Itoa(port),
		Ctx:  ctx,
	}
}
func (s *AuthServer) Start() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("../static"))
	r.PathPrefix("/").Handler(fs)
	//postRouter := r.Methods(http.MethodPost).Subrouter()
	//postRouter.Handle("/get/", handlers.LoggerMiddleware(s.Ctx, http.HandlerFunc(handlers.GetHandler)))
	//postRouter.Handle("/modify/", handlers.LoggerMiddleware(s.Ctx, http.HandlerFunc(handlers.ModifyHandler)))

	if err := http.ListenAndServe(s.port, r); err != nil {
		logger.GetLogger(s.Ctx).Fatal(err)
	}
}
