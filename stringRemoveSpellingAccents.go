package util

import (
	"strings"
)

func StringRemoveSpellingAccents(text string) string {
	var search []string = []string{"á", "à", "ã", "â", "ä", "é", "è", "ẽ", "ê", "ë", "í", "ì", "î", "ĩ", "ï", "ó", "ò", "ô", "õ", "ö", "ú", "ù", "û", "ũ", "ü", "ç", "ñ", "Á", "À", "Ã", "Â", "Ä", "É", "È", "Ẽ", "Ê", "Ë", "Í", "Ì", "Î", "Ĩ", "Ï", "Ó", "Ò", "Ô", "Õ", "Ö", "Ú", "Ù", "Û", "Ũ", "Ü", "Ç", "Ñ"}
	var replace []string = []string{"a", "a", "a", "a", "a", "e", "e", "e", "e", "e", "i", "i", "i", "i", "i", "o", "o", "o", "o", "o", "u", "u", "u", "u", "u", "c", "n", "A", "A", "A", "A", "A", "E", "E", "E", "E", "E", "I", "I", "I", "I", "I", "O", "O", "O", "O", "O", "U", "U", "U", "U", "U", "C", "N"}

	for k, v := range search {
		text = strings.Replace(text, v, replace[k], -1)
	}

	return text
}
