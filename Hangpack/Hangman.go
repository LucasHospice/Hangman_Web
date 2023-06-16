package Hangpack

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

type HangmanData struct {
	Position int
	Lettre   string
	Word     string // Word composed of '_', ex: H_ll_
	ToFind   string // Final word chosen by the program at the beginning. It is the word to find
	Attempts int    // Number of attempts left
	valid    map[string]bool
	Police   string
}

func Mot(Level string) string { // Choisi un mot aléatoirement
	file := ""
	a := 0
	f, err := os.Open(Level)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a = a + 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	b := rand.Intn(a)
	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	for i := 0; i <= b; i++ {
		scanner.Scan()
		if i == b {
			file = scanner.Text()
		}
	}
	return file
}

func Affichage(ToFind string) string { // Affiche le mot transformé en _ avec des lettres placées aléatoirement
	var d string
	a := []rune(ToFind)
	b := []rune(d)
	n := len(ToFind)/2 - 1
	for j := 0; j < len(ToFind); j++ {
		b = append(b, '_')
	}
	for i := 0; i < n; i++ {
		c := rand.Intn(len(ToFind))
		b[c] = a[c]
	}
	word := string(b)
	return word
}

// func Game(ToFind string, Word string, Attempts int, Position int) { //jeu + advence future
// 	valid := map[string]bool{}
// 	for {
// 		var b bool = false
// 		c := []rune(ToFind)
// 		d := []rune(Word)
// 		fmt.Print("Choisi une lettre : ")
// 		reader := bufio.NewReader(os.Stdin)
// 		a, _ := reader.ReadString('\n')
// 		// if a == "stop\n" {
// 		// 	fmt.Println("Les données sont sauvegardées...")
// 		// 	Save(Jeu)
// 		// 	break
// 		// }
// 		if a == "\n" {
// 			fmt.Println("Merci de rentrer une lettre.")
// 			continue
// 		}
// 		if len(a) > 2 {
// 			if a == ToFind+"\n" {
// 				fmt.Println("Bravo ! Vous avez trouvé le mot qui était", string(c), "!")
// 				break
// 			} else {
// 				Attempts = Attempts - 2
// 				fmt.Println("Le mot rentré ne correspond pas au mot à trouver, il vous reste", Attempts, "tentatives.")
// 				Jeu.Position = Position + 2
// 				LettreX(Jeu)
// 				if Jeu.Attempts == 0 {
// 					fmt.Println("Vous avez utilisé toutes vos tentatives, vous avez perdu !")
// 					break
// 				}
// 				continue
// 			}
// 		}
// 		for i := 0; i <= len(ToFind)-1; i++ {
// 			if a[0] == ToFind[i] {
// 				b = true
// 				Jeu.Lettre = string(a)
// 			}
// 		}
// 		if !LettreUtilise(a, valid) {
// 			continue
// 		}
// 		if b == true {
// 			for j := 0; j < len(Jeu.ToFind); j++ {
// 				if c[j] == rune(a[0]) {
// 					d[j] = c[j]
// 					Jeu.Word = string(d)
// 				}
// 			}
// 			Printall(Jeu)
// 			if Jeu.Word == Jeu.ToFind {
// 				fmt.Println("Bravo ! Vous avez trouvé le mot qui était", string(c), "!")
// 				break
// 			}
// 		} else {
// 			Jeu.Attempts = Jeu.Attempts - 1
// 			fmt.Println("La lettre", string(a[0]), "n'est pas présente dans le mot. Il vous reste", Jeu.Attempts, "tentatives.")
// 			Jeu.Position = Jeu.Position + 1
// 			LettreX(Jeu)
// 			if Jeu.Attempts == 0 {
// 				fmt.Println("Vous avez utilisé toutes vos tentatives, vous avez perdu !")
// 				break
// 			}
// 		}
// 	}
// }
