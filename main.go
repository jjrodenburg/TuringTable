package main

import "strings"

func main() {
	currentTapePosition := 0 // the tape can technically stretch from -inf to +inf
	tape := "1 0 1 0 1 0 1"  // Turing separates digits by empty slots.

	states := map[string]state{
		"b": {
			operations:   []operation{{P}, {R}},
			finalMConfig: "c",
		},
		"c": {
			operations:   []operation{{R}},
			finalMConfig: "k",
		},
		"k": {
			operations:   []operation{{P}},
			finalMConfig: "e",
		},
		"e": {
			operations:   []operation{{R}},
			finalMConfig: "b",
		},
	}

	curState := states["e"]

	for currentTapePosition < len(tape) {
		for _, operation := range curState.operations {
			switch operation.operationalMode {
			case P:
				curValue := string(tape[currentTapePosition])
				if strings.TrimSpace(curValue) != "" {
					print(curValue)
				}
				break

			case R:
				currentTapePosition++
				break

			case E:
				break
			}

			curState = states[curState.finalMConfig]
		}
	}

}

type state struct {
	operations   []operation
	finalMConfig string
}

type operation struct {
	operationalMode
}

type operationalMode string

const (
	R operationalMode = "R"
	P operationalMode = "P"
	E operationalMode = "E"
)
