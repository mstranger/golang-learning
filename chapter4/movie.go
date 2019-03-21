package main

import (
  "encoding/json"
  "fmt"
  "log"
)

type Movie struct {
  Title  string
  Year   int  `json:"released"`
  Color  bool `json:"color,omitempty"`
  Actors []string
}

var movies = []Movie{
  {Title: "Casablanca", Year: 1942, Color: false,
    Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
  {Title: "Cool Hand Luke", Year: 1967, Color: true,
    Actors: []string{"Paul Newman"}},
  {Title: "Bullitt", Year: 1968, Color: true,
    Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

var titles []struct {
  Title string
  Year  int `json:"released"`
}

func main() {
  // marshaling
  // data, err := json.Marshal(movies)
  data, err := json.MarshalIndent(movies, "", " ")
  if err != nil {
    log.Fatalf("Error marshaling JSON: %s", err)
  }

  fmt.Printf("%s\n", data)
  // unmarshaling
  if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("Error unmarshaling JSON: %s", err)
  }

  fmt.Println(titles)
}
