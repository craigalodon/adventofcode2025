package main

import (
	"bufio"
	"fmt"
	"os"
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
	parser := NewParser()

	for scanner.Scan() {
		line := scanner.Text()
		_, err := parser.Deserialize(line)
		if err != nil {
			return fmt.Errorf("error deserializing: %v", err)
		}
	}

	root, err := parser.GetRoot()
	if err != nil {
		return fmt.Errorf("error getting root: %w", err)
	}

	seen := make(map[string]bool)

	var recursiveAdd func(n *Node)
	recursiveAdd = func(n *Node) {
		if seen[n.Name] {
			return
		}
		seen[n.Name] = true
		for _, child := range n.children {
			recursiveAdd(child)
		}
	}

	recursiveAdd(root)

	println("Seen: ", len(seen))

	return nil
}

type Node struct {
	Name     string
	children map[string]*Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		children: make(map[string]*Node),
	}
}

func (n *Node) AddChild(child *Node) error {
	if _, ok := n.children[child.Name]; ok {
		return fmt.Errorf("child %s already exists", child.Name)
	}
	n.children[child.Name] = child
	return nil
}

func (n *Node) ForEachChild(fn func(n *Node)) {
	for _, child := range n.children {
		fn(child)
	}
}

type Parser struct {
	state   ParserState
	builder strings.Builder
	curr    *Node
	seen    map[string]*Node
}

func NewParser() *Parser {
	return &Parser{
		state:   readingName,
		builder: strings.Builder{},
		seen:    make(map[string]*Node),
	}
}

func (p *Parser) GetNode(name string) (*Node, error) {
	if node, ok := p.seen[name]; ok {
		return node, nil
	}

	return nil, fmt.Errorf("curr not found: %s", name)
}

func (p *Parser) GetRoot() (*Node, error) {
	root, err := p.GetNode("you")
	if err != nil {
		return nil, fmt.Errorf("root not found")
	}

	return root, nil
}

func (p *Parser) Deserialize(s string) (*Node, error) {
	p.clearCurrent()
	p.state = readingName

	for _, r := range s {
		err := p.parse(r)
		if err != nil {
			return nil, err
		}
	}

	if p.state == readingChild {
		err := p.createChild()
		if err != nil {
			return nil, err
		}

		curr, err := p.getCurrent()
		if err != nil {
			return nil, err
		}

		return curr, nil
	}

	return nil, fmt.Errorf("invalid state")
}

func (p *Parser) parse(r rune) error {
	switch r {
	case ':':
		err := p.setCurrent()
		if err != nil {
			return err
		}
	case ' ':
		switch p.state {
		case readingChild:
			err := p.createChild()
			if err != nil {
				return err
			}
		case seekingChild:
		default:
			return fmt.Errorf("invalid character '%c'", r)
		}
	default:
		p.writeRune(r)
		if p.state == seekingChild {
			p.state = readingChild
		}
	}

	return nil
}

func (p *Parser) writeRune(r rune) {
	p.builder.WriteRune(r)
}

func (p *Parser) getString() string {
	return p.builder.String()
}

func (p *Parser) resetString() {
	p.builder.Reset()
}

func (p *Parser) createChild() error {
	if p.state != readingChild {
		return fmt.Errorf("invalid state")
	}

	parent, err := p.getCurrent()
	if err != nil {
		return err
	}

	name := p.getString()

	if parent.Name == name {
		return fmt.Errorf("curr cannot be own parent")
	}

	var child *Node
	if _, ok := p.seen[name]; !ok {
		child = NewNode(name)
		p.seen[name] = child
	} else {
		child = p.seen[name]
	}

	err = parent.AddChild(child)
	if err != nil {
		return err
	}

	p.state = seekingChild
	p.resetString()

	return nil
}

func (p *Parser) setCurrent() error {
	if p.state != readingName || p.curr != nil {
		return fmt.Errorf("invalid state")
	}

	name := p.getString()

	if ptr, ok := p.seen[name]; ok {
		p.curr = ptr
	} else {
		p.curr = NewNode(name)
		p.seen[name] = p.curr
	}

	p.state = seekingChild
	p.resetString()

	return nil
}

func (p *Parser) getCurrent() (*Node, error) {
	if p.curr == nil {
		return nil, fmt.Errorf("current not found")
	}
	return p.curr, nil
}

func (p *Parser) clearCurrent() {
	p.curr = nil
}

type ParserState int

const (
	readingName ParserState = iota
	seekingChild
	readingChild
)
