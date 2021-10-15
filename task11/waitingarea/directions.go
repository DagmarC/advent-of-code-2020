package waitingarea

type direction [2]int

var up direction = [2]int{-1, 0}
var down direction = [2]int{1, 0}
var right direction = [2]int{0, 1}
var left direction = [2]int{0, -1}

var dUpLeft direction = [2]int{-1, -1}
var dUpRight direction = [2]int{-1, 1}
var dDownLeft direction = [2]int{1, -1}
var dDownRight direction = [2]int{1, 1}

type Directions struct {
	allDirections []direction
}

func (d *Directions) Initialize() {
	d.allDirections = append(d.allDirections, up, down, right, left, dUpLeft, dUpRight, dDownLeft, dDownRight)
}
