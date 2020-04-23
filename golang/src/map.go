package main
import(
		"fmt"
      )
func main(){

testMap := map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5}

	for key, value := range testMap{
		fmt.Printf("\n Key is : %v Value is: %v \n", key,value)
	}

testMap["A"] = 100
testMap["C"] =1973

fmt.Println("\n")
fmt.Println(testMap)
delete(testMap, "C")

}
