package codeChallenges

import (
	"errors"
	"fmt"
)

func ValidInputSubArrayLength(arr1, arr2 []int) error {
	// Validate array lengths
	if len(arr1) == 0 || len(arr2) == 0 || len(arr1) > 1000 || len(arr2) > 1000 {
		return errors.New("Invalid array length")
	}

	// Validate array values
	for _, value := range arr1 {
		if value < 0 || value > 100 {
			return errors.New("Invalid value in arr1")
		}
	}
	for _, value := range arr2 {
		if value < 0 || value > 100 {
			return errors.New("Invalid value in arr2")
		}
	}
	return nil
}

func maxCommonSubarrayLength(arr1 []int, arr2 []int) (int, error) {
	if err := ValidInputSubArrayLength(arr1, arr2); err != nil {
		return 0, err
	}

	// Initialize binary search boundaries
	left := 0
	right := len(arr1)
	maxLength := 0

	// Perform binary search on the minimum length of the common subarray
	for left <= right {
		mid := left + (right-left)/2
		if isCommonSubarray(arr1, arr2, mid) {
			maxLength = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return maxLength, nil
}

// Check if there exists a common subarray of length k
func isCommonSubarray(arr1 []int, arr2 []int, k int) bool {
	// Generate a set of all subarrays of length k in arr1 using hashing
	subarrays := make(map[int]bool)
	for i := 0; i <= len(arr1)-k; i++ {
		subarrayHash := hash(arr1[i : i+k])
		subarrays[subarrayHash] = true
	}

	// Check if any subarray of length k in arr2 matches a subarray in arr1
	for i := 0; i <= len(arr2)-k; i++ {
		subarrayHash := hash(arr2[i : i+k])
		if subarrays[subarrayHash] {
			return true
		}
	}

	return false
}

// Generate a unique hash value for an array
func hash(arrs []int) int {
	hashValue := 0
	for _, value := range arrs {
		hashValue += value
	}
	return hashValue
}

func RunMaximumLengthofRepeatedSubarray() {
	arr1 := []int{1, 2, 3, 2, 1}
	arr2 := []int{3, 2, 1, 4, 7}
	result, err := maxCommonSubarrayLength(arr1, arr2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result) // Output: 3

	arr1 = []int{0, 0, 0, 0, 0}
	arr2 = []int{0, 0, 0, 0, 0}
	result, err = maxCommonSubarrayLength(arr1, arr2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result) // Output: 5

	// Invalid input example - empty array
	// arr1 = []int{}
	// arr2 = []int{1, 2, 3}
	// result, err = maxCommonSubarrayLength(arr1, arr2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result) // Output: 0

	// Invalid input example - array length exceeding constraint
	// arr1 = make([]int, 1001)
	// arr2 = make([]int, 1000)
	// result, err = maxCommonSubarrayLength(arr1, arr2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result) // Output: 0

	// Invalid input example - array value exceeding constraint
	arr1 = []int{1, 2, 3, 101}
	arr2 = []int{4, 5, 6, 7}
	result, err = maxCommonSubarrayLength(arr1, arr2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result) // Output: 0
}
