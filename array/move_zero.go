package array

func MoveZeroes(nums []int) {
	n := len(nums)
	result := make([]int, n)
	j := 0

	// First pass: accumulate non-zero elements
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			result[j] = nums[i]
			j++
		}
	}
	
	// Second pass: fill remaining positions with zeroes
	for ; j < n; j++ {
		result[j] = 0
	}

	// Copy back to original slice
	for i := 0; i < n; i++ {
		nums[i] = result[i]
	}
}

// two pointer
func MoveZeroes1(nums []int) {
	writePos := 0 

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[writePos] = nums[i]
			writePos++
		}
	}


	for writePos < len(nums) {
		nums[writePos] = 0
		writePos++
	}
}

func MoveZeroes2(nums []int) {
	writePos := 0 // next slot for a non-zero
	for readPos := 0; readPos < len(nums); readPos++ {
		if nums[readPos] != 0 {
			if readPos != writePos {
				nums[readPos], nums[writePos] = nums[writePos], nums[readPos]
			}
			writePos++
		}
	}
}