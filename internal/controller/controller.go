package controller

import "cafe_role/internal/service"

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}
