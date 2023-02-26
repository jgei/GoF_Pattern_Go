// Abstract Factory defines an interface for creating all distinct products but leaves the actual product
// creation to concrete factory classes. Each factory type corresponds to a certain product variety.

// The client code calls the creation methods of a factory object instead of creating products directly with
// a constructor call (new operator). Since a factory corresponds to a single product variant, all its products
// will be compatible.

// Client code works with factories and products only through their abstract interfaces. This lets the client
// code work with any product variants, created by the factory object. You just create a new concrete factory
// class and pass it to the client code.

package main

import (
	"abstract_factory_lib"
	"fmt"
)

func main() {
	adidasFactory, _ := abstract_factory_lib.GetSportsFactory("adidas")
	nikeFactory, _ := abstract_factory_lib.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s abstract_factory_lib.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s abstract_factory_lib.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
