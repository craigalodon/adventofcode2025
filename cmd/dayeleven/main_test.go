package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func (n *Node) collectChildNames() []string {
	var names []string
	n.ForEachChild(func(child *Node) {
		names = append(names, child.Name)
	})
	return names
}

func TestParser_Deserialize(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
		children []string
		error    error
	}{
		{
			name:     "Only-Child",
			s:        "aaa: bbb",
			expected: "aaa",
			children: []string{"bbb"},
			error:    nil,
		},
		{
			name:  "Childless",
			s:     "aaa: ",
			error: errors.New("invalid state"),
		},
		{
			name:     "Average Family",
			s:        "aaa: bbb ccc",
			expected: "aaa",
			children: []string{"bbb", "ccc"},
			error:    nil,
		},
		{
			name:  "Parent Of Self",
			s:     "aaa: aaa",
			error: errors.New("curr cannot be own parent"),
		},
		{
			name:  "Duplicate Children",
			s:     "aaa: bbb bbb",
			error: errors.New("child bbb already exists"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			result, err := p.Deserialize(tt.s)
			if tt.error != nil {
				if err.Error() != tt.error.Error() {
					t.Errorf("Deserialize() error = %v, wantErr %v", err, tt.error)
				}
			} else {
				if err != nil {
					t.Errorf("Deserialize() error = %v, wantErr %v", err, tt.error)
				} else {
					if tt.expected != result.Name {
						t.Errorf("Deserialize() = %v, want %v", result.Name, tt.expected)
					}
					names := result.collectChildNames()
					if !reflect.DeepEqual(names, tt.children) {
						t.Errorf("Deserialize() = %v, want %v", names, tt.children)
					}
				}
			}
		})
	}
}

func TestParser_Deserialize_Cycle(t *testing.T) {
	data := []string{
		"you: aaa",
		"aaa: bbb",
		"bbb: ccc",
		"ccc: you",
	}

	p := NewParser()

	for _, d := range data {
		_, err := p.Deserialize(d)
		if err != nil {
			t.Errorf("Deserialize() error = %v, wantErr %v", err, nil)
		}
	}

	root, err := p.GetRoot()
	if err != nil {
		t.Errorf("GetRoot() error = %v, wantErr %v", err, nil)
	}

	found, depth := root.Reachable(root, 4)

	if found != true {
		t.Errorf("Reachable() = %v, want %v", found, true)
	}

	if depth != 4 {
		t.Errorf("Reachable() = %v, want %v", depth, 4)
	}
}

func TestNode_CountPaths_ExampleInput(t *testing.T) {
	filepath := "../../testdata/dayeleven/example_part_one.txt"
	parser, err := readExampleFile(filepath)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	root, err := parser.GetRoot()
	if err != nil {
		t.Fatalf("GetRoot() error = %v, wantErr %v", err, nil)
	}

	exit, err := parser.GetExit()
	if err != nil {
		t.Fatalf("GetExit() error = %v, wantErr %v", err, nil)
	}

	count := root.CountPaths(exit, 100)
	if count != 5 {
		t.Errorf("CountPaths() = %v, want %v", count, 5)
	}
}

func TestNode_CountPathsWithStops_ExampleInput(t *testing.T) {
	filepath := "../../testdata/dayeleven/example_part_two.txt"
	parser, err := readExampleFile(filepath)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	server, err := parser.GetNode("svr")
	if err != nil {
		t.Fatalf("GetNode() error = %v, wantErr %v", err, nil)
	}

	exit, err := parser.GetExit()
	if err != nil {
		t.Fatalf("GetExit() error = %v, wantErr %v", err, nil)
	}

	stops, err := getStops(parser)
	if err != nil {
		t.Fatalf("error getting stops: %v", err)
	}

	count := server.CountPathsWithStops(exit, 100, stops)
	if count != 2 {
		t.Errorf("CountPaths() = %v, want %v", count, 2)
	}
}

func readExampleFile(filepath string) (*Parser, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", filepath)
	}
	defer func() {
		if err := file.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	parser := NewParser()
	for scanner.Scan() {
		_, err := parser.Deserialize(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error deserializing: %v", err)
		}
	}
	return parser, nil
}
