package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type Receiver interface {
	On()
	Off()
}

type Light struct {
	name  string
	color string
}

type TV struct {
	name  string
	brand string
}

type LightOnCommand struct {
	light *Light
}

type LightOffCommand struct {
	light *Light
}

type TVOnCommand struct {
	tv *TV
}

type TVOffCommand struct {
	tv *TV
}

func (l *Light) On() {
	fmt.Printf("%s light is on (Color: %s)\n", l.name, l.color)
}

func (l *Light) Off() {
	fmt.Printf("%s light is off (Color: %s)\n", l.name, l.color)
}

func (t *TV) On() {
	fmt.Printf("%s TV is on (Brand: %s)\n", t.name, t.brand)
}

func (t *TV) Off() {
	fmt.Printf("%s TV is off (Brand: %s)\n", t.name, t.brand)
}

func (c LightOnCommand) Execute() {
	c.light.On()
}

func (c LightOffCommand) Execute() {
	c.light.Off()
}

func (c TVOnCommand) Execute() {
	c.tv.On()
}

func (c TVOffCommand) Execute() {
	c.tv.Off()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{name: "Living Room", color: "White"}
	tv := &TV{name: "Family Room", brand: "Sony"}

	lightOn := LightOnCommand{light}
	lightOff := LightOffCommand{light}
	tvOn := TVOnCommand{tv}
	tvOff := TVOffCommand{tv}

	remote := RemoteControl{}

	for {
		fmt.Println("Enter a command (lighton, lightoff, tvon, tvoff, quit):")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "lighton":
			remote.SetCommand(lightOn)
			remote.PressButton()
		case "lightoff":
			remote.SetCommand(lightOff)
			remote.PressButton()
		case "tvon":
			remote.SetCommand(tvOn)
			remote.PressButton()
		case "tvoff":
			remote.SetCommand(tvOff)
			remote.PressButton()
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command. Try again.")
		}
	}
}
