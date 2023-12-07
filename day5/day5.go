package main

import (
  "fmt"
  "strconv"
  "strings"
  "os"
)

func str_to_int(in string) int {
  out, _ := strconv.Atoi(in)
  return out
}

func main() {
  readFile, _ := os.ReadFile("day5.txt")
  input := strings.Split(string(readFile), "\n\n")

//Only solution for Part One
  seed := []int{1132132257, 323430997, 2043754183, 4501055, 2539071613, 1059028389, 1695770806, 60470169, 2220296232, 251415938, 1673679740, 6063698, 962820135, 133182317, 262615889, 327780505, 3602765034, 194858721, 2147281339, 37466509} 

  for m := 1; m < len(input); m++ {
    almanac := strings.Split(input[m], "\n")
    for x := 0; x < len(seed); x++ {
      for i := 1; i < len(almanac); i++ {
        map_range := strings.Split(almanac[i], " ")
        if map_range[0] == "" {
          break
        }
        if seed[x] >= str_to_int(map_range[1]) && seed[x] < str_to_int(map_range[1])+str_to_int(map_range[2]) {
          seed[x] = seed[x] + (str_to_int(map_range[0])-str_to_int(map_range[1]))
          break
        }
      }
    }
}
  lowest := seed[0]
  for i := 0; i < len(seed); i++ {
    if seed[i] < lowest {
      lowest = seed[i]
    }
  }

  fmt.Println(lowest)
}
