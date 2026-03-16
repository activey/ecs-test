package world

import "fmt"

type NavigationPath struct {
	nodes []*NavigationNode
}

func NewNavigationPath() *NavigationPath {
	return &NavigationPath{
		nodes: make([]*NavigationNode, 0),
	}
}

func (p *NavigationPath) AddNode(node *NavigationNode) {
	p.nodes = append(p.nodes, node)
}

func (p *NavigationPath) Cancel() {
	fmt.Println("canceling path")
	clear(p.nodes)

	p.nodes = make([]*NavigationNode, 0)
}

func (p *NavigationPath) HasNodes() bool {
	return p.nodes != nil && len(p.nodes) > 0
}

func (p *NavigationPath) LastNode() *NavigationNode {
	return p.nodes[len(p.nodes)-1]
}

func (p *NavigationPath) RemoveLastNode() {
	if len(p.nodes) > 0 {
		p.nodes = p.nodes[:len(p.nodes)-1]
	}
}

func (p *NavigationPath) ComputeOrientations() {
	if len(p.nodes) < 3 {
		return
	}

	for i := 1; i < len(p.nodes)-1; i++ { // Connect at 1 and end at len(p.nodes)-1 to check both previous and next
		prev := p.nodes[i-1]
		current := p.nodes[i]

		if prev.Y == current.Y {
			current.Orientation = NodeHorizontal
		} else if prev.X == current.X {
			current.Orientation = NodeVertical
		}
	}
}
