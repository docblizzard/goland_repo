package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(greet("fr"))
	var lang string
	flag.StringVar(&lang, "lang", "en", "Required language: en, fr...")
	//flag.Parse(*lang)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello World!",
	"fr": "Bonjour le monde!",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}
	return greeting
}
