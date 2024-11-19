package controller

import (
	"go-chatbot/internal/service"
)

type VectorController interface {
}

type vectorController struct {
	vectorService service.VectorService
}

func NewVectorController(vectorService service.VectorService) VectorController {
	return &vectorController{vectorService: vectorService}
}
