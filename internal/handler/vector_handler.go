package handler

import "go-chatbot/internal/controller"

type VectorHandler struct {
	vectorController controller.UserController
}

func NewVectorHandler(vectorController controller.UserController) *VectorHandler {
	return &VectorHandler{vectorController: vectorController}
}
