package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"pokemon",
	"cricket",
	"football",
	"messi",
	"charizard",
	"mclaren",
	"verstappen",
	"tomlinson",
	"golang",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randomWord := getRandomWord()

	guessedLetters := initializeGuessedWord(randomWord)
	printGameState(randomWord, guessedLetters)
	// userInput(guessedLetters)
}

func initializeGuessedWord(randomWord string) map[rune]bool {

	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(randomWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(randomWord[len(randomWord)-1]))] = true

	return guessedLetters

}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func printGameState(targetWord string, guessedWord map[rune]bool) {

	//football
	for _, ch := range targetWord {
		if ch == ' ' {
			fmt.Printf(" ")
		} else if guessedWord[unicode.ToLower(ch)] == true {
			fmt.Printf("%c", ch)
		} else {
			fmt.Printf("_")
		}
		print(" ")
	}
}

// func userInput(guessed map[rune]bool) {
// 	var letter string
// 	fmt.Scanln(&letter)

// 	fmt.Println(letter)
// 	// fmt.Println(c, string(c), strconv.QuoteRune(c))
// 	guessed[letter] = true
// }
