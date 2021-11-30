package main
import ("fmt")

var x int=100 
var y = 20
var a bool
var z = 10;
func main(){
	var b int=50;

	fmt.Println("value of X = ", x);
	fmt.Println("value of y = ", y);
	fmt.Println("Value of a = ", a);
	myFun();
	fmt.Println("Value of z = ", z);
	fmt.Println("Value of b = ", b);
}

func myFun()	{
	fmt.Println("Inside myFun...!");
}
