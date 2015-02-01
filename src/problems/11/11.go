package main

import "fmt"
import "os"
import "log"
import "bufio"
import "strings"
import "strconv"

var numbers = 4

func main(){
  grid := getGrid()
  fmt.Println(calulateLargestForIndex(grid,0,0))
}

func calulateLargestForIndex(grid [][]uint64, x int, y int) uint64 {
  if len(grid) > y {
    line := grid[y]
    return line[x]

  }
  return 0
}

func getGrid() [][]uint64 {
  file, err := os.Open("grid.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  grid := make([][]uint64,0)

  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    var row []uint64
    line := scanner.Text()
    splits := strings.Split(line," ")
    for _, element := range splits{
      val, _ :=strconv.ParseUint(element,10,8)
      row = append(row,val)
    }
    grid = append(grid,row)
  }
  return grid
}
