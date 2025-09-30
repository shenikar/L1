package main

import (
	"fmt"
	"strings"
)

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{3, 4, 5}

	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1, nums2 []int) string {
	set := make(map[int]struct{})
	var res []int

	for _, num := range nums1 {
		set[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, exist := set[num]; exist {
			res = append(res, num)
		}
	}

	// форматируем результат в строку с запятыми
	if len(res) == 0 {
		return "[]"
	}
	strs := make([]string, len(res))
	for i, v := range res {
		strs[i] = fmt.Sprint(v)
	}
	return "[" + strings.Join(strs, ",") + "]"
}
