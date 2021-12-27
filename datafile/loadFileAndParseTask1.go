package datafile

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/DagmarC/advent-of-code-2020/constants"
	"github.com/DagmarC/advent-of-code-2020/task1/binarytree"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

type loadOption struct {
	bt   *binarytree.Tree
	data *[]int
}

// LoadFileIntoBinaryTree Loads given file of integers into the binary tree.
func LoadFileIntoBinaryTree(tree *binarytree.Tree) {
	option := loadOption{bt: tree}
	loadFile(option)
}

// LoadFileIntoSlice Loads given file of integers into the slice.
func LoadFileIntoSlice(data *[]int) {
	option := loadOption{data: data}
	loadFile(option)
}

func loadFile(option loadOption) {

	node := &binarytree.Node{Value: 0}
	file, err := os.Open(path.Join(utils.GetWd(), constants.Input))
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		// Decide whether to load it into BT or Slice
		if option.bt != nil {
			node = option.bt.Insert(node, number)
		} else {
			*option.data = append(*option.data, number)
		}
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}
