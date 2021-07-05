package calculate

import (
	"fmt"
	"testing"
)

func TestEven(t *testing.T) {
	inputs := []struct{a int; b bool}{
		{2, true},
		{3, false},
		{4, true},
	}

	for _, data := range inputs {

		if result := Even(data.a); result != data.b {
			t.Errorf("Even(%d), expect result is %v, actual is %v\n", data.a, data.b, result)
		}
	}
}

func BenchmarkEven(b *testing.B) {

	input := struct {a int; b bool}{
		2, true,
	}

	for i := 0; i < b.N; i++ {

		if result := Even(input.a); result != input.b {
			b.Errorf("Even(%d), expect result is %v, actual is %v\n", input.a, input.b, result)
		}
	}
}

func ExampleEven() {
	fmt.Println(Even(2))
	// ouput:
	// true
}