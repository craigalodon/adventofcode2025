package main

import (
	"errors"
	"reflect"
	"testing"
)

func (n *Node) CollectChildNames() []string {
	var names []string
	n.ForEachChild(func(child *Node) {
		names = append(names, child.Name)
	})
	return names
}

func TestDeserialize(t *testing.T) {
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
					names := result.CollectChildNames()
					if !reflect.DeepEqual(names, tt.children) {
						t.Errorf("Deserialize() = %v, want %v", names, tt.children)
					}
				}
			}
		})
	}
}

func TestDeserializeCycle(t *testing.T) {
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

	var found bool
	var recursiveSearch func(*Node, int)
	recursiveSearch = func(n *Node, depth int) {
		if found {
			return
		}

		if n == root && depth > 0 {
			found = true
			return
		}

		if depth > 10 {
			return
		}

		n.ForEachChild(func(child *Node) {
			recursiveSearch(child, depth+1)
		})
	}

	recursiveSearch(root, 0)

	if found != true {
		t.Errorf("Deserialize() = %v, want %v", found, true)
	}
}
