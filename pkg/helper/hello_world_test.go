package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dhafin")
	require.Equal(t, "Hello Dhafin", result, "result must be 'Hello Dhafin'")
	fmt.Println("TestHelloWorld with require done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Dhafin")
	assert.Equal(t, "Hello Dhafin", result, "result must be 'Hello Dhafin'")
	fmt.Println("TestHelloWorld with assert done")
}

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
