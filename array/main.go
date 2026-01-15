package array
import "fmt"

func Array(){
	// binary search tree
	/*
	arr := []int{1,2,3,4,55,6,7}
	res := BinarySearch(arr, 3)
	fmt.Println(res)
	*/

	// zeroes move brute force
	/*
	arr2 := []int{1,2,0,4,0,6,7}
	MoveZeroes(arr2)
	fmt.Println(arr2)

	arr3 := []int{1,0,12,0,4,3,6,7}
	MoveZeroes1(arr3) //two pinter
	fmt.Println(arr3)
	*/
	arr3 := []int{1,0,12,0,4,3,6,7}
	MoveZeroes1(arr3) //two pinter
	fmt.Println(arr3)
}