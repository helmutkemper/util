package util

import "fmt"

func ExampleRemoveSpellingAccents() {
	var text string = "áéíóúàèìòùâêîôûãẽĩõũäëïöüñçÁÉÍÓÚÀÈÌÒÙÂÊÎÔÛÃẼĨÕŨÄËÏÖÜÑÇ"

	fmt.Print(RemoveSpellingAccents(text))

	// Output:
	// aeiouaeiouaeiouaeiouaeiouncAEIOUAEIOUAEIOUAEIOUAEIOUNC
}

func ExampleRound() {

	fmt.Printf("%v\n", Round(1.22))
	fmt.Printf("%v\n", Round(1.22, 0.2))
	fmt.Printf("%v\n", Round(1.22, 0.5))
	fmt.Printf("%v\n", Round(1.22, 0.7))
	fmt.Printf("%v\n", Round(1.22, 0.5, 1.0))
	fmt.Printf("%v\n", Round(1.33, 0.5, 1.0))

	// Output:
	// 1
	// 2
	// 1
	// 1
	// 1.2
	// 1.3
}
