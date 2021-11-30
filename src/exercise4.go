
package main_
import "fmt"

type MyType int

var x MyType=10;
var y int = 20;
func main()	{
	fmt.Printf("\t %T \n", int(x));
	fmt.Println(x);
	fmt.Printf("type of y = %T \n", y);
	y = int(x);
	fmt.Printf("\nT = %T\n", y);
	fmt.Printf("x = %d, y = %d \n", x, y);
}

