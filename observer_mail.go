package main

import (
	"fmt"
)

type Observer interface {
	Update(temperature float64)
}

type Subject struct {
	observers   []Observer
	temperature float64
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Deregister(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.temperature)
	}
}

func (s *Subject) SetTemperature(temperature float64) {
	s.temperature = temperature
	s.Notify()
}

type WeatherObserver struct {
	id string
}

func (o *WeatherObserver) Update(temperature float64) {
	fmt.Printf("[%s] Temperature updated: %.2f\n", o.id, temperature)
}

func main() {
	weatherStation := &Subject{}

	observer1 := &WeatherObserver{id: "Observer 1"}
	observer2 := &WeatherObserver{id: "Observer 2"}

	weatherStation.Register(observer1)
	weatherStation.Register(observer2)

	weatherStation.SetTemperature(25.5)
	weatherStation.SetTemperature(30.0)

	weatherStation.Deregister(observer2)

	weatherStation.SetTemperature(32.5)

	observer3 := &WeatherObserver{id: "Observer 3"}
	observer4 := &WeatherObserver{id: "Observer 4"}

	weatherStation.Register(observer3)
	weatherStation.Register(observer4)

	weatherStation.SetTemperature(28.0)
	weatherStation.SetTemperature(29.5)

	weatherStation.Deregister(observer1)

	weatherStation.SetTemperature(27.0)
}
