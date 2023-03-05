package core

import (
	"fmt"
	"strings"
)

func Process(addresses []string, names []string) {
	var matrix = [][]NodeInfo{}
	assignments := []Node{}
	lastSizeAssignments := 0

	hashsetCols := make(map[int]int)
	hashsetRows := make(map[int]int)

	// prepare matrix
	for i := 0; i < len(addresses); i++ {
		address := addresses[i]
		streetInfo := ExtractStreetInfo(address)

		cols := []NodeInfo{}
		for j := 0; j < len(names); j++ {
			name := names[j]
			ss := float64(0)
			if streetInfo.IsEven {
				ss = float64(CountVowels(name)) * streetInfo.Multiplier
			} else {
				ss = float64(CountConsonants(name)) * streetInfo.Multiplier
			}

			lenStreet := len(streetInfo.Street)
			lenName := len(streetInfo.Street)

			if lenStreet < lenName {
				for j := 2; j < lenStreet; j++ {
					if lenStreet%j == 0 && lenName%j == 0 {
						ss = ss * 1.5
						break
					}
				}
			} else {
				for j := 2; j < lenName; j++ {
					if lenStreet%j == 0 && lenName%j == 0 {
						ss = ss * 1.5
						break
					}
				}
			}

			cols = append(cols, NodeInfo{
				street:        streetInfo.Street,
				name:          name,
				originalValue: ss,
				value:         ss,
			})
		}
		matrix = append(matrix, cols)
	}

	// zeroes
	for w := 0; w < 100; w++ { // to be endless loop
		zeroesCols := make(map[int]int)
		roundZeros := []Node{}
		skipRow := false
		for h := 0; h < len(addresses); h++ {

			for _, a := range assignments {
				if a.X == h {
					skipRow = true
					break
				}
			}

			if skipRow {
				skipRow = false
				continue
			}

			maxssline := float64(0)
			for i := 0; i < len(names); i++ {
				skipCol := false
				for _, A := range assignments {
					if i == A.Y {
						skipCol = true
						break
					}
				}
				if skipCol {
					continue
				}

				ss := matrix[h][i].value

				if ss > maxssline { // calculate max ss row
					maxssline = ss
				}
			}

			for k := 0; k < len(names); k++ {
				nodeinfo := matrix[h][k]

				skipCol := false
				for _, A := range assignments {
					if k == A.Y {
						skipCol = true
						break
					}
				}
				if skipCol {
					continue
				}

				if nodeinfo.value == float64(0) {
					continue // next col
				}

				nodeinfo.value = maxssline - nodeinfo.value
				matrix[h][k] = nodeinfo
				if nodeinfo.value == float64(0) {
					roundZeros = append(roundZeros, Node{h, k})
					zeroesCols[k] = k
				}
			}

		}

		// assignments
		maxss := float64(0)
		for _, zero := range zeroesCols {
			assignment := Node{}
			isNewAssignment := false
			for i := 0; i < len(roundZeros); i++ {
				if roundZeros[i].Y == zero {
					node := matrix[roundZeros[i].X][roundZeros[i].Y]
					v := node.originalValue
					if v > maxss {
						maxss = v
						assignment.X = roundZeros[i].X
						assignment.Y = roundZeros[i].Y
						isNewAssignment = true
					}
				}
			}
			if isNewAssignment {
				assignments = append(assignments, assignment)
				hashsetRows[assignment.X] = assignment.Y
				hashsetCols[assignment.Y] = assignment.X
				isNewAssignment = false
			}
		}

		if len(assignments) == len(addresses) {
			// done
			break
		}

		if lastSizeAssignments == len((assignments)) {
			for x := 0; x < len(addresses); x++ {
				if _, ok := hashsetRows[x]; !ok {
					for y := 0; y < len(names); y++ {
						if _, ok := hashsetCols[y]; !ok {
							assignment := Node{x, y}
							assignments = append(assignments, assignment)
							hashsetRows[assignment.X] = assignment.Y
							hashsetCols[assignment.Y] = assignment.X
							break
						}
					}
				}
			}
			break
		}
		lastSizeAssignments = len((assignments))
	}

	fmt.Println("-Assignments-")
	for _, a := range assignments {
		fmt.Println(addresses[a.X], "-", names[a.Y], "SS", matrix[a.X][a.Y].originalValue)
	}
}

func CountVowels(word string) int {
	t := 0
	for _, v := range word {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			t++
		}
	}
	return t
}

func CountConsonants(word string) int {
	t := 0
	for _, v := range word {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', ' ':
			continue
		default:
			t++
		}
	}
	return t
}

func ExtractStreetInfo(address string) StreetInfo {
	// e.g. 63187 Volkman Garden Suite 447, San Diego, CA 92126
	street := strings.Split(address, ",")[0] // 63187 Volkman Garden Suite 447

	// try to remove inner number
	street = strings.Split(strings.Split(street, "Suite")[0], "Apt.")[0] // 63187 Volkman Garden

	// split in particles
	particles := strings.Split(street, " ")

	// discard the first one particle to remove external number
	street = strings.Join(particles[1:], " ")

	len := len(street)
	isEven := true
	multiplier := 1.5

	if len%2 != 0 {
		isEven = false
		multiplier = 1
	}

	return StreetInfo{street, isEven, multiplier, len}

}
