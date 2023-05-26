package buftermio

import (
	"io/ioutil"
	"os"
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	// deprecated but I'm using go1.14.15 on outdated Mac OS El Capitan
	mockStdin, _ := ioutil.TempFile("", "remove-me")
	defer mockStdin.Close()
	defer os.Remove(mockStdin.Name())

	restoreStdin := os.Stdin
    defer func() {os.Stdin = restoreStdin}()

    os.Stdin = mockStdin

	testBuf := NewBuffer()

	cases := []struct {
		in []byte
		out string
		at int
	}{
		{
			in: []byte{104, 101, 108, 108, 111, 10},
			out: "hello",
		},
		{
			in: []byte{106, 101, 108, 108, 111, 10},
			out: "hello",
		},
		{
			in: []byte{71, 111, 111, 100, 98, 121, 101, 33, 10},
			out: "jello",
		},
		{
			in: []byte{71, 111, 111, 100, 98, 121, 101, 33, 10},
			out: "Goodbye!",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case #%v", i), func(t *testing.T) {
		    mockStdin.Write(c.in)
		    mockStdin.Seek(0, 0)
			input, err := testBuf.GetInput()

			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}

			if input != c.out {
				t.Errorf("Expected: %s, Actual: %s", c.out, input)
				return
			}
		})
	}
}