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

	//size mảng:sử dụng len()
	fmt.Println(len(arrays))

	//**khai báo mảng không cần set size**
	arrays3 := [...]string{"Merc", "Porch", "Toyo"}
	fmt.Println(arrays3)

	//array là value type không phải là reference type**
	//value type
	//reference type:
	countries := [...]string{"VN", "US", "UK", "JP"}
	copyCountries := countries // countries1 = countries,  copyCountries = countries1

	// a->(A) b = a <-> b -> (A) b = xyx

	copyCountries[0] = "Canada"
	fmt.Println(countries)
	fmt.Println(copyCountries)
	fmt.Println(countries)
	//Loop array
	//C1
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
	//C2
	for index, value := range countries {
		fmt.Printf("i = %d value = %s ", index, value)
	}
	//Blank-identifier = _

	//Mang 2 chieu [][]

	matrix := [4][2]int{
		{1, 2},
		{10, 22},
		{15, 25},
		{13, 27},
	}

	fmt.Println(matrix)
}
