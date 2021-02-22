package handler

import (
	"bwastartup/pkg/service"
)

type handler struct {
	service service.Service
}

func InitHandler(service service.Service) *handler {
	return &handler{service}
}
