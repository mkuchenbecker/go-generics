package main

import (
	"fmt"
	"testing"
)


func TestMap(t *testing.T)	{
	Map[int, string]( SliceIterator[int]([]int{1, 2, 3, 4, 5}), func(i int) string {
		return fmt.Sprintf("%d", i)
	})

	Map[int, string](SliceIterator[int]([]int{}), func(i int) string {
		return fmt.Sprintf("%d", i)
	})
	// Map[string,int](toStrings, func(s string) int {
	// 	// convert string s to int
	// 	intValue, err := strconv.Atoi(s)
	// 	assert.NoError(t, err)
	// 	return intValue
	// })
}