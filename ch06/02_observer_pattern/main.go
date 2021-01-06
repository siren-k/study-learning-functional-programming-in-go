package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch06/02_observer_pattern/observer"
)

func main() {
	subject := Subject{}

	oa := Observable{Name: "A"}
	ob := Observable{Name: "B"}
	subject.AddObserver(&Observer{})
	subject.NotifyObservers(oa, ob)
	fmt.Println("-----")

	oc := Observable{Name: "C"}
	subject.NotifyObservers(oa, ob, oc)
	fmt.Println("-----")

	// If you remove an observer from the Service Locator by executing
	// subject.DeleteObserver(&Observer{}), the observer who will
	// respond to the event is considered, and subsequent notifications
	// are meaningless.
	subject.DeleteObserver(&Observer{})
	subject.NotifyObservers(oa, ob, oc)
	fmt.Println("-----")

	od := Observable{Name: "D"}
	subject.NotifyObservers(oa, ob, oc, od)
}
