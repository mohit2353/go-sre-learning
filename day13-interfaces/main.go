package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct{}
type Bike struct{}
type Truck struct{}

func (b Bike) Drive() string {
	return "Bike Started"
}

func (c Car) Drive() string {
	return "Car Started"
}

func (t Truck) Drive() string {
	return "Truck Started"
}

func main() {

	for _, v := range []Vehicle{Car{}, Bike{}, Truck{}} {
		fmt.Println(v.Drive())
	}
}
