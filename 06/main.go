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

	mmap, robot, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	WalkMap(&robot, mmap)

	fmt.Printf("A: %d\n", CountVisitedPositions(mmap))
}

type Robot struct {
	X, Y        int
	Orientation rune
}

type Map [][]byte

func (m Map) String() string {
	var s string
	for _, line := range m {
		s += string(line) + "\n"
	}
	return s
}

func ParseInput(in io.Reader) (Map, Robot, error) {
	var (
		robot Robot
		line  []byte
		lines [][]byte
	)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line = make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())

		for x, c := range line {
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				robot = Robot{x, len(lines), rune(c)}
				line[x] = 'X'
			}
		}
		lines = append(lines, line)
	}

	return Map(lines), robot, scanner.Err()
}

func Step(robot *Robot, mmap Map) bool {
	for i := 0; i < 4; i++ {
		switch robot.Orientation {
		case '^':
			if robot.Y == 0 {
				robot.Y--
				return true
			}
			if mmap[robot.Y-1][robot.X] == '#' {
				robot.Orientation = '>'
				continue
			}
			robot.Y--
			mmap[robot.Y][robot.X] = 'X'
			return false
		case '>':
			if robot.X == len(mmap[robot.Y])-1 {
				robot.X++
				return true
			}
			if mmap[robot.Y][robot.X+1] == '#' {
				robot.Orientation = 'v'
				continue
			}
			robot.X++
			mmap[robot.Y][robot.X] = 'X'
			return false
		case 'v':
			if robot.Y == len(mmap)-1 {
				robot.Y++
				return true
			}
			if mmap[robot.Y+1][robot.X] == '#' {
				robot.Orientation = '<'
				continue
			}
			robot.Y++
			mmap[robot.Y][robot.X] = 'X'
			return false
		case '<':
			if robot.X == 0 {
				robot.X--
				return true
			}
			if mmap[robot.Y][robot.X-1] == '#' {
				robot.Orientation = '^'
				continue
			}
			robot.X--
			mmap[robot.Y][robot.X] = 'X'
			return false
		}
	}
	panic("robot deadlocked")
}

func WalkMap(robot *Robot, mmap Map) {
	for {
		if Step(robot, mmap) {
			break
		}
	}
}

func CountVisitedPositions(mmap Map) int {
	ctr := 0
	for _, line := range mmap {
		for _, c := range line {
			if c == 'X' {
				ctr++
			}
		}
	}
	return ctr
}
