package main

import "fmt"

type OldSystem interface {
	OldMethod() string
}

type NewSystem interface {
	NewMethod() string
}

type OldImplementation struct{}

func (o *OldImplementation) OldMethod() string {
	return "Hello from Old System"
}

type Adapter struct {
	oldSystem OldSystem
}

func (a *Adapter) NewMethod() string {
	return a.oldSystem.OldMethod()
}

func main() {
	oldSystem := &OldImplementation{}

	adapter := &Adapter{oldSystem: oldSystem}

	fmt.Println(adapter.NewMethod())
}
