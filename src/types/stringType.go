package main

import ("fmt")


func main()	{
	var str string
	str = "Sachin Mokashe"
	
	var len int = len(str);
	for index:=0; index < len; index++	{
		fmt.Printf("%c", str[index]);
	}
	var str1[] byte = []byte(str);	
	for index := 0; index < len; index++	{
		fmt.Printf("%x \t", str1[index]);
	}

	fmt.Println();
}
