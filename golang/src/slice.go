package main

import(
	"fmt"
)

func main(){

	mySlice :=[]int{5,10,15,20,25}
	fmt.Println(" \n Initial value ",mySlice)
	fmt.Println(" \n ") 

	for _,i := range mySlice{
		fmt.Println(" Iterating through slice:", i)
	}
	newSlice :=[]int{100,200,300,400}
	mySlice = append(mySlice, newSlice...)

	fmt.Println(" \n Final value ",mySlice)
}
