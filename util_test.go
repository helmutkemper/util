package util

import "fmt"

func ExampleRemoveSpellingAccents() {
  var text string = "áéíóúàèìòùâêîôûãẽĩõũäëïöüñçÁÉÍÓÚÀÈÌÒÙÂÊÎÔÛÃẼĨÕŨÄËÏÖÜÑÇ"

  fmt.Print( RemoveSpellingAccents( text ) )

  // Output:
  // aeiouaeiouaeiouaeiouaeiouncAEIOUAEIOUAEIOUAEIOUAEIOUNC
}
