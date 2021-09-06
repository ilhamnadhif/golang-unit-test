package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// go test
// go test -v
// go test -v -run (nama function)
// go test -v ./...  // untuk di root folder

type TableHelloWorld struct {
	Name     string
	Request  string
	Expected string
}

type TableBenchmark struct {
	Name    string
	Request string
}

func BenchmarkTable(b *testing.B) {
	var benchmarks []TableBenchmark = []TableBenchmark{
		{Name: "Muhammad", Request: "Muhammad"},
		{Name: "Ilham", Request: "Ilham"},
		{Name: "Nadhif", Request: "Nadhif"},
	}
	for _, bench := range benchmarks {
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.Request)
			}
		})
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Ilham", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ilham")
		}
	})
	b.Run("Nadhif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nadhif")
		}
	})
}

func BenchmarkHelloWorldIlham(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ilham")
	}
}

func BenchmarkHelloWorldNadhif(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nadhif")
	}
}

func TestTableTest(t *testing.T) {
	tests := []TableHelloWorld{
		{
			Name:     "HelloWorld(muhammad)",
			Request:  "Muhammad",
			Expected: "Hello Muhammad",
		},
		{
			Name:     "HelloWorld(ilham)",
			Request:  "Ilham",
			Expected: "Hello Ilham",
		},
		{
			Name:     "HelloWorld(nadhif)",
			Request:  "Nadhif",
			Expected: "Hello Nadhif",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("ilham", func(t *testing.T) {
		result := HelloWorld("Ilham")
		assert.Equal(t, "Hello Ilhamm", result, "Harusnya Hello Ilham")
	})
	t.Run("nadhif", func(t *testing.T) {
		result := HelloWorld("Nadhif")
		assert.Equal(t, "Hello Nadhif", result, "Harusnya Hello Nadhif")
	})
}
func TestMain(m *testing.M) {
	fmt.Println("-----BEFORE UNIT TEST-----")
	m.Run()
	fmt.Println("-----AFTER UNIT TEST-----")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on Mac OS")
	}
	// kode ke bawah tidak dijalankan
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Asser", result, "Result Must be Hello Assert")
	fmt.Println("Test Hello World with Assert Done")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result Must be Hello Require")
	fmt.Println("Test Hello World with Require Done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result Must be Hello Assert")
	fmt.Println("Test Hello World with Assert Done")
}
func TestHelloWorldIlham(t *testing.T) {
	result := HelloWorld("Ilham")
	if result != "Hello Ilham" {
		// error
		//t.Fail()
		t.Error("Harusnya Hello Ilham")
	}
	fmt.Println("Test Hello World Done")
}

func TestHelloWorldNadhif(t *testing.T) {
	result := HelloWorld("Nadhif")
	if result != "Hello Nadhif" {
		// error
		//t.FailNow()
		t.Fatal("Harusnya Hello Nadhif")
	}
	fmt.Println("Test Hello World Nadhif Done")
}
