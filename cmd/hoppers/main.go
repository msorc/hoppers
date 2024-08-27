package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/msorc/hoppers/pkg/hoppers"
)

var currentLine int

func main() {
	processInput()
}

func processInput() {
    reader := bufio.NewReader(os.Stdin)

	numbers, err := parseLine(reader, 1)
	if err != nil { panic(err) }
	testsCount := numbers[0]

	for i := 0; i < testsCount; i++ {
		numbers, err := parseLine(reader, 2)
		if err != nil { panic(err) }
		width, height := numbers[0], numbers[1]

		numbers, err = parseLine(reader, 4)
		if err != nil { panic(err) }
		start := hoppers.Point{numbers[0], numbers[1]}
		finish := hoppers.Point{numbers[2], numbers[3]}

		numbers, err = parseLine(reader, 1)
		if err != nil { panic(err) }
		obstaclesCount := numbers[0]
		obstacles := make([][2]hoppers.Point, obstaclesCount)

		for j := 0; j < obstaclesCount; j++ {
			numbers, err = parseLine(reader, 4)
			if err != nil { panic(err) }
			s := hoppers.Point{numbers[0], numbers[2]}
			f := hoppers.Point{numbers[1], numbers[3]}

			obstacles = append(obstacles, [2]hoppers.Point{s, f})
		}
		
		hops, err := hoppers.HopsCount(width, height, start, finish, obstacles)
		fmt.Println(hoppers.PresentResult(hops, err))
	}
}

func parseLine(reader *bufio.Reader, count int) ([]int, error) {
	currentLine++
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Line %d: %w", currentLine, err)
	}

	elements := strings.Fields(text)
	if len(elements) != count {
		return nil, fmt.Errorf("Line %d: arguments count mismatch. Expected: %d, got: %d in: %s",
			currentLine, count, len(elements), text)
	}
	
	numbers := make([]int, len(elements))

	for i, s := range elements {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("Line %d: %w", currentLine, err)
		}
		numbers[i] = num
	}

	return numbers, nil
}


