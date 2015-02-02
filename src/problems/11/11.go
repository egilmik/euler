package main

import "fmt"
import "os"
import "log"
import "bufio"
import "strings"
import "strconv"

func main() {
	grid := getGrid()
	var sum uint64
	var largest uint64
	sum = 0
	largest = 0
	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {

			sum = calulateLargestForIndex(grid, x, y)
			if sum > largest {
				largest = sum
			}
		}

	}
	fmt.Println(largest)
}

func calulateLargestForIndex(grid [][]uint64, x int, y int) uint64 {
	var largest uint64
	largest = 0
	var value uint64

	down := len(grid) > y+3
	right := len(grid[y]) > x+3
	up := y > 2
	//Right
	if right {
		value = grid[y][x] * grid[y][x+1] * grid[y][x+2] * grid[y][x+3]
		if value > largest {
			largest = value
		}
	}

	//Down
	if down {
		value = grid[y][x] * grid[y+1][x] * grid[y+2][x] * grid[y+3][x]
		if value > largest {
			largest = value
		}
	}

	if right && down {
		value = grid[y][x] * grid[y+1][x+1] * grid[y+2][x+2] * grid[y+3][x+3]
		if value > largest {
			largest = value
		}
	}

	if up && right {
		value = grid[y][x] * grid[y-1][x+1] * grid[y-2][x+2] * grid[y-3][x+3]
		if value > largest {
			largest = value
		}
	}

	return largest
}

func getGrid() [][]uint64 {
	file, err := os.Open("grid.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []uint64
		line := scanner.Text()
		splits := strings.Split(line, " ")
		for _, element := range splits {
			val, _ := strconv.ParseUint(element, 10, 8)
			row = append(row, val)
		}
		grid = append(grid, row)
	}
	return grid
}
