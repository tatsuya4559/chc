package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/mattn/go-isatty"
)

func printHelp() {
	usage := `usage: chc -c <case> <word>
chc change word to specified case.

	-c=STRING	Style to change to.
			[s[nake], u[pper], c[amel], p[ascal]] (default: snake)
`
	fmt.Fprintf(os.Stderr, "%s\n", usage)
	os.Exit(0)
}

const (
	SNAKE  = "SNAKE"
	UPPER  = "UPPER"
	CAMEL  = "CAMEL"
	PASCAL = "PASCAL"
	LISP   = "LISP"
)

func main() {
	flag.Usage = printHelp
	casePtr := flag.String("c", SNAKE, "case to change to")

	flag.Parse()

	if isatty.IsTerminal(os.Stdin.Fd()) {
		// interactive
		for _, word := range flag.Args() {
			fmt.Println(ChangeCase(*casePtr, word))
		}
	} else {
		// pipe
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			word := strings.TrimSpace(scanner.Text())
			fmt.Println(ChangeCase(*casePtr, word))
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("error while reading pipe: %v", err)
		}
	}
}

// ChangeCase changes case-style of word.
func ChangeCase(opt, word string) string {
	switch strings.ToLower(opt) {
	case "s", "snake":
		return toSnake(word)
	case "u", "upper":
		return toUpper(word)
	case "c", "camel":
		return toCamel(word)
	case "p", "pascal":
		return toPascal(word)
	case "l", "lisp":
		return toLisp(word)
	default:
		return toSnake(word)
	}
}

func determineCase(word string) string {
	if strings.Contains(word, "-") {
		return LISP
	}
	if word == strings.ToLower(word) {
		return SNAKE
	}
	if word == strings.ToUpper(word) {
		return UPPER
	}
	if unicode.IsLower(rune(word[0])) {
		return CAMEL
	}
	return PASCAL
}

// tokenize splits word considering style of word.
func tokenize(word string) []string {
	switch determineCase(word) {
	case SNAKE:
		return strings.Split(word, "_")
	case UPPER:
		return strings.Split(strings.ToLower(word), "_")
	case LISP:
		return strings.Split(strings.ToLower(word), "-")
	}

	rword := []rune(word)
	if len(rword) < 2 {
		return []string{word}
	}
	var result []string
	var start int
	for i := 1; i < len(rword); i++ {
		if unicode.IsUpper(rword[i]) {
			result = append(result, strings.ToLower(string(rword[start:i])))
			start = i
		}
	}
	result = append(result, strings.ToLower(string(rword[start:])))
	return result
}

// toSnake changes word to snake_case.
func toSnake(word string) string {
	return strings.Join(tokenize(word), "_")
}

// toUpper changes word to UPPER_CASE.
func toUpper(word string) string {
	return strings.ToUpper(strings.Join(tokenize(word), "_"))
}

// toCamel changes word to camelCase.
func toCamel(word string) string {
	tokens := tokenize(word)
	switch len(tokens) {
	case 0:
		return ""
	case 1:
		return tokens[0]
	}
	var b strings.Builder
	b.WriteString(tokens[0])
	for _, t := range tokens[1:] {
		b.WriteString(title(t))
	}
	return b.String()
}

// toPascal changes word to PascalCase.
func toPascal(word string) string {
	var b strings.Builder
	for _, t := range tokenize(word) {
		b.WriteString(title(t))
	}
	return b.String()
}

// toLisp changes word to lisp-case.
func toLisp(word string) string {
	return strings.Join(tokenize(word), "-")
}

// title capitalizes first letter of argument.
func title(s string) string {
	if len(s) < 2 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
