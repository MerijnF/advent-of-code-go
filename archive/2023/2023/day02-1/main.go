package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("loading game data")
	dat, err := os.ReadFile("../../input/2023/day02/input.txt")
	check(err)
	datString := string(dat)
	lines := strings.Split(datString, "\n")
	fmt.Println("parsing games")
	games := make([]game, len(lines))
	for index, line := range lines {
		games[index] = parseGame(line)
	}
	fmt.Println("checking games for 12 red 13 green 14 blue")
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	sum := 0
	for _, game := range games {
		validGame := true
		for _, set := range game.sets {
			if set.red > maxRed {
				validGame = false
				break
			}
			if set.green > maxGreen {
				validGame = false
				break
			}
			if set.blue > maxBlue {
				validGame = false
				break
			}
		}
		if validGame {
			fmt.Println("game " + strconv.Itoa(game.id) + " is valid")
			sum = sum + game.id
		} else {
			fmt.Println("game " + strconv.Itoa(game.id) + " is not valid")
		}
	}
	fmt.Println("final value: " + strconv.Itoa(sum))
}

func parseGame(line string) game {
	idSetsSplit := strings.Split(line, ":")
	gameIdString := idSetsSplit[0]
	id, err := parseGameId(gameIdString)
	check(err)
	setsString := idSetsSplit[1]
	setsSplit := strings.Split(setsString, ";")
	sets := make([]set, len(setsSplit))
	for index, setString := range setsSplit {
		sets[index] = parseSet(setString)
	}
	return game{id: id, sets: sets}
}

func parseGameId(gameIdString string) (int, error) {
	match := regexp.MustCompile(`[0-9]+`)
	return strconv.Atoi(match.FindString(gameIdString))
}

func parseSet(setString string) set {
	red := 0
	green := 0
	blue := 0
	matchColor := regexp.MustCompile(`red|green|blue`)
	matchCount := regexp.MustCompile(`[0-9]+`)
	colorsSplit := strings.Split(setString, ",")
	for _, colorString := range colorsSplit {
		color := matchColor.FindString(colorString)
		count, err := strconv.Atoi(matchCount.FindString(colorString))
		check(err)
		switch color {
		case "red":
			red = count
		case "green":
			green = count
		case "blue":
			blue = count
		}
	}
	return set{blue: blue, red: red, green: green}
}

type game struct {
	id   int
	sets []set
}

type set struct {
	blue  int
	red   int
	green int
}
