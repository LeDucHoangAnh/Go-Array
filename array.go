package main

import "fmt"

func main() {
	//Array : Mảng
	var myArray [4]int
	fmt.Println(myArray)

	//gán
	myArray[0] = 5
	fmt.Println(myArray)
	myArray[3] = 10
	fmt.Println(myArray)

	//khai báo và khởi tạo giá trị
	arrays := [3]int{1, 2, 3}
	fmt.Println(arrays)
}
