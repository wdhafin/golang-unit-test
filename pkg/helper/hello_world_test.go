package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dhafin")
	if result != "Hello Dhafin" {
		// error
		t.Error("Result must be 'Hello Dhafin'")
	}
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldRizqullah(t *testing.T) {
	result := HelloWorld("Rizqullah")
	if result != "Hello Rizqullah" {
		// error
		t.Fatal("Result must be 'Hello Rizqullah'")
	}
	fmt.Println("TestHelloWorldRizqullah done")
}
