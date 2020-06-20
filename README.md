
# Turing Machine Runner
A program that runs Turing Machines as defined in chapter 12 of the [Open Logic Textbook](https://slc.openlogicproject.org/slc-screen.pdf).

## Example Machines
Several example machines are provided. Including the even machine, the doubler machine, and the adder machine as mentioned in the textbook.

## Creating Your Own Machine
To create your own machine, create a text file with the fields `tape`, `transition` and `start`.

You can tell the script which file to read from by changing this line

```go
file, err  := os.Open("doubler.in") // <-- change this file to run a different machine
```

### Defining an Initial State
You can define the initial state of your machine as ![](https://latex.codecogs.com/svg.latex?q_0) by including the line.
```
start q0
```

### Defining the Tape
To run your machine, you need a tape. For example, If you want `1000101` to be the initial contents of your tape, you can do so by including in your text file the line
```
tape 1000101
```

### Defining Transitions
As described in the Open Logic Textbook, different transitions exist for the machine based on the state of the machine and the reading of the tape head.

An example is ![](https://latex.codecogs.com/svg.latex?%5Cdelta%28q_0%2C%201%29%3D%3Cq_1%2C1%2CR%3E) which describes the behavior of the machine in state ![](https://latex.codecogs.com/svg.latex?q_0) and when the tape head reads 1.

We describe this function in the input file by listing the terms in the inputs and outputs of the function from left to right.

```
transition q0 1 q1 1 R
```

> **Note:** You can only describe the direction of the tape head in terms of R, L, and N in the input file.

As it turns out, this function describes the arrow that points from ![](https://latex.codecogs.com/svg.latex?q_0) to ![](https://latex.codecogs.com/svg.latex?q_1) in the following diagram, which illustrates the even machine.

![](https://mermaid.ink/svg/eyJjb2RlIjoiZ3JhcGggTFJcblxucTAoKHEwKSkgLS0gMSwgMSwgUiAtLT4gcTEoKHExKSlcblxucTEoKHExKSkgLS0gMSwgMSwgUiAtLT4gcTAoKHEwKSlcblxucTEoKHExKSkgLS0gMCwgMCwgUiAtLT4gcTEoKHExKSkiLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)
