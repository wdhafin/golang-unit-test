package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Dhafin", func(t *testing.T) {
		result := HelloWorld("Dhafin")
		require.Equal(t, "Hello Dhafin", result, "result must be 'Hello Dhafin'")
	})
	t.Run("Rizqullah", func(t *testing.T) {
		result := HelloWorld("Rizqullah")
		require.Equal(t, "Hello Rizqullah", result, "result must be 'Hello Rizqullah'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("before unit test")

	m.Run()

	fmt.Println("after unit test")
	// after
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on linux")
	}
	result := HelloWorld("Dhafin")
	require.Equal(t, "Hello Dhafin", result, "result must be 'Hello Dhafin'")
}

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
