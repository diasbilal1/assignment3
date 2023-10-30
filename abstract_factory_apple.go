package main

import "fmt"

type AppleFactory interface {
	CreatePhone() ApplePhone
	CreateLaptop() AppleLaptop
}

type ApplePhone interface {
	GetInfo() string
}

type AppleLaptop interface {
	GetInfo() string
}

type iPhone struct {
	Model      string
	Generation int
}

func (i iPhone) GetInfo() string {
	return fmt.Sprintf("iPhone %s (Generation: %d)", i.Model, i.Generation)
}

type MacBook struct {
	Model        string
	Year         int
	RAM          int
	Storage      string
	Screen       string
	Speakers     int
	BassSpeakers int
}

func (m MacBook) GetInfo() string {
	return fmt.Sprintf("MacBook %s (Year: %d)\nRAM: %d GB\nStorage: %s\nScreen: %s\nSpeakers: %d (%d with bass)",
		m.Model, m.Year, m.RAM, m.Storage, m.Screen, m.Speakers, m.BassSpeakers)
}

type iPhoneFactory struct{}

func (f iPhoneFactory) CreatePhone() ApplePhone {
	return iPhone{
		Model:      "iPhone 13",
		Generation: 2021,
	}
}

func (f iPhoneFactory) CreateLaptop() AppleLaptop {
	return nil
}

type MacBookFactory struct{}

func (f MacBookFactory) CreatePhone() ApplePhone {
	return nil
}

func (f MacBookFactory) CreateLaptop() AppleLaptop {
	return MacBook{
		Model:        "M1 Pro MacBook Pro",
		Year:         2023,
		RAM:          16,
		Storage:      "1 terabyte",
		Screen:       "High-quality screen",
		Speakers:     6,
		BassSpeakers: 4, // 60% of the speakers
	}
}

type Application struct{}

func (app Application) CreateAndDisplayDevices(factory AppleFactory) {
	phone := factory.CreatePhone()
	laptop := factory.CreateLaptop()

	if phone != nil {
		fmt.Println("Created Phone:", phone.GetInfo())
	}

	if laptop != nil {
		fmt.Println("Created Laptop:", laptop.GetInfo())
	}
}

func main() {
	app := Application{}

	iphoneFactory := iPhoneFactory{}
	macbookFactory := MacBookFactory{}

	app.CreateAndDisplayDevices(iphoneFactory)
	app.CreateAndDisplayDevices(macbookFactory)
}
