package util

import "fmt"

func ExampleRound() {

	fmt.Printf("%v\n", Round(1.22))
	fmt.Printf("%v\n", Round(1.22, 0.2))
	fmt.Printf("%v\n", Round(1.22, 0.5))
	fmt.Printf("%v\n", Round(1.22, 0.7))
	fmt.Printf("%v\n", Round(1.22, 0.5, 1.0))
	fmt.Printf("%v\n", Round(1.33, 0.5, 1.0))
	fmt.Printf("%v\n", Round(3.14159265358979323, 0.5, 8.0))

	// Output:
	// 1
	// 2
	// 1
	// 1
	// 1.2
	// 1.3
	// 3.14159265
}
