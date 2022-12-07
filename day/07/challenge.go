package day_07

import (
	"fmt"
	"strings"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var (
	nonBloatedThreshold     = 100000
	totalSystemSize         = 70000000
	minSpaceNeededForUpdate = 30000000
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
	child := NewFSNode()
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

func NewFSNode() *FSNode {
	return &FSNode{
		children: make(map[string]*FSNode),
		size:     0,
	}
}

// Find all of the directories with a total size of at most 100000.
// What is the sum of the total sizes of those directories?
func Part1(input string) int {
	var (
		root  = NewFSNode()
		lines = utils.StrToSlice(input, "\n")
	)
	buildFS(lines, root)
	totalBloat := calculateTotalBloat(root)
	return totalBloat
}

func Part2(input string) int {
	var (
		root  = NewFSNode()
		lines = utils.StrToSlice(input, "\n")
	)
	buildFS(lines, root)
	sizeOfDirToDelete := findSmallestDirAboveThreshold(root, root.GetSize(), root.GetSize())
	return sizeOfDirToDelete
}

func buildFS(lines []string, root *FSNode) {
	currentNode := root
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
}

func calculateTotalBloat(n *FSNode) int {
	var totalBloat int
	for _, x := range n.children {
		size := x.GetSize()
		if x.IsDir() && size < nonBloatedThreshold {
			totalBloat += size
		}
		totalBloat += calculateTotalBloat(x)
	}
	return totalBloat
}

func findSmallestDirAboveThreshold(n *FSNode, s int, rootSize int) int {
	for _, x := range n.children {
		if !x.IsDir() {
			continue
		}
		size := x.GetSize()
		isBigEnough := (totalSystemSize-rootSize)+size >= minSpaceNeededForUpdate
		if size < s && isBigEnough {
			s = size
			s = findSmallestDirAboveThreshold(x, s, rootSize)
		}
	}
	return s
}
