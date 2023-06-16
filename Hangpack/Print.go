package Hangpack

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func LettreX(Lettre *HangmanData) { // Affiche le hangman
	f, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var ligne int
	for scanner.Scan() {
		if ligne < Lettre.Position*8 && ligne >= (Lettre.Position-1)*8 {
			fmt.Println(scanner.Text())
		}
		ligne += 1
	}
}

func Art(Ascii *HangmanData) { // ascii art
	var tableau [9][]string
	tabword := []rune(Ascii.Word)
	for d := 0; d < len(Ascii.Word); d++ {
		a := 0
		f, err := os.Open(Ascii.Police)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		var line int
		for scanner.Scan() {
			if a == 9 {
				a = 0
			}
			if line < ((9*(int(tabword[d])-32))+9)+1 && line >= (9*(int(tabword[d])-32))+1 {
				tableau[a] = append(tableau[a], scanner.Text())
			}
			line++
			a++
		}
	}
	for c := 0; c < 9; c++ {
		fmt.Println(strings.Join(tableau[c], " "))
	}
}

func Printall(Print *HangmanData) {
	if Print.Police != "" {
		Art(Print)
	} else {
		fmt.Println(Print.Word)
	}
}
