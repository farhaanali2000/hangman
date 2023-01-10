package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)

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
	hangmanState := 0

	for !isWordGuessed(randomWord, guessedLetters) && !isHangManComplete(hangmanState) {
		printGameState(randomWord, guessedLetters, hangmanState)
		input := readInput()

		if len(input) != 1 {
			fmt.Println("invalid input. please use letters only...")
			continue
		}

		letter := rune(input[0])
		if isCorrectGuess(randomWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}
	fmt.Println("Game Over!")
	if isWordGuessed(randomWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangManComplete(hangmanState) {
		fmt.Println("You lose!")
	} else {
		panic("invalid state. Game is over there is no winner")
	}

	// userInput(guessedLetters)

}

func initializeGuessedWord(randomWord string) map[rune]bool {

	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(randomWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(randomWord[len(randomWord)-1]))] = true

	return guessedLetters
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isHangManComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func printGameState(targetWord string, guessedWord map[rune]bool, hangmanState int) {

	//football
	fmt.Println(getWordGuessingProgress(targetWord, guessedWord))
	fmt.Println(renderHangman(hangmanState))
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

func renderHangman(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))

	if err != nil {
		panic(err.Error())
	}

	return string(data)
}

func readInput() string {
	fmt.Printf(">")
	r := inputReader
	input, err := r.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}

// func userInput(guessed map[rune]bool) {
// 	var letter string
// 	fmt.Scanln(&letter)

// 	fmt.Println(letter)
// 	// fmt.Println(c, string(c), strconv.QuoteRune(c))
// 	guessed[letter] = true
// }
