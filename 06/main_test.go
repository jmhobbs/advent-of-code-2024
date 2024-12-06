package main_test

import (
	"strings"
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/06"

	"github.com/stretchr/testify/assert"
)

func Test_ParseInput(t *testing.T) {
	mmap, robot, err := main.ParseInput(strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`))

	assert.Nil(t, err)
	assert.Equal(
		t,
		main.Map{
			[]byte("....#....."),
			[]byte(".........#"),
			[]byte(".........."),
			[]byte("..#......."),
			[]byte(".......#.."),
			[]byte(".........."),
			[]byte(".#..X....."),
			[]byte("........#."),
			[]byte("#........."),
			[]byte("......#..."),
		},
		mmap,
	)
	assert.Equal(t, main.Robot{4, 6, '?'}, robot)
}

func Test_Step(t *testing.T) {
	mmap := main.Map{
		[]byte("....#....."),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte("..#......."),
		[]byte(".......#.."),
		[]byte(".........."),
		[]byte(".#..X....."),
		[]byte("........#."),
		[]byte("#........."),
		[]byte("......#..."),
	}

	robot := main.Robot{4, 6, '^'}

	// Following the above protocol, the guard moves up several times until she reaches an obstacle
	for i := 0; i < 5; i++ {
		assert.False(t, main.Step(robot, mmap))
	}
	assert.Equal(t, main.Robot{4, 1, '^'}, robot)
	assert.Equal(t, main.Map{
		[]byte("....#....."),
		[]byte("....X....#"),
		[]byte("....X....."),
		[]byte("..#.X....."),
		[]byte("....X..#.."),
		[]byte("....X....."),
		[]byte(".#..X....."),
		[]byte("........#."),
		[]byte("#........."),
		[]byte("......#..."),
	}, mmap)

	// Because there is now an obstacle in front of the guard, she turns right before continuing straight in her new facing direction:
	assert.False(t, main.Step(robot, mmap))
	assert.Equal(t, main.Robot{5, 1, '>'}, robot)
	assert.Equal(t, main.Map{
		[]byte("....#....."),
		[]byte("....XX...#"),
		[]byte("....X....."),
		[]byte("..#.X....."),
		[]byte("....X..#.."),
		[]byte("....X....."),
		[]byte(".#..X....."),
		[]byte("........#."),
		[]byte("#........."),
		[]byte("......#..."),
	}, mmap)

	for i := 0; i < 3; i++ {
		assert.False(t, main.Step(robot, mmap))
	}
	assert.Equal(t, main.Robot{8, 1, '>'}, robot)
	assert.Equal(t, main.Map{
		[]byte("....#....."),
		[]byte("....XXXXX#"),
		[]byte("....X....."),
		[]byte("..#.X....."),
		[]byte("....X..#.."),
		[]byte("....X....."),
		[]byte(".#..X....."),
		[]byte("........#."),
		[]byte("#........."),
		[]byte("......#..."),
	}, mmap)

	// Reaching another obstacle (a spool of several very long polymers), she turns right again and continues downward
	assert.False(t, main.Step(robot, mmap))
	assert.Equal(t, main.Robot{8, 2, 'v'}, robot)
	assert.Equal(t, main.Map{
		[]byte("....#....."),
		[]byte("....XXXXX#"),
		[]byte("....X...X."),
		[]byte("..#.X....."),
		[]byte("....X..#.."),
		[]byte("....X....."),
		[]byte(".#..X....."),
		[]byte("........#."),
		[]byte("#........."),
		[]byte("......#..."),
	}, mmap)
}

func Test_WalkMap(t *testing.T) {
	mmap := main.Map{
		[]byte("....#....."),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte("..#......."),
		[]byte(".......#.."),
		[]byte(".........."),
		[]byte(".#..X....."),
		[]byte("........#."),
		[]byte("#........."),
		[]byte("......#..."),
	}

	robot := main.Robot{4, 6, '^'}

	main.WalkMap(robot, mmap)

	assert.Equal(t, main.Robot{7, 10, 'v'}, robot)
	assert.Equal(t, main.Map{
		[]byte("....#....."),
		[]byte("....XXXXX#"),
		[]byte("....X...X."),
		[]byte("..#.X...X."),
		[]byte("..XXXXX#X."),
		[]byte("..X.X.X.X."),
		[]byte(".#XXXXXXX."),
		[]byte(".XXXXXXX#."),
		[]byte("#XXXXXXX.."),
		[]byte("......#X.."),
	}, mmap)
}

func Test_CountVisitedPositions(t *testing.T) {
	visitedMap := main.Map{
		[]byte("....#....."),
		[]byte("....XXXXX#"),
		[]byte("....X...X."),
		[]byte("..#.X...X."),
		[]byte("..XXXXX#X."),
		[]byte("..X.X.X.X."),
		[]byte(".#XXXXXXX."),
		[]byte(".XXXXXXX#."),
		[]byte("#XXXXXXX.."),
		[]byte("......#X.."),
	}
	assert.Equal(t, 41, main.CountVisitedPositions(visitedMap))
}
