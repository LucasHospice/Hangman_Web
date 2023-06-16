package main

import (
	"Hangman/Hangpack"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HangmanData struct {
	Word     string
	ToFind   string
	Attempts int
	Valid    map[string]bool
	Level    string
	Nom      string
}

func main() {
	game := HangmanData{}
	tmpl, err := template.ParseFiles("./templates/username.html")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, game)
	})

	http.HandleFunc("/username", func(w http.ResponseWriter, r *http.Request) {
		game.Nom = r.FormValue("nom")
		fmt.Println("Nom:", game.Nom)
		http.Redirect(w, r, "/level", http.StatusFound)
	})

	http.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/level.html")
		tmpl.Execute(w, game)
	})

	http.HandleFunc("/word", func(w http.ResponseWriter, r *http.Request) {
		value := r.FormValue("level")
		fmt.Println("Difficulté:", value)
		if value == "facile" {
			game.Level = "words.txt"
		} else if value == "normal" {
			game.Level = "words2.txt"
		} else if value == "difficile" {
			game.Level = "words3.txt"
		}
		game.ToFind = Hangpack.Mot(game.Level)
		game.Word = Hangpack.Affichage(game.ToFind)
		game.Attempts = 10
		game.Valid = map[string]bool{}
		fmt.Println("Mot:", game.ToFind)
		http.Redirect(w, r, "/game", http.StatusFound)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/layout.html")
		tmpl.Execute(w, game)
	})

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		Lettre := r.FormValue("lettre/mot")
		var b bool = false
		c := []rune(game.ToFind)
		d := []rune(game.Word)
		if Lettre == "2022" {
			http.Redirect(w, r, "/bonnefetes", http.StatusFound)
		} else if len(Lettre) >= 2 {
			if Lettre == game.ToFind {
				http.Redirect(w, r, "/win", http.StatusFound)
			} else {
				game.Attempts = game.Attempts - 2
				if game.Attempts <= 0 {
					http.Redirect(w, r, "/lost", http.StatusFound)
				}
			}
		} else if len(Lettre) == 0 {

		} else {
			for i := 0; i <= len(game.ToFind)-1; i++ {
				if Lettre[0] == game.ToFind[i] {
					b = true
				}
			}
			if !Hangpack.LettreUtilise(Lettre, game.Valid) {

			} else if b == true {
				for j := 0; j < len(game.ToFind); j++ {
					if c[j] == rune(Lettre[0]) {
						d[j] = c[j]
						game.Word = string(d)
					}
				}
				if game.Word == game.ToFind {
					http.Redirect(w, r, "/win", http.StatusFound)
				}
			} else {
				game.Attempts = game.Attempts - 1
				if game.Attempts <= 0 {
					http.Redirect(w, r, "/lost", http.StatusFound)
				}
			}
		}
		http.Redirect(w, r, "/game", http.StatusFound)
	})

	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/win.html")
		tmpl.Execute(w, game)
	})

	http.HandleFunc("/lost", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/lost.html")
		tmpl.Execute(w, game)
	})

	http.HandleFunc("/bonnefetes", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/bonnefete.html")
		tmpl.Execute(w, game)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Connecté au port 8080...")
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		log.Fatal(err)
	}
}
