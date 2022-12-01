package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

var buffer bytes.Buffer

func expect() string {
	output := `🛠 Setup for Url: https://example.com/
🛠 Setup for Method: GET
🛠 Setup for Request: 10
🛠 Setup for Duration: 10
`
	return output
}

func TestMain(t *testing.T) {
	outStream = &buffer
	expectBuf := bytes.NewBufferString(expect()).Bytes()
	arg := []string{"", "testing"}
	t.Run("Test main function", func(t *testing.T) {
		os.Args = arg
		main()
		if !bytes.Equal(buffer.Bytes(), expectBuf) {
			fmt.Println("📝 The Results not as Expected 📝")
			t.Errorf(buffer.String())
		} else {
			fmt.Println("🎉 The Results as Expected 🎉")
		}
	})
	buffer.Reset()
}

func outputTest() string {
	output := `Success
`
	return output
}

func TestAttack(t *testing.T) {
	outStream = &buffer
	t.Run("Test Attack function", func(t *testing.T) {
		variable := &Variable{
			url:      "https://example.com/",
			method:   "GET",
			request:  10,
			duration: 10,
		}
		atteckOperation(variable)
		if buffer.String() == outputTest() {
			fmt.Println("🎉 Good News, Function running work fine 🎉")
		} else {
			t.Errorf(buffer.String())
		}
	})
}
