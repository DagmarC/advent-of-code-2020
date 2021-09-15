package binarytree

import (
	"errors"
	"fmt"
	"reflect"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (t Tree) IsEmpty() bool {
	return reflect.DeepEqual(t, Tree{})
}

// PrintTree will print the binary tree - sorted numbers.
func (t Tree) PrintTree(root *Node) {
	if root == nil {
		return
	}
	t.PrintTree(root.Left)
	fmt.Println(root.Value)
	t.PrintTree(root.Right)
}

// Search will find the number and if it exists it returns the *Node, otherwise it returns an error.
func (t Tree) Search(number int) (*Node, error) {

	matchNode := t.Root
	for {
		if matchNode == nil {
			return nil, errors.New("Match not found.")
		}
		if matchNode.Value == number {
			break
		}
		if number > matchNode.Value {
			matchNode = matchNode.Right
		} else {
			matchNode = matchNode.Left
		}
	}
	return matchNode, nil
}

// Insert will insert into the binary tree.
// Usage: node = tree.Insert(node, number), where first inserted node is nil
func (t *Tree) Insert(node *Node, value int) *Node {

	if node == nil {
		newNode := &Node{Value: value}

		if t.Root == nil && node == nil {
			t.Root = newNode
		}
		return newNode
	}

	if value > node.Value {
		node.Right = t.Insert(node.Right, value)
	} else {
		node.Left = t.Insert(node.Left, value)
	}
	return node
}

// FindMultiplier will find two numbers where x+y=sum abd returns x*y.
// The x is the node from tree traversal and y is the node value obtained and found via sum-x.
// 1. Traverses the tree from the root and takes each node as nodeX.
// 2. Finds the nodeY value via sum - nodeX.
// 3. Searches for the nodeY value inside the tree.
// 4. If it is present then the multiplication of nodeX * nodeY wil take place, otherwise continue to the next nodeX.
func (t *Tree) FindMultiplier(nodeX *Node, sum int) int {
	if nodeX == nil {
		return 1
	}

	result := t.FindMultiplier(nodeX.Left, sum)
	if result != 1 {
		return result
	}

	findY := sum - nodeX.Value

	nodeY, _ := t.Search(findY)
	if nodeY != nil && nodeY != nodeX {
		fmt.Println("Found numbers: ", nodeX.Value, nodeY.Value)
		return nodeX.Value * nodeY.Value
	}

	result = t.FindMultiplier(nodeX.Right, sum)

	return result
}
