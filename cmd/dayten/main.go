package main

import (
	"adventofcode2025/internal/mathutils"
	"bufio"
	"fmt"
	"math"
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

	configPresses := 0
	joltPresses := 0.0

	for scanner.Scan() {
		line := scanner.Text()
		machine, err := deserialize(line)
		if err != nil {
			return fmt.Errorf("error deserializing line: %w", err)
		}
		{
			presses, err := machine.configure()
			if err != nil {
				return fmt.Errorf("error configuring machine: %w", err)
			}
			configPresses += presses
		}
		{
			presses, err := machine.jolt()
			if err != nil {
				return fmt.Errorf("error jolting machine: %w", err)
			}
			joltPresses += presses
		}
	}

	fmt.Printf("Configured all machines with %d presses\n", configPresses)
	fmt.Printf("Jolted all machines with %v presses\n", joltPresses)

	return nil
}

type Machine struct {
	configuration int
	buttons       []int
	joltage       []int
}

func (m *Machine) jolt() (float64, error) {
	jmap := make(map[int]int)

	matrix := make([][]float64, len(m.joltage))
	ptr := 1
	for i := range matrix {
		matrix[i] = make([]float64, len(m.buttons)+1)
		matrix[i][len(m.buttons)] = float64(m.joltage[i])
		for j, button := range m.buttons {
			if button&ptr > 0 {
				matrix[i][j] = 1
				if _, ok := jmap[button]; !ok {
					jmap[button] = m.joltage[i]
				} else {
					if jmap[button] > m.joltage[i] {
						jmap[button] = m.joltage[i]
					}
				}
			} else {
				matrix[i][j] = 0
			}
		}
		ptr = ptr << 1
	}

	rref, err := mathutils.MatrixReduce(matrix)
	if err != nil {
		return 0, fmt.Errorf("error reducing matrix: %w", err)
	}

	freeVariables, params := mathutils.Parametrize(rref)

	ranges := make([]*mathutils.Range, len(params[0]))
	ranges[0] = mathutils.NewRange(1, 2)

	for i, j := range freeVariables {
		if _, ok := jmap[m.buttons[i]]; !ok {
			ranges[j+1] = mathutils.NewRange(0, 1)
		} else {
			ranges[j+1] = mathutils.NewRange(0, jmap[m.buttons[i]]+1)
		}
	}

	combs := mathutils.GenerateCombinations(ranges)

	var best *float64
	for _, comb := range combs {
		coefs := make(map[int]float64)
		valid := true
		for i, exp := range params {
			acc := 0.0
			for i, val := range comb {
				acc += exp[i] * float64(val)
			}

			rounded := math.Round(acc)
			if !mathutils.IsZero(acc - rounded) {
				valid = false
				break
			}
			if rounded < 0 {
				valid = false
				break
			}

			coefs[i] = rounded
		}

		if valid {
			for _, row := range matrix {
				acc := 0.0
				for j := 0; j < len(row)-1; j++ {
					acc += row[j] * coefs[j]
				}
				if !mathutils.IsZero(acc - row[len(row)-1]) {
					valid = false
					break
				}
			}
		}

		if valid {
			acc := 0.0
			for _, val := range coefs {
				acc += val
			}
			if best == nil || acc < *best {
				best = &acc
			}
		}
	}

	if best == nil {
		return 0, fmt.Errorf("no combinations found")
	}

	return *best, nil
}

func (m *Machine) configure() (int, error) {

	combs := calcCombs(len(m.buttons))
	var best *int

	for i := 0; i <= combs; i++ {
		val, presses := getValPresses(i, m)
		if val == m.configuration {
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

func getValPresses(comb int, m *Machine) (int, int) {
	val := 0
	presses := 0
	ptr := 1
	button := 0
	for ptr <= comb {
		if comb&ptr > 0 {
			val = press(val, m.buttons[button])
			presses++
		}
		ptr = ptr << 1
		button++
	}
	return val, presses
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

	joltage := make([]int, 0)

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
				joltage = append(joltage, res)
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
				joltage = append(joltage, res)
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
		configuration: configuration,
		buttons:       buttons,
		joltage:       joltage,
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
