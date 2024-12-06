package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jmhobbs/advent-of-code-2024/util"
)

func main() {
	f := util.OpenInput()
	defer f.Close()

	puzzle, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("A: %d\n", CountXmas(puzzle))
	fmt.Printf("B: %d\n", CountXMas(puzzle))
}

func ParseInput(in io.Reader) ([][]byte, error) {
	lines := [][]byte{}

	var buf []byte
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		buf = scanner.Bytes()
		line := make([]byte, len(buf))
		copy(line, buf)
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

// !! assumption: all rows will be the same length
func CountXmas(puzzle [][]byte) int {
	var count int

	for y := 0; y < len(puzzle); y++ {
		row := puzzle[y]
		for x := 0; x < len(row); x++ {
			if row[x] == 'X' {
				// backward
				if x >= 3 {
					if row[x-1] == 'M' && row[x-2] == 'A' && row[x-3] == 'S' {
						count += 1
					}
				}

				if y >= 3 {
					// up
					if puzzle[y-1][x] == 'M' && puzzle[y-2][x] == 'A' && puzzle[y-3][x] == 'S' {
						count += 1
					}
					// diagonal up-right
					if len(row)-x >= 4 {
						if puzzle[y-1][x+1] == 'M' && puzzle[y-2][x+2] == 'A' && puzzle[y-3][x+3] == 'S' {
							count += 1
						}
					}
					// diagonal up-left
					if x >= 3 {
						if puzzle[y-1][x-1] == 'M' && puzzle[y-2][x-2] == 'A' && puzzle[y-3][x-3] == 'S' {
							count += 1
						}
					}
				}

				if len(puzzle)-y >= 4 {
					// down
					if puzzle[y+1][x] == 'M' && puzzle[y+2][x] == 'A' && puzzle[y+3][x] == 'S' {
						count += 1
					}
					// diagonal down-right
					if len(row)-x >= 4 {
						if puzzle[y+1][x+1] == 'M' && puzzle[y+2][x+2] == 'A' && puzzle[y+3][x+3] == 'S' {
							count += 1
						}
					}
					// diagonal down-left
					if x >= 3 {
						if puzzle[y+1][x-1] == 'M' && puzzle[y+2][x-2] == 'A' && puzzle[y+3][x-3] == 'S' {
							count += 1
						}
					}
				}
				// forward - comes last since we can optimize x increment if we know we have an XMAS
				if len(row)-x >= 4 {
					if row[x+1] == 'M' && row[x+2] == 'A' && row[x+3] == 'S' {
						count += 1
						x += 3
					}
				}
			}
		}
	}
	return count
}

func CountXMas(puzzle [][]byte) int {
	var count int

	// MAS requires at least one line above and one below
	for y := 1; y < len(puzzle)-1; y++ {
		row := puzzle[y]
		// MAS requires at least one column before and one after
		for x := 1; x < len(row)-1; x++ {
			if row[x] == 'A' {
				// high left -> low right
				highToLow := (puzzle[y-1][x-1] == 'S' && puzzle[y+1][x+1] == 'M') || (puzzle[y-1][x-1] == 'M' && puzzle[y+1][x+1] == 'S')
				// low left -> high right
				lowToHigh := (puzzle[y+1][x-1] == 'S' && puzzle[y-1][x+1] == 'M') || (puzzle[y+1][x-1] == 'M' && puzzle[y-1][x+1] == 'S')
				if highToLow && lowToHigh {
					count += 1
				}
			}
		}
	}

	return count
}
