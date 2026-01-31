package main

import (
	"fmt"
)

const (
	helloEn = "Hello, "
	helloSp = "Hola, "
	helloJp = "こんにちは, "
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "sp":
		prefix = helloSp
	case "jp":
		prefix = helloJp
	default:
		prefix = helloEn
	}
	return
}

func Hello(s, lang string) string {
	if s == "" {
		s = "world"
	}
	return greetingPrefix(lang) + s
}

func main() {
	fmt.Println(Hello("name", ""))
}
