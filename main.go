package main

import (
	"fmt"
	"github.com/ivandi1980/my-go/pattern"
)

func main() {

	/**
	 * Observer Pattern
	 * It is a design pattern that defines a one-to-many dependency relationship between objects,
	 * where changes to the state of one object (called the subject) are automatically propagated
	 * to all its dependents (called observers)
	 */

	name := pattern.GetObserverName("Observer")
	println("----------------------------------")
	println(name)
	println("----------------------------------")

	publisher := pattern.GetNewPublisher()
	subscriber1 := pattern.GetNewSubscriber("1")
	subscriber2 := pattern.GetNewSubscriber("2")
	publisher.Register(subscriber1)
	publisher.Register(subscriber2)

	publisher.NotifyAll("Hello, This is a message from Publisher")

	/**
	 * Prototype Pattern
	 * a creation design pattern that allows creating new objects by copying or cloning existing objects.
	 * It provides a way to create new objects based on an existing object, without specifying the exact class
	 * of the object that will be created.
	 */

	name = pattern.GetPrototypeName("Prototype")
	println("----------------------------------")
	println(name)
	println("----------------------------------")

	pattern.LoadToRegistry()

	square := pattern.RegistryList[int(pattern.SquareType)]
	sq, ok := square.(pattern.Square) // Type Assertion
	if ok {
		fmt.Println("Old Properties : ")
		sq.PrintTypeProp()
		newSquare := sq.Clone() // Prototyping
		fmt.Println("Cloned Object Properties : ")
		newSquare.PrintTypeProp()
	}

	circle := pattern.RegistryList[int(pattern.CircleType)]
	cr, ok := circle.(pattern.Circle) // Type Assertion
	_ = ok
	if ok {
		fmt.Println("Old Properties : ")
		cr.PrintTypeProp()
		newCircle := cr.Clone().(pattern.Circle) // Prototyping i,e. Cloning Existing Object
		newCircle.Radius = 100
		fmt.Println("Cloned Object Change Properties : ")
		newCircle.PrintTypeProp()
	}

	/**
	 * SingletonPattern
	 * a design pattern that restricts the instantiation of a class to one object.
	 * This is useful when exactly one object is needed to coordinate actions across the system.
	 */

	name = pattern.GetSingletonName("Singleton")
	println("----------------------------------")
	println(name)
	println("----------------------------------")

	for i := 0; i < 3; i++ {
		go pattern.GetConfigInstance()
	}
	fmt.Scanln()
}
