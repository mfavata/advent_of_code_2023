package main

import (
  "fmt"
  "strings"
  "os"
  "regexp"
)

func cardScore(win []string, have []string, part string) int {
  var score, count int
  for x := 0; x < len(win); x++ {
    for y := 0; y < len(have); y++ {
      if win[x] == have[y] {
        count++
        if score == 0 {
          score = 1 
          break
        }
        score = score+score
      }
    }
  }
  if part == "1" {
    return score
  }
  return count
}

func partTwo(card_wins []int) int {
  cards := [202]int{}
  for i := 0; i < len(card_wins); i++ {
    cards[i]++
    for y := 0; y < cards[i]; y++ {
      for x := 1; x <= card_wins[i]; x++ {
        cards[i+x]++
      }
    }
  }
  var total int
  for j := 0; j < len(cards); j++ {
    total += cards[j]
  }
  return total
}

func main() {
  readFile, _ := os.ReadFile("day4.txt")

  input := strings.Split(string(readFile), "\n")
  var part_one_total int  
  var card_wins []int

  for i := 0; i < len(input)-1; i++ {
    digits := regexp.MustCompile(`[\d]+`)
    card := strings.Split(input[i], "|")
    winning_numbers := digits.FindAllString(card[0][9:], -1)
    numbers_you_have := digits.FindAllString(card[1], -1)
    part_one_total += cardScore(winning_numbers, numbers_you_have, "1")
    card_wins = append(card_wins, cardScore(winning_numbers, numbers_you_have, "2"))
  }

  fmt.Println("Part One Solution:", part_one_total)
  fmt.Println("Part Two Solution:", partTwo(card_wins))
}
