package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "bgp"
	author := "yakov rekther"
	bestFinish := bestLeagueFinishes(12, 10, 12, 17, 18, 15)
	fmt.Println(bestFinish)

	fmt.Println(converter(module, author))
}

func converter(module, author string) (str1, str2 string) {
	str1 = strings.Title(module)
	str2 = strings.ToUpper(author)
	return str1, str2
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i > best {
			best = i
		}
	}
	return best
}
