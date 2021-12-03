package utils

import (
	"fmt"
)

//MakeSet initialize the set
func MakeSet() *CustomSet {
	return &CustomSet{
		container: make(map[int]struct{}),
	}
}

type CustomSet struct {
	container map[int]struct{}
}

func (c *CustomSet) Exists(key int) bool {
	_, exists := c.container[key]
	return exists
}

func (c *CustomSet) Add(key int) {
	c.container[key] = struct{}{}
}

func (c *CustomSet) Remove(key int) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *CustomSet) First() int {
	for key := range c.container {
		return key
	}
	return -1
}

func (c *CustomSet) Size() int {
	return len(c.container)
}
