// p. 141
package main

// TODO: task 4.10 - 4.13

// TODO: main and github packages

// TODO: chapter 4.6, task 4.14

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "net/url"
  "os"
  "strings"
  "time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
  TotalCount int `json:"total_count"`
  Items      []*Issue
}

type Issue struct {
  Number    int
  HTMLURL   string
  Title     string
  State     string
  User      *User
  CreatedAt time.Time `json:"created_at"`
  Body      string
}

type User struct {
  Login string
  HTML  string `json:"html_url"`
}

// SearchIssues request GitHub
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
  q := url.QueryEscape(strings.Join(terms, " "))
  resp, err := http.Get(IssueURL + "?q=" + q)
  if err != nil {
    return nil, err
  }

  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return nil, fmt.Errorf("Error request: %s", resp.Status)
  }

  var result IssuesSearchResult
  if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    resp.Body.Close()
    return nil, err
  }
  resp.Body.Close()
  return &result, nil
}

func main() {
  result, err := SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%d themes:\n", result.TotalCount)
  for _, item := range result.Items {
    fmt.Printf("#%-5d %9.9s %.55s\n",
      item.Number, item.User.Login, item.Title)
  }
}
