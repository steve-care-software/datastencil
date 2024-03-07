package layers

// NewLayersForTests creates a new layers for tests
func NewLayersForTests(list []Layer) Layers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(instructions Instructions, output Output, input string) Layer {
	ins, err := NewLayerBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputWithExecuteForTests creates a new output with execute for tests
func NewOutputWithExecuteForTests(variable string, kind Kind, execute string) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputForTests creates a new output for tests
func NewOutputForTests(variable string, kind Kind) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKindWithContinueForTests creates a new kind with continue for tests
func NewKindWithContinueForTests() Kind {
	ins, err := NewKindBuilder().Create().IsContinue().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKindWithPromptForTests creates a new kind with prompt for tests
func NewKindWithPromptForTests() Kind {
	ins, err := NewKindBuilder().Create().IsPrompt().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionsForTests creates new instructions for tests
func NewInstructionsForTests(list []Instruction) Instructions {
	ins, err := NewInstructionsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithAssignmentForTests creates a new instruction with assignment for tests
func NewInstructionWithAssignmentForTests(assignment Assignment) Instruction {
	ins, err := NewInstructionBuilder().Create().WithAssignment(assignment).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithConditionForTests creates a new instruction with condition for tests
func NewInstructionWithConditionForTests(condition Condition) Instruction {
	ins, err := NewInstructionBuilder().Create().WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithRaiseErrorForTests creates a new instruction with raiseError for tests
func NewInstructionWithRaiseErrorForTests(raiseError uint) Instruction {
	ins, err := NewInstructionBuilder().Create().WithRaiseError(raiseError).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithStopForTests creates a new instruction with stop for tests
func NewInstructionWithStopForTests() Instruction {
	ins, err := NewInstructionBuilder().Create().IsStop().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionForTest creates a new condition for tests
func NewConditionForTest(variable string, instructions Instructions) Condition {
	ins, err := NewConditionResourceBuilder().Create().WithVariable(variable).WithInstructions(instructions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignmentForTests creates a new assignment for tests
func NewAssignmentForTests(name string, assignable Assignable) Assignment {
	ins, err := NewAssignmentBuilder().Create().WithName(name).WithAssignable(assignable).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithBytesForTests creates a new assignable with bytes for tests
func NewAssignableWithBytesForTests(input Bytes) Assignable {
	ins, err := NewAssignableBuilder().Create().WithBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithLayerForTests creates a new execution with layer for tests
func NewExecutionWithLayerForTests(input string, layer string) Execution {
	ins, err := NewExecutionBuilder().Create().WithInput(input).WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(input string) Execution {
	ins, err := NewExecutionBuilder().Create().WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithHashBytesForTests creates a new bytes with hashBytes for tests
func NewBytesWithHashBytesForTests(input string) Bytes {
	ins, err := NewBytesBuilder().Create().WithHashBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithCompareForTests creates a new bytes with compare for tests
func NewBytesWithCompareForTests(input []string) Bytes {
	ins, err := NewBytesBuilder().Create().WithCompare(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithJoinForTests creates a new bytes with join for tests
func NewBytesWithJoinForTests(join []string) Bytes {
	ins, err := NewBytesBuilder().Create().WithJoin(join).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
