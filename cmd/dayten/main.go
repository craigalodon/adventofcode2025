package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: go run . <path/to/input/file>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	presses := 0

	for scanner.Scan() {
		line := scanner.Text()
		machine, err := deserialize(line)
		if err != nil {
			return fmt.Errorf("error deserializing line: %w", err)
		}
		presses += machine.configure()
	}

	fmt.Printf("Configured all machines with %d presses\n", presses)

	return nil
}

type Machine struct {
	indicators          []bool
	buttons             []map[int]bool
	joltageRequirements map[int]bool
}

func (m *Machine) configure() int {
	return 1
}

type ParserState int

const (
	null ParserState = iota
	bracketsOpened
	bracketsClosed
	parenthesesOpened
	parenthesesClosed
	curlyBracesOpened
	curlyBracesClosed
)

func deserialize(s string) (*Machine, error) {

	parserState := null
	var builder strings.Builder

	indicators := make([]bool, 0)

	buttons := make([]map[int]bool, 0)
	button := make(map[int]bool)

	joltageRequirements := make(map[int]bool)

	for i, r := range s {
		switch r {
		case '[':
			if parserState == null {
				parserState = bracketsOpened
				continue
			}
			return nil, fmt.Errorf("invalid character '[' at position %d", i)
		case ']':
			if parserState == bracketsOpened {
				parserState = bracketsClosed
				continue
			}
			return nil, fmt.Errorf("invalid character ']' at position %d", i)
		case '(':
			if parserState == bracketsClosed || parserState == parenthesesClosed {
				parserState = parenthesesOpened
				continue
			}
			return nil, fmt.Errorf("invalid character '(' at position %d", i)
		case ')':
			if parserState == parenthesesOpened {
				parserState = parenthesesClosed
				res, err := builderToInt(builder)
				if err != nil {
					return nil, err
				}
				builder.Reset()
				button[res] = true
				buttons = append(buttons, button)
				button = make(map[int]bool)
				continue
			}
			return nil, fmt.Errorf("invalid character ')' at position %d", i)
		case '{':
			if parserState == parenthesesClosed {
				parserState = curlyBracesOpened
				continue
			}
			return nil, fmt.Errorf("invalid character '{' at position %d", i)
		case '}':
			if parserState == curlyBracesOpened {
				parserState = curlyBracesClosed
				res, err := builderToInt(builder)
				if err != nil {
					return nil, err
				}
				builder.Reset()
				joltageRequirements[res] = true
				continue
			}
			return nil, fmt.Errorf("invalid character '{' at position %d", i)
		case ' ':
			if parserState == bracketsClosed || parserState == parenthesesClosed {
				continue
			}
			return nil, fmt.Errorf("invalid character ' ' at position %d", i)
		case ',':
			switch parserState {
			case parenthesesOpened:
				res, err := builderToInt(builder)
				if err != nil {
					return nil, err
				}
				builder.Reset()
				button[res] = true
				continue
			case curlyBracesOpened:
				res, err := builderToInt(builder)
				if err != nil {
					return nil, err
				}
				builder.Reset()
				joltageRequirements[res] = true
				continue
			default:
				return nil, fmt.Errorf("invalid character ',' at position %d", i)
			}
		case '.':
			if parserState == bracketsOpened {
				indicators = append(indicators, false)
				continue
			}
			return nil, fmt.Errorf("invalid character '.' at position %d", i)
		case '#':
			if parserState == bracketsOpened {
				indicators = append(indicators, true)
				continue
			}
			return nil, fmt.Errorf("invalid character '#' at position %d", i)
		default:
			if parserState == parenthesesOpened || parserState == curlyBracesOpened {
				_, err := builder.WriteRune(r)
				if err != nil {
					return nil, err
				}
				continue
			}
			return nil, fmt.Errorf("invalid character at position %d", i)
		}
	}

	return &Machine{
		indicators:          indicators,
		buttons:             buttons,
		joltageRequirements: joltageRequirements,
	}, nil
}

func builderToInt(b strings.Builder) (int, error) {
	s := b.String()
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
