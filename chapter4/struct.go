package main

import (
  "fmt"
  "time"
)

type Employee struct {
  ID            int
  Name, Address string
  DoB           time.Time
  Position      string
  Salary        int
  ManagerID     int
}

func main() {
  var dilbert Employee

  dilbert.Salary -= 5000
  position := &dilbert.Position
  *position = "Senior"
  fmt.Printf("%#v\n", dilbert)
  fmt.Printf("%v\n", dilbert)
}
