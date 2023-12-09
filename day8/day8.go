package main

import (
  "fmt"
  "os"
  "strings"
  "regexp"
)

func mapInput(s string) map[string][]string {
  regex := regexp.MustCompile("[A-Z]+") 
  lines := strings.Split(s, "\n")
  result := make(map[string][]string)

  for i := 0; i < len(lines)-1; i++ {
    temp := regex.FindAllString(lines[i], -1)
    key := temp[0]
    values := []string{temp[1], temp[2]}
    result[key] = values
  }
  return result
}

func partOne(a map[string][]string, step string) int {
  var count, i = 0, 0
  next := "AAA"

  for next != "ZZZ" {
    switch string(step[i]) {
    case "L":
      next = a[next][0]
    case "R":
      next = a[next][1]
    }

    if i == len(step)-1 {
      i = 0
    } else {
      i++
    }

    count++
  }
  return count
}

func main() {
  readFile, _ := os.ReadFile("day8.txt")
  input := strings.Split(string(readFile), "\n\n")
  
  fmt.Println("Part One Solution:", partOne(mapInput(input[1]), input[0]))
}
