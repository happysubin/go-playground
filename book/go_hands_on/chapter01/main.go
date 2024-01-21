package main

import (
	"fmt"
	"bufio" //입출력 관련 패키지. 여기에 Reader와 Writer가 포함된다.
	"io"
	"errors"
	"strconv"
	"os"
)


var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|-help]

A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])


func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}
	name, err := getName(r, w)
	if(err != nil) {
		return err
	}
	greetUser(c, name, w)
	return nil
}

type config struct {
	numTimes int //인사 횟수 결정
	printUsage bool //사용법을 보여주는지 여부
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater tahn 0") 
	}
	return nil
}



func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error

	c := config{}
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}
	if args[0] == "-h" || args[0] == "-help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0]) //숫자열의 문자열을 그에 해당하는 정숫 값으로 변환

	if err != nil {
		return c, err
	}

	c.numTimes = numTimes
	return c, nil
}

//os.Stdin, os.Stdout가 대표적인 Reader, Writer 구현체
func getName(r io.Reader, w io.Writer) (string, error) { //return 값이 string과 error
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}
	return name, nil //nil은 해당 타입의 영제 값이나 초기화되지 않은 포인터를 나타냅니다.
}


func main() {
	//os.Args는 프로그램을 실행할 때 전달된 모든 명령행 인수를 포함하는 문자열 슬라이스입니다. 
	//os.Args[1:]는 첫 번째 요소 이후의 모든 명령행 인수를 나타냅니다.
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

//go build -o application

// ./application
// ./application -h
// ./application 5