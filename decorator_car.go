package main

import "fmt"

type Car interface {
	Drive() string
	Cost() float64
}

type BaseCar struct{}

func (c *BaseCar) Drive() string {
	return "Driving a basic car"
}

func (c *BaseCar) Cost() float64 {
	return 20000.0
}

type Decorator struct {
	car Car
}

func (d *Decorator) Drive() string {
	return d.car.Drive()
}

func (d *Decorator) Cost() float64 {
	return d.car.Cost()
}

type GPSDecorator struct {
	Decorator
}

func NewGPSDecorator(car Car) *GPSDecorator {
	return &GPSDecorator{Decorator{car}}
}

func (d *GPSDecorator) Drive() string {
	return d.car.Drive() + " with GPS"
}

func (d *GPSDecorator) Cost() float64 {
	return d.car.Cost() + 500.0
}

type LeatherSeatsDecorator struct {
	Decorator
}

func NewLeatherSeatsDecorator(car Car) *LeatherSeatsDecorator {
	return &LeatherSeatsDecorator{Decorator{car}}
}

func (d *LeatherSeatsDecorator) Drive() string {
	return d.car.Drive() + " with leather seats"
}

func (d *LeatherSeatsDecorator) Cost() float64 {
	return d.car.Cost() + 1000.0
}

type AlloyWheelsDecorator struct {
	Decorator
}

func NewAlloyWheelsDecorator(car Car) *AlloyWheelsDecorator {
	return &AlloyWheelsDecorator{Decorator{car}}
}

func (d *AlloyWheelsDecorator) Drive() string {
	return d.car.Drive() + " with alloy wheels"
}

func (d *AlloyWheelsDecorator) Cost() float64 {
	return d.car.Cost() + 800.0
}

func main() {
	baseCar := &BaseCar{}
	fmt.Printf("Base Car: %s, Cost: $%.2f\n", baseCar.Drive(), baseCar.Cost())

	gpsCar := NewGPSDecorator(baseCar)
	fmt.Printf("GPS Car: %s, Cost: $%.2f\n", gpsCar.Drive(), gpsCar.Cost())

	leatherSeatsCar := NewLeatherSeatsDecorator(gpsCar)
	fmt.Printf("Leather Seats Car: %s, Cost: $%.2f\n", leatherSeatsCar.Drive(), leatherSeatsCar.Cost())

	alloyWheelsCar := NewAlloyWheelsDecorator(leatherSeatsCar)
	fmt.Printf("Alloy Wheels Car: %s, Cost: $%.2f\n", alloyWheelsCar.Drive(), alloyWheelsCar.Cost())
}
