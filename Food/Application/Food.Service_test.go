package Food_Application_test

import (
	"testing"
	"time"

	foodService "FoodInHouse/Food/Domain"
)

func TestCreate(t *testing.T) {

	food := foodService.Food{
		Name:       "FILETE DE POLLO APANADO",
		Price:      7000,
		Type:       "",
		Restaurant: "60c2abb1912e689fbdb05e41",
		CreatedAt:  time.Now(),
	}

	err := foodService.Create(food)

	if err != nil {
		t.Error("La prueba fallo")
		t.Fail()
	} else {
		t.Log("Prueba correcta")
	}
}

func TestRead(t *testing.T) {
	food, err := foodService.ReadById("60c2abb1912e689fbdb05e41")

	if err != nil {
		t.Error("Se presentó un error")
		t.Fail()
	}

	if len(food) == 0 {
		t.Error("No hay usuarios")
		t.Fail()
	} else {
		t.Log("Hay registros")
	}
}
