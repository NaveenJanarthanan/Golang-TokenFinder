package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//IsLetter checks for letter
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {

	if len(os.Args) >= 2 {
		fmt.Println("please run program, then enter file name")
		return
	}
	fmt.Print("Enter file name: ")

	var first string
	fmt.Scanln(&first)

	f, err := os.Open(first)
	check(err)

	file, err := os.Create("input.out")
	check(err)

	writer := bufio.NewWriter(file)

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	fmt.Println("Results in Output File input.out")
	var token int
	for scanner.Scan() {

		if string(scanner.Text()[0]) == "$" {
			fmt.Fprintf(writer, "ID[STRING]: ")
			a := strings.Trim(scanner.Text(), `$`)
			fmt.Fprintln(writer, a)
			token++

		} else if string(scanner.Text()[0]) == "#" {
			fmt.Fprintf(writer, "ID[INT]: ")
			b := strings.Trim(scanner.Text(), `#`)
			fmt.Fprintln(writer, b)
			token++
		} else if string(scanner.Text()[0]) == "%" {
			fmt.Fprintf(writer, "ID[REAL]: ")
			c := strings.Trim(scanner.Text(), `%`)
			fmt.Fprintln(writer, c)
			token++
		} else if string(scanner.Text()[0]) == ("\"") && strings.HasSuffix(scanner.Text(), "\"") {
			fmt.Fprintf(writer, "STRING: ")
			s := strings.Trim(scanner.Text(), `"`)
			fmt.Fprintln(writer, s)
			token++
		} else if string(scanner.Text()[0]) == "B" {
			fmt.Fprintln(writer, "BEGIN")
			token++
		} else if scanner.Text() == "<=" {
			fmt.Fprintln(writer, "ASSIGN")
			token++
		} else if string(scanner.Text()[0]) == ("\"") {
			fmt.Fprintf(writer, "STRING: ")
			s := strings.Trim(scanner.Text(), `"`)
			fmt.Fprintf(writer, s)
			token++
		} else if strings.HasSuffix(scanner.Text(), "\"") {
			s := strings.Trim(scanner.Text(), `"`)
			fmt.Fprintln(writer, s)

		} else if scanner.Text() == ":" {
			fmt.Fprintln(writer, "COLON ")

			token++
		} else if string(scanner.Text()[0]) == "(" {
			fmt.Fprintln(writer, "LPAREN ")
			token++
		} else if string(scanner.Text()[0]) == ")" {
			fmt.Fprintln(writer, "RPAREN ")
			token++
		} else if scanner.Text() == "+" {
			fmt.Fprintln(writer, "PLUS ")
			token++
		} else if scanner.Text() == "*" {
			fmt.Fprintln(writer, "TIMES ")
			token++
		} else if scanner.Text() == "/" {
			fmt.Fprintln(writer, "DIVISION ")
			token++
		} else if scanner.Text() == "-" {
			fmt.Fprintln(writer, "MINUS ")
			token++
		} else if scanner.Text() == "^" {
			fmt.Fprintln(writer, "POWER")
			token++
		} else if scanner.Text() == "WRITE" {
			fmt.Fprintln(writer, "WRITE")
			token++
		} else if scanner.Text() == "END" {
			fmt.Fprintln(writer, "END")
			token++
		} else if IsLetter(scanner.Text()) == true {
			fmt.Fprintf(writer, scanner.Text())
		} else {
			if _, err := strconv.Atoi(scanner.Text()); err == nil {

				fmt.Fprintf(writer, "INT_CONST: ")
				fmt.Fprintln(writer, scanner.Text())
				token++
			} else {

				fmt.Fprintf(writer, "REAL_CONST: ")
				fmt.Fprintln(writer, scanner.Text())
				token++

			}
		}
	}
	fmt.Println(token, "Tokens Produced")
	writer.Flush()
	f.Close()

}
