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
	dat, err := os.ReadFile("input.txt")
	check(err)
	datString := string(dat)
	lines := strings.Split(datString, "\n")
	fmt.Println("parsing games")
	games := make([]game, len(lines))
	for index, line := range lines {
		games[index] = parseGame(line)
	}
	fmt.Println("calculating min viable cubes")
	sum := 0
	for _, game := range games {
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, set := range game.sets {
			if set.red > minRed {
				minRed = set.red
			}
			if set.green > minGreen {
				minGreen = set.green
			}
			if set.blue > minBlue {
				minBlue = set.blue
			}
		}
		fmt.Println("game " + strconv.Itoa(game.id) + " requires at least " + strconv.Itoa(minRed) + " red " + strconv.Itoa(minGreen) + " green " + strconv.Itoa(minBlue) + " blue ")
		power := minRed * minGreen * minBlue
		sum = sum + power
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
