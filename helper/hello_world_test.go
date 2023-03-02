package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// ini fungsi dasar benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Willi")
	}
}

// ini fungsi sub benchmark
func BenchmarkHelloWorldSubBench(b *testing.B) {
	b.Run("Willi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Willi")
		}
	})
	b.Run("Buli", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Buli")
		}
	})
}

func BenchmarkHelloWorldTableBench(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Ama)",
			request: "Ama",
		},
		{
			name:    "HelloWorld(Willi)",
			request: "Willi",
		},
		{
			name:    "HelloWorld(Buli)",
			request: "Buli",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// proses ini adalah before dan after
func TestMain(m *testing.M) {
	fmt.Println("Before")
	m.Run()
	fmt.Println("After")
}

func TestHelloWorld(t *testing.T) {
	t.Skip("TestHelloWorld is Skip") // ini untuk skip test

	// proses test dibawah ini adalah proses test dasar yang saya sarankan
	result := HelloWorld("Willi")
	require.Equal(t, "Hello Willi", result, "Result is not 'Hello Willi'")
	fmt.Println("TestHelloWorld Done")
}

// proses ini adalah subtest
func TestSubTest(t *testing.T) {
	t.Run("Willi", func(t *testing.T) {
		result := HelloWorld("Willi")
		require.Equal(t, "Hello Willi", result, "Result is not 'Hello Willi'")
	})
	t.Run("Buli", func(t *testing.T) {
		result := HelloWorld("Buli")
		require.Equal(t, "Hello Buli", result, "Result is not 'Hello Buli'")
	})
}

// ini adalah proses tabel test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Ama)",
			request:  "Ama",
			expected: "Hello Ama",
		},
		{
			name:     "HelloWorld(Willi)",
			request:  "Willi",
			expected: "Hello Willi",
		},
		{
			name:     "HelloWorld(Buli)",
			request:  "Buli",
			expected: "Hello Buli",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

	fmt.Println("TestHelloWorld Done")
}
