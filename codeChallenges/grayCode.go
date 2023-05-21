package codeChallenges

import (
	"errors"
	"fmt"
)

// validInputN Check value of n
func validInputN(n int) error {
	if n < 1 || n > 16 {
		return errors.New("error: The value of n is outside the valid range [1, 16]")
	}
	return nil
}

// GrayCode ...
func GrayCode(n int) ([]int, error) {
	if err := validInputN(n); err != nil {
		return nil, err
	}
	count := 1 << uint(n)
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = i ^ (i >> 1)
	}
	return result, nil
}

func RunGrayCode() {
	// GrayCode
	n := 9
	result, err := GrayCode(n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
