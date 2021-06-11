package Order_Application

import (
	orderRepository "FoodInHouse/Order/Domain"
)

func Create(order orderRepository.Order) error {

	err := orderRepository.Create(order)

	if err != nil {
		return err
	}

	return nil
}

func Read() (orderRepository.Orders, error) {

	orders, err := orderRepository.Read()

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func Update(order orderRepository.Order, OrderId string) error {

	err := orderRepository.Update(order, OrderId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(OrderId string) error {

	err := orderRepository.Delete(OrderId)

	if err != nil {
		return err
	}

	return nil
}
