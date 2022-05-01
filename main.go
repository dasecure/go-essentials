package main

import (
	"fmt"
)

type Ordered interface {
	int | string | float64
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func main() {
	fmt.Println(min([]int{1, 2, 3, 4, 5}))
	fmt.Println(min([]string{"f", "b", "A", "d", "a"}))
	fmt.Println(min([]float64{1.1, 2.2, 3.3, 4.4, 5.5, 0.1}))
}
