package main

func main() {
	currentTapePosition := 0 // the tape technically stretches -inf <-> +inf

	states := map[string]state{
		"b": {
			operations:   []operation{{P, "0"}, {R, ""}},
			finalMConfig: "c",
		},
		"c": {
			operations:   []operation{{R, ""}},
			finalMConfig: "k",
		},
		"k": {
			operations:   []operation{{P, "1"}},
			finalMConfig: "e",
		},
		"e": {
			operations:   []operation{{R, ""}},
			finalMConfig: "b",
		},
	}

	curState := states["b"]

	for currentTapePosition <= 10 { // bound to stop endless loop
		for _, operation := range curState.operations {
			switch operation.operationalMode {
			case P:
				print(operation.value)
				break

			case R:
				currentTapePosition++
				break
			}

			curState = states[curState.finalMConfig]
		}
	}

}

type state struct {
	symbol       string
	operations   []operation
	finalMConfig string
}

type operation struct {
	operationalMode
	value string
}

type operationalMode string

const (
	R operationalMode = "R"
	P operationalMode = "P"
)
