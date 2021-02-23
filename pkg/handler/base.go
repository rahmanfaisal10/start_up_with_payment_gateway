package handler

import (
	"bwastartup/auth"
	"bwastartup/pkg/service"
)

type handler struct {
	service       service.Service
	authorization auth.Auth
}

func InitHandler(service service.Service, authorization auth.Auth) *handler {
	return &handler{service, authorization}
}
