// It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance.
// However, we can still implement the basic version of the pattern, the Simple Factory.

// In this example, we’re going to build various types of weapons using a factory struct.

// First, we create the iGun interface, which defines all methods a gun should have. There is a gun struct type that implements
// the iGun interface. Two concrete guns—ak47 and musket—both embed gun struct and indirectly implement all iGun methods.

// The gunFactory struct serves as a factory, which creates guns of the desired type based on an incoming argument.
// The main.go acts as a client. Instead of directly interacting with ak47 or musket, it relies on gunFactory to create
// instances of various guns, only using string parameters to control the production.

package main

import (
	"factory_methode_lib"
	"fmt"
)

func main() {
	ak47, _ := factory_methode_lib.GetGun("ak47")
	musket, _ := factory_methode_lib.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory_methode_lib.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
