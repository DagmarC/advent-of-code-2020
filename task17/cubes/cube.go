package cubes

import (
	"fmt"
	"strings"
)

func BootProcess(c Cubes) *Cubes {

	fmt.Println("START Boot process on cube x, y, z:", c.width, c.height, c.z)
	c.PrintPlanes()

	// 1. Expand dimensions - append corner neighbours (x, y dimensions) and planes (z dimension).
	c.AppendNeighbours()

	// 2. Create new cubes with empty planes.
	newCubes := CreateEmptyCubes()
	newCubes.height = c.height
	newCubes.width = c.width
	newCubes.z = c.z

	// 3. Loop over all cubes and change state accordingly, then attach it to a new cube.
	for z, plane := range c.planes {

		// Construct new plane.
		newPlane := make([]string, c.height)

		for y, line := range *plane {

			// This lCopy will be modified during the Boot process and appended to a newPlane.
			lCopy := []rune(line)
			// 4. Loop over each line element == cube.
			for x, cube := range line {

				// 5. Count Active neighbours for each cube.
				activeN := countActiveNeighbours(z, x, y, &c)

				// 6. Check if the cube state should be changed, if yes then modify lCopy at x position.
				if newState, ok := changeState(cube, activeN); ok {
					lCopy[x] = newState
				}
			}
			newPlane[y] = string(lCopy)
		}
		newCubes.planes[z] = &newPlane
	}
	return newCubes
}

func changeState(cube rune, count int) (rune, bool) {
	// If a cube is active and exactly 2 or 3 of its neighbors are also active,
	// the cube remains active. Otherwise, the cube becomes inactive.
	if cube == '#' && !(count == 2 || count == 3) {
		return '.', true
	}
	// If a cube is inactive but exactly 3 of its neighbors are active,
	// the cube becomes active. Otherwise, the cube remains inactive.
	if cube == '.' && count == 3 {
		return '#', true
	}
	// No change.
	return cube, false
}

func countActiveNeighbours(z, x, y int, cube *Cubes) int {
	count := 0
	for zz := z - 1; zz <= z+1; zz++ {

		for xx := x - 1; xx <= x+1; xx++ {

			for yy := y - 1; yy <= y+1; yy++ {

				if cube.checkRange(xx, yy, zz) && !(xx == x && yy == y && zz == z) {
					next := (*cube.planes[zz])[yy][xx]

					if rune(next) == '#' {
						count++
					}
				}
			}

		}
	}
	return count
}

type Cubes struct {
	planes map[int]*[]string
	z      int
	width  int
	height int
}

func CreateEmptyCubes() *Cubes {
	c := &Cubes{}
	c.planes = make(map[int]*[]string, 0)
	return c
}

func (c *Cubes) Initialize(input []string) {
	c.planes = make(map[int]*[]string, 1)
	c.planes[0] = &input
	c.z = 0
	c.width = len(input[0])
	c.height = len(input)
}

func (c *Cubes) AppendNeighbours() {
	c.appendBoundaries()
	c.appendPlanes()
}

func (c *Cubes) Active() int {
	sum := 0
	for _, plane := range c.planes {
		sum += activePerPlane(plane)
	}
	return sum
}

func (c *Cubes) appendBoundaries() {
	c.width += 2
	c.height += 2

	for _, plane := range c.planes {
		// For line in plane append . at the beginning and the end -> width.
		for i, line := range *plane {
			(*plane)[i] = "." + line + "."
		}
		// Append non-active neighbours up and down the plane -> height.
		*plane = append([]string{strings.Repeat(".", c.width)}, *plane...)
		*plane = append(*plane, strings.Repeat(".", c.width))
	}
}

func (c *Cubes) appendPlanes() {
	upperBoundary := nonActivePlane(c.width, c.height)
	lowerBoundary := nonActivePlane(c.width, c.height)
	c.z++
	c.AddPlane(c.z, &upperBoundary)
	c.AddPlane(-c.z, &lowerBoundary)
}

func (c *Cubes) AddPlane(z int, plane *[]string) {
	c.planes[z] = plane
}

func (c *Cubes) PrintPlanes() {
	for z, _ := range c.planes {
		fmt.Println("Plane at z = ", z)
		c.PrintPlane(z)
	}
}

func (c *Cubes) PrintPlane(z int) {
	for _, line := range *c.planes[z] {
		fmt.Println(line)
	}
}

func (c *Cubes) checkRange(x, y, z int) bool {
	_, ok := c.planes[z]
	if !ok {
		return false
	}
	if y >= c.height || y < 0 {
		return false
	}
	if x >= c.width || x < 0 {
		return false
	}
	return true
}

func nonActivePlane(w, h int) []string {
	var plane []string
	for i := 0; i < h; i++ {
		plane = append(plane, strings.Repeat(".", w))
	}
	return plane
}

func activePerPlane(plane *[]string) int {
	c := 0
	for _, x := range *plane {
		c += strings.Count(x, "#")
	}
	return c
}
