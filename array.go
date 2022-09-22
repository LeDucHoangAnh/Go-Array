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

	//khai báo nhưng khởi tạo không đủ giá trị
	arrays2 := [3]int{100} //100 sẽ được gán vào phần tử đầu tiên
	fmt.Println(arrays2)   //[100 0 0]

	//size mang
	fmt.Println(len(arrays))
}
