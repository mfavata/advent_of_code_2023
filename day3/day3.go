package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "regexp"
)

func checkLimit(x int) int {
  if x > 139 {
    return 139
  } else if x < 0{
    return 0
  }
  return x
}

func checkLines(a, b, c int, sch []string) int {
  symbols := regexp.MustCompile(`[^.^0-9]`)
  start := checkLimit(b-1)
  end := checkLimit(c+1)

  for i := checkLimit(a-1); i <= checkLimit(a+1); i++ {
    if symbols.MatchString(sch[i][start:end]) {
      return string_to_int(sch[a][b:c])
      break
    }
  }
  return 0
}

func partTwo(a, b, c int, sch []string) (string, int) {
  symbols := regexp.MustCompile(`[*]`)
  start := checkLimit(b-1)
  end := checkLimit(c+1)

  for i := checkLimit(a-1); i <= checkLimit(a+1); i++ {
    if symbols.MatchString(sch[i][start:end]) {
      first, _ := returnUpperAndLower(symbols.FindStringIndex(sch[i][start:end]))
      return strconv.Itoa(i)+":"+strconv.Itoa(start+first), string_to_int(sch[a][b:c])
    }
  }
  return string(0), 0
}

func gearRatios(gear [1000]string, num [1000]int) int {
  var running_total int
  for i := 0; i < len(gear); i++ {
    for x:= i+1; x < len(gear); x++ {
      if (gear[i] == gear[x]) && num[i] != 0 {
        running_total += num[i]*num[x]
      }
    }
  }

  return running_total
}

func returnUpperAndLower(a []int) (int, int) {
  return a[0], a[1]
}

func string_to_int(a string) int {
  b, _ := strconv.Atoi(a)
  return b
}

func main() {
  readFile, err := os.ReadFile("day3.txt")
  if err != nil {
    fmt.Println(err)
  }
  
  schematic := strings.Split(string(readFile), "\n")
  
  var part_one_total, tracking int // part_two_total int
  gear_positions := [1000]string{}
  numbers := [1000]int{}

  for i := 0; i < len(schematic); i++ {
    re := regexp.MustCompile("[0-9]+")
    num_index := re.FindAllStringIndex(schematic[i], -1)
    for x := 0; x < len(num_index); x++ {
      a, b := returnUpperAndLower(num_index[x])
      part_one_total += checkLines(i, a, b, schematic)
      gear_positions[tracking], numbers[tracking] = partTwo(i, a, b, schematic)
      if numbers[tracking] != 0 {
        tracking++
      }
    }
  }
  fmt.Println("Part One Solution:", part_one_total)
  fmt.Println("Part Two Solution:", gearRatios(gear_positions, numbers))

}
