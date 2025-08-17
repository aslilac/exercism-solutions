package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
    greeting := greeter.Greet(name)
    return fmt.Sprintf("I can speak %v: %v", greeter.LanguageName(), greeting)
}

type Italian struct {}

func (self Italian) LanguageName() string {
    return "Italian"
}

func (self Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %v!", name)
}

type Portuguese struct {}

func (self Portuguese) LanguageName() string {
    return "Portuguese"
}

func (self Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %v!", name)
}