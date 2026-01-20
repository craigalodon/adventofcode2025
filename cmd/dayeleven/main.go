package main

import (
	"fmt"
	"strings"
)

func main() {
	println("Hello, Day Eleven!")
}

type Node struct {
	Name     string
	children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		children: []*Node{},
	}
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) ForEachChild(fn func(n *Node)) {
	for _, child := range n.children {
		fn(child)
	}
}

type Parser struct {
	state    ParserState
	builder  strings.Builder
	Node     *Node
	children map[string]*Node
	seen     map[string]*Node
}

func NewParser(seen map[string]*Node) *Parser {
	return &Parser{
		state:    readingName,
		builder:  strings.Builder{},
		children: make(map[string]*Node),
		seen:     seen,
	}
}

func (p *Parser) WriteRune(r rune) {
	p.builder.WriteRune(r)
}

func (p *Parser) GetString() string {
	return p.builder.String()
}

func (p *Parser) ResetString() {
	p.builder.Reset()
}

func (p *Parser) CreateChild() error {
	if p.state != readingChild {
		return fmt.Errorf("invalid state")
	}

	name := p.GetString()

	if p.Node.Name == name {
		return fmt.Errorf("node cannot be own parent")
	}
	if _, ok := p.children[name]; ok {
		return fmt.Errorf("child %s already exists", name)
	}

	if child, ok := p.seen[name]; ok {
		p.children[name] = child
	} else {
		p.children[name] = NewNode(name)
		p.seen[name] = p.children[name]
	}

	p.Node.AddChild(p.children[name])

	p.state = seekingChild
	p.ResetString()

	return nil
}

func (p *Parser) SetName() error {
	if p.state != readingName {
		return fmt.Errorf("invalid state")
	}

	name := p.GetString()

	if ptr, ok := p.seen[name]; ok {
		p.Node = ptr
	} else {
		p.Node = NewNode(name)
		p.seen[name] = p.Node
	}

	p.state = seekingChild
	p.ResetString()

	return nil
}

func (p *Parser) GetNode() *Node {
	for _, child := range p.children {
		p.Node.AddChild(child)
	}

	return p.Node
}

type ParserState int

const (
	readingName ParserState = iota
	seekingChild
	readingChild
)

func Deserialize(s string, seen map[string]*Node) (*Node, error) {
	p := NewParser(seen)

	for _, r := range s {
		err := p.Parse(r)
		if err != nil {
			return nil, err
		}
	}

	err := p.CreateChild()
	if err != nil {
		return nil, err
	}

	return p.Node, nil
}

func (p *Parser) Parse(r rune) error {
	switch r {
	case ':':
		err := p.SetName()
		if err != nil {
			return err
		}
	case ' ':
		switch p.state {
		case readingChild:
			err := p.CreateChild()
			if err != nil {
				return err
			}
		case seekingChild:
		default:
			return fmt.Errorf("invalid character '%c'", r)
		}
	default:
		p.WriteRune(r)
		if p.state == seekingChild {
			p.state = readingChild
		}
	}

	return nil
}
