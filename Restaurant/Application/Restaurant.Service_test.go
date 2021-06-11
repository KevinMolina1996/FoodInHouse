package Restaurant_Application_test

import (
	"testing"
	"time"

	resService "FoodInHouse/Restaurant/Domain"
)

func TestCreate(t *testing.T) {

	restaurant := resService.Restaurant{
		Name:      "COMIDAS RAPIDAS KM",
		CreatedAt: time.Now(),
		Address:   "CALLE SIEMPRE VIVA 123",
	}

	err := resService.Create(restaurant)

	if err != nil {
		t.Error("La prueba fallo")
		t.Fail()
	} else {
		t.Log("Prueba correcta")
	}
}
