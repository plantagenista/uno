package uno

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s !!", name)
}

type Counter struct {
	Value int
}

func (c *Counter) Inc()    { c.Value++ }
func NewCounter() *Counter { return &Counter{5} }

type Printer interface {
	Print(s string) string
}

func PrintHello(p Printer) string {
	return p.Print("Hello, World!")
}
