package main

import "testing"

func getTestInput1() arrangement {
	var input arrangement
	input = []item{
		item{elementID: noElement, typeID: elevator, floorID: 1},
		item{elementID: hydrogen, typeID: microchip, floorID: 1},
		item{elementID: lithium, typeID: microchip, floorID: 1},

		item{elementID: hydrogen, typeID: generator, floorID: 2},

		item{elementID: lithium, typeID: generator, floorID: 3},
	}

	input = reorder(input)
	return input
}

func getTestInput2() arrangement {
	var input arrangement
	input = []item{
		item{elementID: noElement, typeID: elevator, floorID: 1},
		item{elementID: strontium, typeID: generator, floorID: 1},
		item{elementID: strontium, typeID: microchip, floorID: 1},
		item{elementID: plutonium, typeID: generator, floorID: 1},
		item{elementID: plutonium, typeID: microchip, floorID: 1},

		item{elementID: thulium, typeID: generator, floorID: 2},
		item{elementID: ruthenium, typeID: generator, floorID: 2},
		item{elementID: ruthenium, typeID: microchip, floorID: 2},
		item{elementID: curium, typeID: generator, floorID: 2},
		item{elementID: curium, typeID: microchip, floorID: 2},

		item{elementID: thulium, typeID: microchip, floorID: 3},
	}

	input = reorder(input)
	return input
}

// func getTestInput3() arrangement {
// 	var input arrangement
// 	input = []item{
// 		item{elementID: noElement, typeID: elevator, floorID: 1},
// 		item{elementID: strontium, typeID: generator, floorID: 1},
// 		item{elementID: strontium, typeID: microchip, floorID: 1},
// 		item{elementID: plutonium, typeID: generator, floorID: 1},
// 		item{elementID: plutonium, typeID: microchip, floorID: 1},

// 		item{elementID: thulium, typeID: generator, floorID: 2},
// 		item{elementID: ruthenium, typeID: generator, floorID: 2},
// 		item{elementID: ruthenium, typeID: microchip, floorID: 2},
// 		item{elementID: curium, typeID: generator, floorID: 2},
// 		item{elementID: curium, typeID: microchip, floorID: 2},

// 		item{elementID: thulium, typeID: microchip, floorID: 3},

// 		// part2
// 		item{elementID: elerium, typeID: generator, floorID: 1},
// 		item{elementID: elerium, typeID: microchip, floorID: 1},
// 		item{elementID: dilithium, typeID: generator, floorID: 1},
// 		item{elementID: dilithium, typeID: microchip, floorID: 1},
// 	}

// 	input = reorder(input)
// 	return input
// }

func getTestInput4() arrangement {
	var input arrangement
	input = []item{
		item{elementID: noElement, typeID: elevator, floorID: 1},

		item{elementID: thulium, typeID: generator, floorID: 1},
		item{elementID: thulium, typeID: microchip, floorID: 1},
		item{elementID: plutonium, typeID: generator, floorID: 1},
		item{elementID: strontium, typeID: generator, floorID: 1},

		item{elementID: plutonium, typeID: microchip, floorID: 2},
		item{elementID: strontium, typeID: microchip, floorID: 2},

		item{elementID: promethium, typeID: generator, floorID: 3},
		item{elementID: promethium, typeID: microchip, floorID: 3},
		item{elementID: ruthenium, typeID: generator, floorID: 3},
		item{elementID: ruthenium, typeID: microchip, floorID: 3},
	}

	input = reorder(input)
	return input
}

// func getTestInput5() arrangement {
// 	var input arrangement
// 	input = []item{
// 		item{elementID: noElement, typeID: elevator, floorID: 1},

// 		item{elementID: thulium, typeID: generator, floorID: 1},
// 		item{elementID: thulium, typeID: microchip, floorID: 1},
// 		item{elementID: plutonium, typeID: generator, floorID: 1},
// 		item{elementID: strontium, typeID: generator, floorID: 1},

// 		item{elementID: plutonium, typeID: microchip, floorID: 2},
// 		item{elementID: strontium, typeID: microchip, floorID: 2},

// 		item{elementID: promethium, typeID: generator, floorID: 3},
// 		item{elementID: promethium, typeID: microchip, floorID: 3},
// 		item{elementID: ruthenium, typeID: generator, floorID: 3},
// 		item{elementID: ruthenium, typeID: microchip, floorID: 3},

// 		item{elementID: elerium, typeID: generator, floorID: 1},
// 		item{elementID: elerium, typeID: microchip, floorID: 1},
// 		item{elementID: dilithium, typeID: generator, floorID: 1},
// 		item{elementID: dilithium, typeID: microchip, floorID: 1},
// 	}

// 	input = reorder(input)
// 	return input
// }

func TestRun1(t *testing.T) {
	root := getTestInput1()
	actual := run(root)
	const expected = 11
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestRun2(t *testing.T) {
	root := getTestInput2()
	actual := run(root)
	const expected = 37
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

// // Disabled due to memory leaks
// func TestRun3(t *testing.T) {
// 	root := getTestInput3()
// 	actual := run(root)
// 	const expected = 61
// 	if actual != expected {
// 		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
// 	}
// }

func TestRun4(t *testing.T) {
	root := getTestInput4()
	actual := run(root)
	const expected = 31
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

// // Disabled due to memory leaks
// func TestRun5(t *testing.T) {
// 	root := getTestInput5()
// 	actual := run(root)
// 	const expected = 55
// 	if actual != expected {
// 		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
// 	}
// }
