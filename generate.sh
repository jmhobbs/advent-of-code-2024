#!/bin/bash -

set -o errexit
set -o nounset
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <day-number>"
    exit 1
fi

dirname="$(printf "%02d" "$1")"
fullpath="$__dir/$dirname"


if [ -x "$fullpath" ]; then
    echo "Directory for day $1 already exists"
    exit 1
fi

mkdir -p "$fullpath"
touch "$fullpath/problem.md"

cat <<EOF > "$fullpath/main.go"
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
    fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func PartOne() int {
  return 0
}
EOF

cat <<EOF > "$fullpath/main_test.go"
package main_test

import (
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/$dirname"

	"github.com/stretchr/testify/assert"
)

func Test_Thing(t *testing.T) {
  assert.Equal(t, 1, main.PartOne())
}
EOF
