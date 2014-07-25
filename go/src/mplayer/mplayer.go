package main

import (
	"bufio"
	"fmt"
	"mplayer/lib"
	"mplayer/mp"
	"os"
	"strconv"
	"strings"
)

var manager *lib.MusicManager
var id int = 1

//var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < manager.Len(); i++ {
			e, _ := manager.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			manager.Add(&lib.Music{strconv.Itoa(id),
				tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])

	}

}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := manager.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	manager = lib.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized commnand:", tokens[0])
		}
	}
}
