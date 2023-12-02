package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

func part_one(a string) int {
  var temp string
  
  for i := 0; i < len(a); i++ {
    if a[i] >= 48 && a[i] <= 57 {
      temp += string(a[i])
    }
  }

  retval, _ := strconv.Atoi(string(temp[0]) + string(temp[len(temp)-1]))
  return retval
}

func part_two(a string) string {
  list := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
  b := a
  for i := 1; i <= len(b); i++ {
    for j := 0; j < len(list); j++ {
      if strings.Contains(b[:i], list[j]) {
        b = strings.Replace(b[:i], list[j], string(byte(j)+48)+b[i-1:], -1)
        i = i - len(list[j])
      }
    }
  }

  return b
}
 

func main() {

  readFile, err := os.ReadFile("day1.txt")
  if err != nil {
    fmt.Println(err)
  }

  input := strings.Split(string(readFile), "\n")
  
  part_one_total := 0
  part_two_total := 0

  for i := 0; i < len(input) - 1; i++ {
    part_one_total += part_one(input[i])
    part_two_total += part_one(part_two(input[i]))

  }

  fmt.Println("Part 1:", part_one_total)
  fmt.Println("Part 2:", part_two_total)
}
