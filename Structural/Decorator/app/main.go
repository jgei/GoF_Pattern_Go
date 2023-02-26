// Decorator is a structural pattern that allows adding new behaviors to objects dynamically by placing them inside
// special wrapper objects, called decorators.

// Using decorators you can wrap objects countless number of times since both target objects and decorators follow the
// same interface. The resulting object will get a stacking behavior of all wrappers.

package main

import (
	"decorator_lib"
	"fmt"
)

func main() {

	pizza := &decorator_lib.VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator_lib.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator_lib.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
