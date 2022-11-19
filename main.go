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

var message string = `url: (https://example.com)
method: (GET, PUT, POST, DELETE)
body: 
header:
request:
duration: (minutes)
`

func main() {
	if arg := os.Args; arg[1] == "testing" {
		fmt.Println("Welcome to Loadtest")
		fmt.Println(message)
		scanner := bufio.NewScanner(os.Stdin)
		variable := &Variable{}
		Loadtest(variable, scanner)
	}
}

func Loadtest(variable *Variable, scanner *bufio.Scanner) {
	exit := false
	for !exit && scanner.Scan() {
		input := scanner.Text()
		arg := convert2Slice(input)
		switch {
		case arg[0] == "url" && len(arg) == 2:
			variable.url = arg[1]
		case arg[0] == "method" && len(arg) == 2:
			variable.method = arg[1]
		case arg[0] == "body" && len(arg) == 2:
			b := []byte(arg[1])
			variable.body = b
		// case arg[0] == "header" && len(arg) == 2:
		// 	variable.header = arg[1]
		case arg[0] == "request" && len(arg) == 2:
			r, err := strconv.Atoi(arg[1])
			if err == nil {
				variable.request = r
			} else {
				fmt.Println("is that INTEGER ??")
			}
		case arg[0] == "duration" && len(arg) == 2:
			t, err := strconv.Atoi(arg[1])
			if err == nil {
				variable.duration = t
			} else {
				fmt.Println("is that NUMBER ??")
			}
		case arg[0] == "attack" && len(arg) == 1:
			atteckOperation(variable)
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
