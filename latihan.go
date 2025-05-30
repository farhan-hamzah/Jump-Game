package main

import "fmt"

func canJump(nums []int) bool {
    maxReach := 0
    for i, num := range nums {
        if i > maxReach {
            return false
        }
        if i+num > maxReach {
            maxReach = i + num
        }
    }
    return true
}

func main() {
    // Contoh penggunaan
    nums1 := []int{2, 3, 1, 1, 4}
    fmt.Println("Can jump (example 1):", canJump(nums1)) // true

    nums2 := []int{3, 2, 1, 0, 4}
    fmt.Println("Can jump (example 2):", canJump(nums2)) // false
}
