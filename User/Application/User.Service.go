package User_Application

import (
	userRepository "FoodInHouse/User/Domain"
)

func Create(user userRepository.User) error {

	err := userRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (userRepository.Users, error) {

	Users, err := userRepository.Read()

	if err != nil {
		return nil, err
	}

	return Users, nil
}

func Update(User userRepository.User, UserId string) error {

	err := userRepository.Update(User, UserId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(UserId string) error {

	err := userRepository.Delete(UserId)

	if err != nil {
		return err
	}

	return nil
}
