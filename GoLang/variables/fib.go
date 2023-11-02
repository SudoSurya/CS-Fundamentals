package main


func init() {
    println("init")

}

// fibannaci using for loop
func fib(n int) int {
    seq := []int{0, 1}

    for i:=2 ; i<=n ; i++ {
        seq = append(seq, seq[i-1] + seq[i-2])
    }
    return seq[n]
}
// The frequency of an element is the number of times it occurs in an array.

// You are given an integer array nums and an integer k. In one operation, you can choose an index of nums and increment the element at that index by 1.

// Return the maximum possible frequency of an element after performing at most k operations.


func check(nums []int) bool {

    for i:=1 ; i<len(nums) ; i++ {
        if nums[i] < nums[i-1] {
            return false
        }
    }
    return true

}
