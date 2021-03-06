// p. 114
package main

import "fmt"

func reverse(s []int) {
  for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}

func main() {

  //  note: type = [6]int  - array
  arr := [...]int{3,4,5,6,7,8}

  // note: type = []int  - slice
  reverse(arr[:])

  fmt.Println(arr)
}