package main

import "fmt"

func main() {
  arr1 := [...]int{3,4,5}
  arr2 := []int{3,4,5}
  arr3 := []int{3,4,5}

  fmt.Printf("%T\n%[1]v\n%T\n[2]%v\n", arr1, arr2)
  // [3]int --- array
  // [2 3 4]
  // []int  --- slice
  // [2 3 4]

  fmt.Printf("%t\n", arr1 == arr2)
  // error -> mismatched types

  fmt.Printf("%t\n", arr2 == arr3)
  // error -> slices cannot be compared!
}