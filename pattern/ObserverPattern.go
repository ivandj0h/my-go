package pattern

import "fmt"

type Publisher interface {
	Register(subscriber Subscriber)
	NotifyAll(msg string)
}

type Subscriber interface {
	ReactToPublisherMsg(msg string)
}

type publisher struct {
	subscriberList []Subscriber
}

func GetNewPublisher() publisher {
	return publisher{subscriberList: make([]Subscriber, 0)}
}

func (p *publisher) Register(subs Subscriber) {
	p.subscriberList = append(p.subscriberList, subs)
}

func (p publisher) NotifyAll(msg string) {
	for _, subs := range p.subscriberList {
		fmt.Println("Publisher Notifying Subscriber with Id: ", subs.(subscriber).subscriberId)
		subs.ReactToPublisherMsg(msg)
	}
}

type subscriber struct {
	subscriberId string
}

func GetNewSubscriber(Id string) subscriber {
	return subscriber{subscriberId: Id}
}

func (s subscriber) ReactToPublisherMsg(msg string) {
	fmt.Println("Subscriber with Id: ", s.subscriberId, " received message: ", msg)
}

func GetName() string {
	return "Observer Pattern"
}
