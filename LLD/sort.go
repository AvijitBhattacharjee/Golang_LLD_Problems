package main

import (
	"fmt"
	"sync"
)


func main() {

    var arr = []int{1,4,8,2,6,3,9,0,7,5}
    var arr2 []int
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        quickSort(arr, 0, len(arr)-1)
        
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        arr2 = mergeSort(arr)
        
    }()

    wg.Wait()
    fmt.Println(arr)
    fmt.Println(arr2)
}

func quickSort(nums []int, low, high int) {

    if low >= high {return}

    p := partition(nums, low, high)

    quickSort(nums, low, p-1)
    quickSort(nums, p+1, high)
}


func partition(nums []int, low, high int) int {

    pivot := nums[high]
    var i,j = low, low

    for j=low;j<high;j++ {

        if nums[j] < pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    nums[i], nums[high] = nums[high], nums[i]
    return i
}




func mergeSort(nums []int) []int {

    if len(nums) <= 1 {return nums}

    mid := len(nums)/2
    left, right := mergeSort(nums[:mid]), mergeSort(nums[mid:])

    return merge(left, right)
}

func merge(nums1, nums2 []int) []int {

    result := make([]int, 0, len(nums1)+len(nums2))
    i,j := 0,0

    for i < len(nums1) && j <  len(nums2) {

        if nums1[i] < nums2[j] {
            result = append(result, nums1[i]); i++
        } else {
            result = append(result, nums2[j]); j++
        }
    }


    if i<len(nums1) {
        result = append(result, nums1[i:]...); i++
    }

    if j<len(nums2) {
        result = append(result, nums2[j:]...); j++
    }

    return result
}


