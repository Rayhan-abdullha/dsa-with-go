package array
import "fmt"

func Array(){
	// binary search tree
	arr := []int{1,2,3,4,55,6,7}
	res := BinarySearch(arr, 3)
	fmt.Println(res)
}