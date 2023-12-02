package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID     int64
	maxMap map[string]int64 // color -> max number
}

func NewGame(s string) *Game {
	g := &Game{
		maxMap: make(map[string]int64),
	}

	lineSplit := strings.Split(s, ": ")
	game := lineSplit[0]
	subsets := lineSplit[1]

	var gameId int64
	gameId, err := strconv.ParseInt(strings.Split(game, " ")[1], 10, 64)
	if err != nil {
		log.Fatalln("Failed to parse input at line :", s)
	}
	g.ID = gameId

	for _, subset := range strings.Split(subsets, "; ") {
		for _, record := range strings.Split(subset, ", ") {
			recSplit := strings.Split(record, " ") // record is in form "<count> <color>"
			color := recSplit[1]
			count, err := strconv.ParseInt(recSplit[0], 10, 64)
			if err != nil {
				log.Fatalln("Failed to parse input at line :", s)
			}
			g.maxMap[color] = max(g.maxMap[color], count)
		}
	}
	return g
}

func (g *Game) power() int64 {
	var power int64 = 1

	for _, count := range g.maxMap {
		power *= count
	}

	return power
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sumOfPowers int64 = 0

	for scanner.Scan() {
		game := NewGame(scanner.Text())
		sumOfPowers += game.power()
	}

	fmt.Println(sumOfPowers)
}
