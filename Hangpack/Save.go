package Hangpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func LettreUtilise(a string, valid map[string]bool) bool { //stock les lettres (advence future)
	var lettreutilise bool = true
	b := valid[string(a[0])]
	if b {
		lettreutilise = false
	} else {
		valid[string(a[0])] = true
	}
	return lettreutilise
}

func Save(Nom string, ToFind string, Attempts int, Level string) { //stock la save
	var save []string
	// Tentatives := string(Attempts)
	if Level == "words.txt" {
		Level = "Facile"
	} else if Level == "words2.txt" {
		Level = "Normal"
	} else if Level == "words3.txt" {
		Level = "Difficile"
	}
	save = append(save, Nom, ToFind, strconv.Itoa(Attempts), Level)
	fmt.Println(Level)
	// file, _ := os.Create("save.txt")
	// convert, _ := json.Marshal(save)
	// file.WriteString(string(convert))
	if Attempts > 5 {
		file, err := os.OpenFile("/templates/save.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		str := strings.Join(save, ", ")
		if _, err := file.WriteString(str); err != nil {
			log.Fatal(err)
		}
		file.WriteString("\n")
	}
}

func ReadSave(scoreboard struct{}) { // lit la save
	file, _ := ioutil.ReadFile("/templates/save.txt")
	json.Unmarshal(file, &scoreboard)

	// fmt.Println("Bonne chance, vous avez", Read.Attempts, "tentatives.")
	// Printall(Read)
}
