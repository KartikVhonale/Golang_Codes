package main

import "fmt"

//outside of any function we can't use : this operator for declaring variable
//like a := 100;
func main() {
	//string
	var username string = "kartik"
	fmt.Println(username)
	fmt.Printf("the type of variable : %T \n\n", username)

	//bool
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("the type of variable : %T \n\n", isLoggedIn)

	//small int val
	var small_uint uint8 = 255
	fmt.Println(small_uint)
	fmt.Printf("the type of variable : %T \n\n", small_uint)

	//small float value
	var small_float float32 = 255.4343434343434
	fmt.Println("value of small float :", small_float)
	fmt.Printf("the type of veriable : %T \n\n", small_float)

	//small float value
	var big_float float64 = 255.4343434343434
	fmt.Println("value of float64 :", big_float)
	fmt.Printf("the type of veriable : %T \n\n", big_float)

	//declare default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("the type of veriable : %T \n\n", anotherVariable)

	//implicit type
	var gmail = "whonalekartik11@gmail.com"
	fmt.Println(gmail)
	fmt.Printf("the type of veriable : %T \n\n", gmail)

	//no var style for skiping var
	users := 3000000
	fmt.Println(users)
	fmt.Printf("the type of veriable : %T \n\n", users)

	//declaring constant
	const Login = "12345678976543456789"
	fmt.Println(Login)
	fmt.Printf("the type of veriable : %T \n\n", Login)

}
