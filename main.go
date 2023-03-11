package main

import "github.com/ivandi1980/my-go/pattern"

func main() {

	// Observer Pattern
	name := pattern.GetName()
	println(name)

	publisher := pattern.GetNewPublisher()
	subscriber1 := pattern.GetNewSubscriber("1")
	subscriber2 := pattern.GetNewSubscriber("2")
	publisher.Register(subscriber1)
	publisher.Register(subscriber2)

	publisher.NotifyAll("Hello, This is a message from Publisher")
}
