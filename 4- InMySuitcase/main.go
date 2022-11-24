package main

import (
	"fmt"
)

func main() {
	sentence := "Dans ma valise, il y a …"
	word := []string{}
	temp := ""

	for temp != "fin" {
		fmt.Printf("\nEntrez un mot : ")
		fmt.Scan(&temp)

		if temp != "fin" {
			word = append(word, temp)
			fmt.Printf(sentence)
			showWords(word)
		}
	}
}

func showWords(word []string) {
	for i := 0; i < len(word); i++ {

		fmt.Printf(" %s", word[i])

		if i < len(word)-2 {
			fmt.Printf(",")
		}

		if i == len(word)-2 {
			fmt.Printf(" et")
		}

		if i == len(word)-1 {
			fmt.Printf(".")
		}

	}
}
