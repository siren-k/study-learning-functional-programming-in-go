package observer

import "fmt"

type Observable struct {
	Name string
}

type Observer struct{}

func (ob *Observer) Notify(o *Observable) {
	fmt.Println(o.Name)
}

type Callback interface {
	Notify(o *Observable)
}
