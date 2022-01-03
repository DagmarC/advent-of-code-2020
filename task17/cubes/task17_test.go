package cubes

import (
	"fmt"
	"testing"
)

var inputTest []string = []string{".#.", "..#", "###"}

func TestActiveCubesCount(t *testing.T) {
	cubes := &Cubes{}
	cubes.Initialize(inputTest)

	cycles := 6
	for cycles > 0 {
		fmt.Println("CYCLE: ------------------", cycles)
		cubes = BootProcess(*cubes)
		cycles--
	}
	expected := 112

	got := cubes.Active()
	if expected != got {
		t.Fatalf("Expected = %d, Got = %d", expected, got)
	}
}
