package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/pterm/pterm"

	"github.com/MarkusZoppelt/go-lazydice/wordlist"
)

func main() {
	numWords := getWordNumFromDialog()
	var words []string

	for i := 0; i < numWords; i++ {
		words = append(words, generateWord())
	}
	fmt.Println(words)
}

func getWordNumFromDialog() int {
	choices := []string{
		"4 words (51 bits of entropy)",
		"5 words (64 bits of entropy)",
		"6 words (77 bits of entropy)",
	}

	selection, _ := pterm.DefaultInteractiveSelect.WithOptions(choices).WithDefaultText("How strong do you want your passphrase to be?").Show()
	num, _ := strconv.Atoi(strings.Split(selection, " ")[0])

	return num
}

func generateWord() string {
	number := 0

	for i := 0; i < 5; i++ {
		diceRoll, _ := rand.Int(rand.Reader, big.NewInt(6))
		number *= 10
		number += int(diceRoll.Int64()) + 1
	}

	return wordlist.GetWord(number)
}
