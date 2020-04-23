package main

import(
	"fmt"
	"time"
)

func main(){

	for timer := 10; timer >=0; timer--{
		if timer  != 0 {
			fmt.Println(timer)
		} else {
			fmt.Println("Boom !")
		}
		time.Sleep(1*time.Second)
	}
}
