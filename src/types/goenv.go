package main

import ("fmt"; "runtime")

func main()	{
	fmt.Println("Go OStype = ", runtime.GOOS);
	fmt.Println("Go Architecture : ", runtime.GOARCH);
}	
