package main

import (
	"fmt"
)

func main() {
	name := "Pravin"
	learning := "Mastering Go"
	fmt.Println("\n Hi ", name, " You are learning", learning)

	changeCourse(&learning)
	fmt.Println("\n Hi ", name, " You are now learning", learning)
}

func changeCourse(course *string) string {
	*course = " First look at mastering Python"
	fmt.Println("\n Modified the course to", *course)
	return *course
}
