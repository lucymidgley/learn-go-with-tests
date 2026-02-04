package iteration

import "strings"

//	func Repeat(x string) (repeated string) {
//		for i := 0; i < 5; i++ {
//			repeated += x
//		}
//		return
//	}

func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
