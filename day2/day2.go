package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func string_to_int(a string) int {
  b, _ := strconv.Atoi(a)
  return b
}

func part_one(a []string, b int) int {
  for i := 2; i < len(a); i=i+2 {
    if strings.Contains(a[i+1], "red") && string_to_int(a[i]) > 12 {
      return 0
      break
    } 
    if strings.Contains(a[i+1], "green") && string_to_int(a[i]) > 13 {
      return 0
      break
    }
    if strings.Contains(a[i+1], "blue") && string_to_int(a[i]) > 14 {
      return 0
      break
    }
  }
  return b+1
}

func part_two(a []string) int {
  var r, g, b = 0, 0, 0
  for i := 2; i < len(a); i=i+2 {
    if strings.Contains(a[i+1], "red") && string_to_int(a[i]) > r {
      r = string_to_int(a[i])
    }
    if strings.Contains(a[i+1], "green") && string_to_int(a[i]) > g {
      g = string_to_int(a[i])
    }
    if strings.Contains(a[i+1], "blue") && string_to_int(a[i]) > b {
      b = string_to_int(a[i])
    }
  }
  return r*g*b  
}

func main() {
  readFile, err := os.ReadFile("day2.txt")
  if err != nil {
    fmt.Println(err)
  }

  input := strings.Split(string(readFile), "\n")
  var part_one_total, part_two_total = 0, 0

  for i := 0; i < len(input)-1; i++ {
    part_one_total += part_one(strings.Split(input[i], " "), i)
    part_two_total += part_two(strings.Split(input[i], " "))
  }

  fmt.Println("Part One:", part_one_total, "Part Two:", part_two_total)
}
