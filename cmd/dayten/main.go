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

	accPresses := 0

	for scanner.Scan() {
		line := scanner.Text()
		machine, err := deserialize(line)
		if err != nil {
			return fmt.Errorf("error deserializing line: %w", err)
		}
		presses, err := machine.configure()
		if err != nil {
			return fmt.Errorf("error configuring machine: %w", err)
		}
		accPresses += presses
	}

	fmt.Printf("Configured all machines with %d presses\n", accPresses)

	return nil
}

type Machine struct {
	configuration       int
	buttons             []int
	joltageRequirements map[int]bool
}

func (m *Machine) jolt() (int, error) {
	// TODO: Implement solution
	return 0, nil
}

func (m *Machine) configure() (int, error) {

	combs := calcCombs(len(m.buttons))
	var best *int

	for i := 0; i <= combs; i++ {
		curr := 0
		presses := 0
		j := 1
		ptr := 0
		for j <= i {
			if i&j > 0 {
				curr = press(curr, m.buttons[ptr])
				presses++
			}
			j = j << 1
			ptr++
		}
		if curr == m.configuration {
			if best == nil || *best > presses {
				best = &presses
			}
		}
	}

	if best == nil {
		return 0, fmt.Errorf("no configuration found")
	}

	return *best, nil
}

func calcCombs(buttons int) int {
	combs := 1
	for range buttons {
		combs = combs << 1
	}
	combs = combs - 1
	return combs
}

type ParserState int

const (
	started ParserState = iota
	bracketsOpened
	bracketsClosed
	parenthesesOpened
	parenthesesClosed
	curlyBracesOpened
	curlyBracesClosed
)

func deserialize(s string) (*Machine, error) {

	parserState := started
	var builder strings.Builder

	configuration := 0
	ptr := 1

	buttons := make([]int, 0)
	var button int

	joltageRequirements := make(map[int]bool)

	for i, r := range s {
		switch r {
		case '[':
			if parserState == started {
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
				res, err := tryBuilderToInt(builder)
				if err != nil {
					return nil, err
				}
				builder.Reset()
				button = addIndicator(button, res)
				buttons = append(buttons, button)
				button = 0
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
				res, err := tryBuilderToInt(builder)
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
				res, err := tryBuilderToInt(builder)
				if err != nil {
					return nil, err
				}
				builder.Reset()
				button = addIndicator(button, res)
				continue
			case curlyBracesOpened:
				res, err := tryBuilderToInt(builder)
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
				ptr = ptr << 1
				continue
			}
			return nil, fmt.Errorf("invalid character '.' at position %d", i)
		case '#':
			if parserState == bracketsOpened {
				configuration = ptr | configuration
				ptr = ptr << 1
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
		configuration:       configuration,
		buttons:             buttons,
		joltageRequirements: joltageRequirements,
	}, nil
}

func tryBuilderToInt(b strings.Builder) (int, error) {
	s := b.String()
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func addIndicator(button, indicator int) int {
	x := 1
	for range indicator {
		x = x << 1
	}
	return button | x
}

func press(curr, button int) int {
	return (curr | button) ^ (curr & button)
}
