package main

import "fmt"

func main() {

//Part One:
//time := []int{56,97,78,75}
//distance := []int{546,1927,1131,1139}

//Part Two:
  time := []int{56977875}
  distance := []int{546192711311139}

  solution := 1

  for i := 0; i < len(time); i++ {
    new_record := 0
    for hold := 1; hold < time[i]; hold++ {
      if (time[i]-hold)*hold > distance[i] {
        new_record++
      }
    }
    fmt.Println("New Record Count:", new_record)
    solution = solution*new_record
  }
  fmt.Println("Solution:", solution)
}
