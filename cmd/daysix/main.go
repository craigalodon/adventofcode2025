package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Col struct {
	Offset int
	Op     rune
	Grid   []string
}

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

	rows, err := readFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("error while reading file: %w", err)
	}

	cols := initCols(rows[len(rows)-1])
	populateGrids(rows[:len(rows)-1], cols)
	leftJustifyGrid(cols[len(cols)-1].Grid)

	part1 := 0
	part2 := 0

	for _, col := range cols {
		{
			roworder, err := getRowOrderValues(col)
			if err != nil {
				return fmt.Errorf("error during row-order traversal: %w", err)
			}
			res, err := agg(col, roworder)
			if err != nil {
				return fmt.Errorf("unable to aggregate row-order values: %w", err)
			}
			part1 += res
		}

		{
			colorder, err := getColOrderValues(col)
			if err != nil {
				return fmt.Errorf("error during col-order traversal: %w", err)
			}
			res, err := agg(col, colorder)
			if err != nil {
				return fmt.Errorf("unable to aggregate col-order values: %w", err)
			}
			part2 += res
		}
	}

	fmt.Printf("Row-Order Traversal: %v\n", part1)
	fmt.Printf("Column-Order Traversal: %v\n", part2)

	return nil
}

func agg(c Col, vals []int) (int, error) {
	if c.Op != '+' && c.Op != '*' {
		return 0, fmt.Errorf("invalid op value: %v", c.Op)
	}

	acc := 0
	if c.Op == '*' {
		acc = 1
	}

	for _, v := range vals {
		if c.Op == '+' {
			acc += v
		} else {
			acc *= v
		}
	}

	return acc, nil
}

func getColOrderValues(c Col) ([]int, error) {
	vals := make([]int, 0)
	for i := range len(c.Grid[0]) {
		builder := strings.Builder{}
		for j := range len(c.Grid) {
			builder.WriteByte(c.Grid[j][i])
		}
		s := builder.String()
		val, err := trimThenAtoi(s)
		if err != nil {
			return nil, fmt.Errorf("unable to parse column: %w", err)
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func getRowOrderValues(c Col) ([]int, error) {
	vals := make([]int, 0)
	for _, s := range c.Grid {
		val, err := trimThenAtoi(s)
		if err != nil {
			return nil, fmt.Errorf("unable to parse row: %w", err)
		}
		vals = append(vals, val)
	}

	return vals, nil
}

func trimThenAtoi(s string) (int, error) {
	prev := 0
	for _, r := range s {
		if r != ' ' {
			curr, err := strconv.Atoi(string(r))
			if err != nil {
				return 0, fmt.Errorf("unable to convert to int: %w", err)
			}
			prev = prev*10 + curr
		}
	}
	return prev, nil
}

func leftJustifyGrid(grid []string) {
	m := len(grid[0])
	for i := 1; i < len(grid); i++ {
		curr := len(grid[i])
		if curr > m {
			m = curr
		}
	}

	for i := range len(grid) {
		curr := len(grid[i])
		if curr < m {
			padding := m - curr
			grid[i] = grid[i] + strings.Repeat(" ", padding)
		}
	}
}

func populateGrids(rows []string, cols []Col) {
	for _, row := range rows {
		for i, col := range cols {
			next := len(row)
			if i != len(cols)-1 {
				next = cols[i+1].Offset - 1
			}
			cols[i].Grid = append(cols[i].Grid, row[col.Offset:next])
		}
	}
}

func initCols(row string) []Col {
	cols := make([]Col, 0)

	for i, r := range row {
		if r == '*' || r == '+' {
			cols = append(cols, Col{Offset: i, Op: r})
		}
	}

	return cols
}

func readFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	rows := make([]string, 0)

	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("error closing file: %w", err)
	}

	return rows, nil
}
