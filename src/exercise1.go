package main

import (
	"fmt"
)

func main(){

	fmt.Println("Inside main function");
	PrintVars();
	fmt.Println("Printed variables and back to main again to get out of this program !!!");
}

func PrintVars()	{
		x := 10;
		y := "Sachin Mokashe";
		z := 29;

		fmt.Printf("Index : %d> name : %s, Age : %d \n", x, y, z);	
	}
