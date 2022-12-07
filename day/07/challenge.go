package day_07

import (
	"fmt"
	"strings"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

type FSNode struct {
	parent   *FSNode
	children map[string]*FSNode
	size     int
}

func (n *FSNode) GetSize() int {
	var size int
	if len(n.children) > 0 {
		// it's a dir
		for _, v := range n.children {
			size += v.GetSize()
		}
	} else {
		// it's a file
		size = n.size
	}
	return size
}

func (n *FSNode) Insert(size int, name string) {
	child := NewFsNode()
	child.size = size
	child.parent = n
	n.children[name] = child
}

func (n *FSNode) Print(offset int) {
	padding := "-"
	for i := 0; i < offset; i++ {
		padding += "-"
	}
	for k, v := range n.children {
		size := v.size

		if v.IsDir() {
			fmt.Printf("%s %s (dir - %d)\n", padding, k, v.GetSize())
			v.Print(offset + 1)
		} else {
			fmt.Printf("%s %s - %d\n", padding, k, size)
		}
	}
}

func (n *FSNode) IsDir() bool {
	return len(n.children) > 0
}

func (n *FSNode) CalculateTotalBloat() int {
	var totalBloat int
	for _, x := range n.children {
		size := x.GetSize()
		if x.IsDir() && size < 100000 {
			totalBloat += size
		}
		totalBloat += x.CalculateTotalBloat()
	}
	return totalBloat
}

func NewFsNode() *FSNode {
	return &FSNode{
		children: make(map[string]*FSNode),
		size:     0,
	}
}

// Find all of the directories with a total size of at most 100000.
// What is the sum of the total sizes of those directories?
func Part1(input string) int {
	var (
		root        = NewFsNode()
		currentNode = root
		lines       = utils.StrToSlice(input, "\n")
	)
	for _, line := range lines {
		parsed := strings.Split(line, " ")

		if string(line[0]) == "$" {
			// it's a command
			cmd := parsed[1]
			if cmd == "cd" {
				if parsed[2] == ".." {
					currentNode = currentNode.parent
					continue
				}
				if parsed[2] == "/" {
					currentNode = root
					continue
				}
				currentNode = currentNode.children[parsed[2]]
				continue
			}
			if parsed[1] == "ls" {
				// ignore ls?! just continue I guess...
				continue
			}
		} else {
			// it's a fs node
			name := parsed[1]
			var size int
			// fmt.Println(currentNode)
			if parsed[0] != "dir" {
				size = utils.StrToInt(parsed[0], 0)
			}
			currentNode.Insert(size, name)
		}
	}
	root.Print(0)
	totalBloat := root.CalculateTotalBloat()
	return totalBloat
}

func Part2(input string) {}
