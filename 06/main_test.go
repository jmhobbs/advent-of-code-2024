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
	assert.Equal(t, main.Robot{4, 6, '^'}, robot)
}

func Test_Step_Minimal(t *testing.T) {
	t.Run("off the top", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..."),
			[]byte("..."),
		}
		robot := main.Robot{1, 0, '^'}
		assert.True(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{1, -1, '^'}, robot)
	})

	t.Run("off the bottom", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..."),
			[]byte("..."),
		}
		robot := main.Robot{1, 2, 'v'}
		assert.True(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{1, 3, 'v'}, robot)
	})

	t.Run("off the left", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..."),
			[]byte("..."),
		}
		robot := main.Robot{0, 1, '<'}
		assert.True(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{-1, 1, '<'}, robot)
	})

	t.Run("off the right", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..."),
			[]byte("..."),
		}
		robot := main.Robot{2, 1, '>'}
		assert.True(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{3, 1, '>'}, robot)
	})

	t.Run("up - blocked", func(t *testing.T) {
		mmap := main.Map{
			[]byte(".#."),
			[]byte("..."),
			[]byte("..."),
		}
		robot := main.Robot{1, 1, '^'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{2, 1, '>'}, robot)

		assert.Equal(
			t,
			main.Map{
				[]byte(".#."),
				[]byte("..X"),
				[]byte("..."),
			},
			mmap,
		)
	})

	t.Run("right - blocked", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..#"),
			[]byte("..."),
		}
		robot := main.Robot{1, 1, '>'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{1, 2, 'v'}, robot)

		assert.Equal(
			t,
			main.Map{
				[]byte("..."),
				[]byte("..#"),
				[]byte(".X."),
			},
			mmap,
		)
	})

	t.Run("down - blocked", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..."),
			[]byte(".#."),
		}
		robot := main.Robot{1, 1, 'v'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{0, 1, '<'}, robot)

		assert.Equal(
			t,
			main.Map{
				[]byte("..."),
				[]byte("X.."),
				[]byte(".#."),
			},
			mmap,
		)
	})

	t.Run("left - blocked", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("#.."),
			[]byte("..."),
		}
		robot := main.Robot{1, 1, '<'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{1, 0, '^'}, robot)

		assert.Equal(
			t,
			main.Map{
				[]byte(".X."),
				[]byte("#.."),
				[]byte("..."),
			},
			mmap,
		)
	})

	t.Run("up - clear", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..."),
			[]byte("..."),
		}
		robot := main.Robot{1, 1, '^'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{1, 0, '^'}, robot)

		assert.Equal(
			t,
			main.Map{
				[]byte(".X."),
				[]byte("..."),
				[]byte("..."),
			},
			mmap,
		)
	})

	t.Run("right - clear", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..X"),
			[]byte("..."),
		}
		robot := main.Robot{1, 1, '>'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{2, 1, '>'}, robot)
	})

	t.Run("down - clear", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("..."),
			[]byte(".X."),
		}
		robot := main.Robot{1, 1, 'v'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{1, 2, 'v'}, robot)
	})

	t.Run("left - clear", func(t *testing.T) {
		mmap := main.Map{
			[]byte("..."),
			[]byte("X.."),
			[]byte("..."),
		}
		robot := main.Robot{1, 1, '<'}
		assert.False(t, main.Step(&robot, mmap))
		assert.Equal(t, main.Robot{0, 1, '<'}, robot)
	})
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
		assert.False(t, main.Step(&robot, mmap))
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
	assert.False(t, main.Step(&robot, mmap))
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
		assert.False(t, main.Step(&robot, mmap))
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
	assert.False(t, main.Step(&robot, mmap))
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

	main.WalkMap(&robot, mmap)

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
