// The Builder pattern is used when the desired product is complex and requires multiple steps to complete.
// In this case, several construction methods would be simpler than a single monstrous constructor.
// The potential problem with the multistage building process is that a partially built and unstable product
// may be exposed to the client. The Builder pattern keeps the product private until it’s fully built.

// In the below code, we see different types of houses (igloo and normalHouse) being constructed by iglooBuilder
// and normalBuilder. Each house type has the same construction steps. The optional director struct helps to
// organize the building process.

package main

import (
	"builder_lib"
	"fmt"
)

func main() {
	normalBuilder := builder_lib.GetBuilder("normal")
	iglooBuilder := builder_lib.GetBuilder("igloo")

	director := builder_lib.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

}
