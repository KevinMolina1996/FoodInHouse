package User_Application_test

import (
	"testing"
	"time"

	userService "FoodInHouse/User/Domain"
)

func TestCreate(t *testing.T) {

	user := userService.User{
		Name:      "Kevin",
		Email:     "kevin.andres125@hotmail.com",
		Address:   "Cra 4D 48 136",
		CreatedAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("La prueba fallo")
		t.Fail()
	} else {
		t.Log("Prueba correcta")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()

	if err != nil {
		t.Error("Se presentó un error")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("No hay usuarios")
		t.Fail()
	} else {
		t.Log("Hay registros")
	}
}
