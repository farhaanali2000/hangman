package main

import (
	"fmt"
	"io/ioutil"
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

	fmt.Println(renderHangman(5))
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
	fmt.Println(getWordGuessingProgress(targetWord, guessedWord))
}

func getWordGuessingProgress(targetWord string, guessedWord map[rune]bool) string {

	var result = ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedWord[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}

	return result
}

func renderHangman(hangman int) string {
	data, err := ioutil.ReadFile("states/hangman6")

	if err != nil {
		panic(err.Error())
	}

	return string(data)
}

// func userInput(guessed map[rune]bool) {
// 	var letter string
// 	fmt.Scanln(&letter)

// 	fmt.Println(letter)
// 	// fmt.Println(c, string(c), strconv.QuoteRune(c))
// 	guessed[letter] = true
// }
