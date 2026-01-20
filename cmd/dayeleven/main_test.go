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
			error: errors.New("node cannot be own parent"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Deserialize(tt.s, map[string]*Node{})
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
