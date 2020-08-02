package util

import "fmt"

func ExampleStringRemoveSpellingAccents() {
	var text string = "áéíóúàèìòùâêîôûãẽĩõũäëïöüñçÁÉÍÓÚÀÈÌÒÙÂÊÎÔÛÃẼĨÕŨÄËÏÖÜÑÇ"

	fmt.Print(StringRemoveSpellingAccents(text))

	// Output:
	// aeiouaeiouaeiouaeiouaeiouncAEIOUAEIOUAEIOUAEIOUAEIOUNC
}
