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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxExpected := Game{
		-1,
		map[string]int64{
			"red":   12,
			"green": 13,
			"blue":  14,
		},
	}

	var sumOfIds int64 = 0

	for scanner.Scan() {
		game := NewGame(scanner.Text())

		hasEnoughCubes := true
		for color, count := range game.maxMap {
			if count > maxExpected.maxMap[color] {
				hasEnoughCubes = false
				break
			}
		}

		if hasEnoughCubes {
			sumOfIds += game.ID
		}
	}

	fmt.Println(sumOfIds)
}
