package main

import (
	"fmt"
)
func main(){
	number := []int{10,20,30,90};
	for i , num := range number {
		fmt.Println("index :", i , "value :" , num)
	}
     letter := []string{"a", "as", "er", "tg"}
	 for i, num:=range letter{
		fmt.Println("index :", i, "value :", num)
	 }

}
