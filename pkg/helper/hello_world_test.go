package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHellowWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Dhafin",
			request:  "Dhafin",
			expected: "Hello Dhafin",
		},
		{
			name:     "Rizqullah",
			request:  "Rizqullah",
			expected: "Hello Rizqullah",
		},
		{
			name:     "Muhammad",
			request:  "Muhammad",
			expected: "Hello Muhammad",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

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

func TestHelloWorldDhafin(t *testing.T) {
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
