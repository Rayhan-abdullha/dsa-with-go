package array

func BinarySearch(a []int, target int) int {
    lo := 0
    hi := len(a) - 1
    
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        if a[mid] == target {
            return mid
        }
        if a[mid] < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    
    return -1
}