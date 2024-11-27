package main

import "fmt"

func countWord(word string) {
  fmt.Println(len(word))

  for i := 0; i < len(word); i++ {
    fmt.Printf("Character at index %d: %c\n", i, word[i])
  }
}

func main() {
  randomWord := "Gabriel"  

  countWord(randomWord)

}
