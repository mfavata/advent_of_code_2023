package main

import (
  "fmt"
  "os"
  "strings"
  "sort"
  "strconv"
)

const cards = "23456789ABCDE"

func str_to_int(s string) int {
  a, _ := strconv.Atoi(s)
  return a
}

func transpose_cards(hand string) string {
  new_hand := hand
  new_hand = strings.Replace(new_hand, "A", "E", -1)
  new_hand = strings.Replace(new_hand, "K", "D", -1)
  new_hand = strings.Replace(new_hand, "Q", "C", -1)
  new_hand = strings.Replace(new_hand, "J", "B", -1)
  new_hand = strings.Replace(new_hand, "T", "A", -1)
  return new_hand
}

func identify_hand(hand string) int {
  if strings.Contains(hand, "5") {
    return 6 //"five_of_a_kind"
  } else if strings.Contains(hand, "4") {
    return 5 //"four_of_a_kind" 
  } else if strings.Contains(hand, "3") && strings.Contains(hand, "2") {
    return 4 //"full_house"
  } else if strings.Contains(hand, "3") {
    return 3 //"three_of_a_kind"
  } else if strings.Contains(hand, "2") && strings.Count(hand, "2") > 2 {
    return 2 //"two_pair"
  } else if strings.Contains(hand, "2") {
    return 1 //"one_pair"
  }
  return 0 //"high_card"
}

func uniqueCount(str string) string {
  var unique_count string
  for _, c := range str {
    unique_count = unique_count+string(strings.Count(str, string(c))+48)  
  }
  return unique_count
}


func main() {
  readFile, _ := os.ReadFile("day7.txt")
  input := strings.Split(string(readFile), "\n")
  hands := []string{}

  for x := 0; x < len(input)-1; x++ {
    hand := strings.Split(input[x], " ")
    hand[0] = transpose_cards(hand[0])
    updated_hand := string(identify_hand(uniqueCount(hand[0]))+48) + hand[0] + " " + hand[1]
    hands = append(hands, updated_hand)
   
  }
  sort.Slice(hands, func(i, j int) bool {
      return hands[i] < hands[j]
   })
 
  var part_one_total int

  for i:=0; i < len(hands); i++ {
    current := strings.Split(hands[i], " ")
    part_one_total += (i+1)*str_to_int(current[1])
  }

  fmt.Println("Part One Solution:", part_one_total)
}
