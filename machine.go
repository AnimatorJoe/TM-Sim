package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Tape class
type Tape struct {
	index uint32
	content []string
}

func (t *Tape) Read() string {
	return t.content[t.index]
}

func (t *Tape) Write(b string) {
	t.content[t.index] = b
}

func (t *Tape) Left() {
	if (t.index <= 0) {
		// log.Fatalf("You moved off the tape lmao")
		return
	}
	t.index--
}

func (t *Tape) Right() {
	if (t.index + 1 >= uint32(len(t.content))) {
		t.content = append(t.content, "0")
	}
	t.index++
}

// Transition struct
type Transition struct {
	state string
	write string
	direction string
}

// Transition Conditions
type TransitionCondition struct {
	state string
	tapeSymbol string
}

var tape *Tape
var transitionFunctions map[TransitionCondition]Transition
var machineState string

// main function
func main() {
	transitionFunctions = make(map[TransitionCondition]Transition)

	// open file
	file, err := os.Open("doubler.in") // <-- change this file to run a different machine
	if err != nil {
		log.Fatalf("file could not be opened: %s", err)
	}

	// parse file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if (len(line) == 0) { continue }

		switch {
		case line[0] == "tape":
			if (len(line) != 2) { log.Fatalf("You provided an incorrect number of arguments for your tape") } 
			input := []string{}
			for _, c := range line[1] {
				input = append(input, string(c))
			}
			tape = &Tape{0, input}
			// fmt.Println(tape)

		case line[0] == "transition":
			if (len(line) != 6) { log.Fatalf("You provided an incorrect number of arguments for a transition function") }
			if (!(line[5] == "R" || line[5] == "L" || line[5] == "N")) { log.Fatalf("You provided an invalid direction \"%v\"", line[5]) }
			transitionFunctions[TransitionCondition{line[1], line[2]}] = Transition{line[3], line[4], line[5]}

		case line[0] == "start":
			machineState = line[1]

		default:
			log.Fatalf("You provided an invlide prefix \"%v\"", line[0])
		}
	}

	switch {
	case machineState == "":
		log.Fatalf("You didn't define an initial state, this can be done by writing: \"start [initial_state]\" in the input file")
	case tape == nil:
		log.Fatalf("You didn't provide a tape, you can provide one by writing: \"tape [your_tape]\" in the input file")
	}

	// run the machine
	fmt.Println("initial state:", machineState)
	fmt.Println("tape:", tape.content)
	fmt.Println("transition functions:", transitionFunctions)
	fmt.Println()

	for {
		transition, hasKey := transitionFunctions[TransitionCondition{machineState, tape.Read()}];
		if (!hasKey) {
			fmt.Println("Computation complete, your tape writes\n", tape.content)
			return
		}
		tape.Write(transition.write)
		machineState = transition.state
		switch transition.direction {
		case "R":
			tape.Right()
		case "L":
			tape.Left()
		default:
			continue
		}
	}
}
