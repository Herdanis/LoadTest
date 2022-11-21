package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

var message string = `* url: (https://example.com)
* method: (GET, PUT, POST, DELETE)
* duration: (minutes)
* request: (request per second)
body: 
header:
attack for execute
use exit to quit from aplication

* : is required
`

func main() {
	argLen := len(os.Args)
	scanner := bufio.NewScanner(os.Stdin)
	switch {
	case argLen > 1:
		inputUnitTest, err := os.Open("c_mainTest.txt")
		if err != nil {
			panic(err)
		}
		defer inputUnitTest.Close()
		scanner = bufio.NewScanner(inputUnitTest)
	default:
		fmt.Println("Welcome to Loadtest")
		fmt.Println(message)
	}
	variable := &Variable{}
	Loadtest(variable, scanner)
}

func Loadtest(variable *Variable, scanner *bufio.Scanner) {
	exit := false
	for !exit && scanner.Scan() {
		input := scanner.Text()
		arg := convert2Slice(input)
		switch {
		case arg[0] == "url" && len(arg) == 2:
			variable.url = arg[1]
			fmt.Println("🛠 Setup for Url:", arg[1])
		case arg[0] == "method" && len(arg) == 2:
			variable.method = arg[1]
			fmt.Println("🛠 Setup for Method:", arg[1])
		case arg[0] == "body" && len(arg) == 2:
			b := []byte(arg[1])
			variable.body = b
			fmt.Println("🛠 Setup for Body:", arg[1])
		// case arg[0] == "header" && len(arg) == 2:
		// 	variable.header = arg[1]
		case arg[0] == "request" && len(arg) == 2:
			r, err := strconv.Atoi(arg[1])
			if err == nil {
				variable.request = r
				fmt.Println("🛠 Setup for Request:", arg[1])
			} else {
				fmt.Println("is that INTEGER ??")
			}
		case arg[0] == "duration" && len(arg) == 2:
			t, err := strconv.Atoi(arg[1])
			if err == nil {
				variable.duration = t
				fmt.Println("🛠 Setup for Duration:", arg[1])
			} else {
				fmt.Println("is that NUMBER ??")
			}
		case arg[0] == "attack" && len(arg) == 1:
			if variable.url == "" {
				fmt.Println("Require the URL")
			} else if variable.method == "" {
				fmt.Println("Require the Method")
			} else if variable.request == 0 {
				fmt.Println("Require the Request")
			} else if variable.duration == 0 {
				fmt.Println("Require the Duration")
			} else {
				atteckOperation(variable)
			}
		case arg[0] == "check" && len(arg) == 1:
			fmt.Println(*variable)
		case arg[0] == "exit" && len(arg) == 1:
			exit = true
		}
	}
}

func atteckOperation(variable *Variable) {
	requestRate := vegeta.Rate{Freq: variable.request, Per: time.Second}
	duration := time.Duration(variable.duration) * time.Minute
	target := target(variable)
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for result := range attacker.Attack(target, requestRate, duration, "Loadtest") {
		metrics.Add(result)
	}
	metrics.Close()
	reporter := vegeta.NewTextReporter(&metrics)
	reporter(os.Stdout)
}

func target(variable *Variable) vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}
		tgt.URL = variable.url
		tgt.Method = variable.method
		tgt.Body = variable.body
		tgt.Header = variable.header
		return nil
	}
}

func convert2Slice(input string) []string {
	slice := strings.Split(input, " ")
	return slice
}
