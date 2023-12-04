package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id                  int
	winners, candidates []int
}

func toIntSlice(s []string) ([]int, error) {
	var ints []int
	for _, str := range s {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ints = append(ints, num)
	}
	return ints, nil
}

func createCard(lineOfText string) Card {
	//fmt.Println(lineOfText)

	// Finding out card number
	re := regexp.MustCompile(`\s+`)
	colonSplit := strings.Split(lineOfText, ":")
	spaceSplit := re.Split(colonSplit[0], -1)
	cardCounter, err := strconv.Atoi(spaceSplit[1])

	re = regexp.MustCompile(`\s+`)
	pipeSplit := strings.Split(colonSplit[1], "|")

	winners, err := toIntSlice(re.Split(strings.TrimSpace(pipeSplit[0]), -1))
	if err != nil {
		fmt.Println("Error converting winners to integers:", err)
	}
	candidates, err := toIntSlice(re.Split(strings.TrimSpace(pipeSplit[1]), -1))
	if err != nil {
		fmt.Println("Error converting candidates to integers:", err)

	}

	var card = Card{id: cardCounter, winners: winners, candidates: candidates}
	return card

}

func parseCards(path string) []Card {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var cards []Card
	var cardNumber = 1

	for fileScanner.Scan() {
		//fmt.Println(cardNumber)
		cards = append(cards, createCard(fileScanner.Text()))
		cardNumber++
	}

	readFile.Close()

	return cards

}

func calculatePoints(cards []Card) int {

	var points int = 0

	for _, card := range cards {
		var wins int = 0
		for _, winner := range card.winners {
			for _, candidate := range card.candidates {
				if winner == candidate {
					wins++
					break
				}
			}
		}
		points = points + max(1, wins)*2
	}

	return points
}

func main() {

	var cards []Card
	cards = parseCards("C:\\Git\\advent-of-code-23\\4\\input.txt")
	points := calculatePoints(cards)

	fmt.Println(points)

	if false {
		fmt.Print(cards)
	}

}
