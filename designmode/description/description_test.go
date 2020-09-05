package description

import (
	"fmt"
	"testing"
)

func TestDescription(t *testing.T) {
	testDrink := drink{}
	testDrink.cost = 1

	coffeeDrink := coffee{&testDrink}
	fmt.Println(coffeeDrink.getCost() == testDrink.cost +coffeeCost)

	milkDrink:= milk{&testDrink}
	fmt.Println(milkDrink.getCost() == testDrink.cost +milkCost)
}