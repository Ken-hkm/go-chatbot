package service

import (
	"go-chatbot/internal/dto"
	"go-chatbot/internal/repository"
	"strconv"
)

type OrderService interface {
	GetUserByID(id uint) ([]dto.OrderDataDto, error)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{orderRepo: orderRepo}
}

func (s *orderService) GetUserByID(id uint) ([]dto.OrderDataDto, error) {
	//return nil, errors.New("test")
	return s.orderRepo.GetOrderDataByUserId(strconv.Itoa(int(id)))
}
