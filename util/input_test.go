package util_test

import (
	"io/fs"
	"os"
	"testing"

	"github.com/jmhobbs/advent-of-code-2024/util"
	"github.com/stretchr/testify/assert"
)

func Test_OpenInput(t *testing.T) {
	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmp)

	err = os.Chdir(tmp)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("missing file", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(*fs.PathError)
				if !ok || err.Err.Error() != "no such file or directory" {
					t.Errorf("Unexpected panic: %v", r)
				}
			} else {
				t.Error("Expected panic, none received")
			}
		}()

		util.OpenInput()
	})

	t.Run("file exists", func(t *testing.T) {
		f, err := os.Create("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		f.Close()

		defer os.Remove("input.txt")

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Unexpected panic: %v", r)
			}
		}()

		r := util.OpenInput()
		defer r.Close()
	})
}

func Test_ReadInput(t *testing.T) {
	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmp)

	err = os.Chdir(tmp)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("missing file", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(*fs.PathError)
				if !ok || err.Err.Error() != "no such file or directory" {
					t.Errorf("Unexpected panic: %v", r)
				}
			} else {
				t.Error("Expected panic, none received")
			}
		}()

		util.ReadInput()
	})

	t.Run("file exists", func(t *testing.T) {
		f, err := os.Create("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		f.Close()

		defer os.Remove("input.txt")

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Unexpected panic: %v", r)
			}
		}()

		buf := util.ReadInput()
		assert.Equal(t, []byte{}, buf)
	})
}
