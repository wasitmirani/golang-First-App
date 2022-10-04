// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界");
// 	test();
// }

// func test(){
// 	fmt.Println("hello world");
// }

package main

import "fmt"

 func main()  {

	fmt.Println("hello world first app on golang");
	veriables();
 }


 func veriables(){
	// var username string ="Wasitmirani";
	var link= "http://localhost";
	fmt.Println("hello world %T",link);
	num := 4;
	fmt.Println("num-",num);
     num1 := 664.66
	fmt.Println("num-",num1);
	sum  := num+int(num1);
	fmt.Println(sum);

 }