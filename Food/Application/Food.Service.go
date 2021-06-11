package Food_Application

import (
	foodRepository "FoodInHouse/Food/Domain"
)

func Create(food foodRepository.Food) error {

	err := foodRepository.Create(food)

	if err != nil {
		return err
	}

	return nil
}

func Read() (foodRepository.Foods, error) {

	foods, err := foodRepository.Read()

	if err != nil {
		return nil, err
	}

	return foods, nil
}

func ReadById(id string) (foodRepository.Foods, error) {

	foods, err := foodRepository.ReadById(id)

	if err != nil {
		return nil, err
	}

	return foods, nil
}

func ReadTypes() (foodRepository.Types, error) {

	types, err := foodRepository.ReadTypes()

	if err != nil {
		return nil, err
	}

	return types, nil
}

func Update(food foodRepository.Food, FoodId string) error {

	err := foodRepository.Update(food, FoodId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(FoodId string) error {

	err := foodRepository.Delete(FoodId)

	if err != nil {
		return err
	}

	return nil
}
