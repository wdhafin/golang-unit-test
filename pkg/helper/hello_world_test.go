package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dhafin")
	if result != "Hello Dhafin" {
		// error
		panic("result is not 'Hello Dhafin'")
	}
}
