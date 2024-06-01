package main

// import (
// 	"golang.org/x/tour/wc"
// )
import "strings"

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, word := range strings.Fields(s) {
		res[word]++
	}
	return res
}

// func main() {
// 	wc.Test(WordCount)
// }
