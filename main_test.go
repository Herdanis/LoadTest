package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	arg := []string{"testing"}
	t.Run("Test main function", func(t *testing.T) {
		os.Args = arg
		main()
	})
}
