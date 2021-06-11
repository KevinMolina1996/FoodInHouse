package Restaurant_Application

import (
	restaurantRepository "FoodInHouse/Restaurant/Domain"
)

func Create(restaurant restaurantRepository.Restaurant) error {

	err := restaurantRepository.Create(restaurant)

	if err != nil {
		return err
	}

	return nil
}

func Read() (restaurantRepository.Restaurants, error) {

	restaurants, err := restaurantRepository.Read()

	if err != nil {
		return nil, err
	}

	return restaurants, nil
}

func Update(restaurant restaurantRepository.Restaurant, RestaurantId string) error {

	err := restaurantRepository.Update(restaurant, RestaurantId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(RestaurantId string) error {

	err := restaurantRepository.Delete(RestaurantId)

	if err != nil {
		return err
	}

	return nil
}
