package intcode

import (
	"adventofcode/util"
	"strings"
)

const Memory = 100000

type Operation struct {
	code           int
	parameterModes []int
}

type Program struct {
	code               []int
	instructionPointer int
	relativeBase       int
	operation          Operation
	IsRunning          bool
	Input              chan int
	Output             chan int
}

func NewIntCodeProgram(numbers []int) *Program {
	return &Program{
		code:               numbers,
		instructionPointer: 0,
		relativeBase:       0,
		Input:              make(chan int, 100),
		Output:             make(chan int, 100),
	}
}

func NewIntCodeProgramFromFile(filename string) *Program {
	line := util.ReadLines(filename)[0]
	numbersAsInt := util.ConvertStrArrToIntArr(strings.Split(line, ","))
	numbesWithMemory := make([]int, Memory)
	copy(numbesWithMemory, numbersAsInt)

	return NewIntCodeProgram(numbesWithMemory)
}

func NewIntCodeProgramWithChannels(numbers []int, inChannel, outChannel chan int) *Program {
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	return &Program{
		code:               numbersCopy,
		instructionPointer: 0,
		relativeBase:       0,
		Input:              inChannel,
		Output:             outChannel,
	}
}

func (program *Program) readNextInstruction() {
	opCode := program.code[program.instructionPointer]

	op := opCode % 100

	opCode /= 100

	var modes []int

	for opCode > 0 {
		modes = append(modes, opCode%10)
		opCode /= 10
	}

	program.operation = Operation{
		code:           op,
		parameterModes: modes,
	}
}

func (program *Program) getParameterMode(position int) int {
	if len(program.operation.parameterModes) < position {
		return 0
	}
	return program.operation.parameterModes[position-1]
}

func (program *Program) getParam(position int) int {
	// mask := getMask(position)
	//
	// paramMode := getParamFromMask(program.opCode, mask)
	paramMode := program.getParameterMode(position)

	switch paramMode {
	case 0:
		return program.code[program.code[program.instructionPointer+position]]
	case 1:
		return program.code[program.instructionPointer+position]
	case 2:
		return program.code[program.relativeBase+program.code[program.instructionPointer+position]]
	default:
		panic("unknown parameter mode!")
	}
}

func (program *Program) getOutputParam(position int) int {
	// mask := getMask(position)
	//
	// paramMode := getParamFromMask(program.opCode, mask)
	paramMode := program.getParameterMode(position)

	switch paramMode {
	case 0:
		return program.code[program.instructionPointer+position]
	case 2:
		return program.relativeBase + program.code[program.instructionPointer+position]
	default:
		panic("unknown output parameter mode!")
	}

}

func (program *Program) Step() {
	// program.opCode = program.code[program.instructionPointer]
	program.readNextInstruction()

	switch program.operation.code {
	case 99:
		program.IsRunning = false
		break
	case 1:
		num1 := program.getParam(1)
		num2 := program.getParam(2)

		param3 := program.getOutputParam(3)

		program.code[param3] = num1 + num2
		program.instructionPointer += 4
	case 2:
		num1 := program.getParam(1)
		num2 := program.getParam(2)

		param3 := program.getOutputParam(3)

		program.code[param3] = num1 * num2
		program.instructionPointer += 4
	case 3:
		param1 := program.getOutputParam(1)

		// fmt.Println("Input operation")

		// select {
		// case in := <-program.Input:

		in := <-program.Input

		// fmt.Print("Enter text: ")
		// text, _ := reader.ReadString('\n')
		// in, _ := strconv.Atoi(strings.TrimSpace(text))

		program.code[param1] = in
		program.instructionPointer += 2

		// default:
		//	panic("no input given")
		// }
	case 4:
		param1 := program.getParam(1)

		// fmt.Println("Output operation")
		// fmt.Println(param1)
		program.Output <- param1

		program.instructionPointer += 2
	case 5:
		param1 := program.getParam(1)
		param2 := program.getParam(2)

		if param1 != 0 {
			program.instructionPointer = param2
		} else {
			program.instructionPointer += 3
		}
	case 6:
		param1 := program.getParam(1)
		param2 := program.getParam(2)

		if param1 == 0 {
			program.instructionPointer = param2
		} else {
			program.instructionPointer += 3
		}
	case 7:
		param1 := program.getParam(1)
		param2 := program.getParam(2)

		param3 := program.getOutputParam(3)

		if param1 < param2 {
			program.code[param3] = 1
		} else {
			program.code[param3] = 0
		}
		program.instructionPointer += 4
	case 8:
		param1 := program.getParam(1)
		param2 := program.getParam(2)

		param3 := program.getOutputParam(3)

		if param1 == param2 {
			program.code[param3] = 1
		} else {
			program.code[param3] = 0
		}
		program.instructionPointer += 4
	case 9:
		param1 := program.getParam(1)

		program.relativeBase += param1

		program.instructionPointer += 2
	default:
		panic("unknown instruction")
	}
}

func (program *Program) RunUntilIO() {
	for program.IsRunning && program.operation.code != 3 {
		program.readNextInstruction()
		program.Step()
	}
}

func (program *Program) Run() {
	program.IsRunning = true

	for program.IsRunning {
		program.readNextInstruction()
		program.Step()
	}
}

func (program *Program) UpdateMemory(location, value int) {
	program.code[location] = value
}

func (program *Program) Location(location int) int {
	return program.code[location]
}
