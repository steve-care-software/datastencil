package instructions

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions"
	results_failures "github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execAccountApp      accounts.Application
	execAssignmentApp   assignments.Application
	execDatabaseApp     databases.Application
	stackBuilder        stacks.Builder
	framesBuilder       stacks.FramesBuilder
	frameBuilder        stacks.FrameBuilder
	assignmentsBuilder  stacks.AssignmentsBuilder
	interruptionBuilder interruptions.Builder
	failureBuilder      results_failures.Builder
}

func createApplication(
	execAccountApp accounts.Application,
	execAssignmentApp assignments.Application,
	execDatabaseApp databases.Application,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	interruptionBuilder interruptions.Builder,
	failureBuilder results_failures.Builder,
) Application {
	out := application{
		execAccountApp:      execAccountApp,
		execAssignmentApp:   execAssignmentApp,
		execDatabaseApp:     execDatabaseApp,
		stackBuilder:        stackBuilder,
		framesBuilder:       framesBuilder,
		frameBuilder:        frameBuilder,
		assignmentsBuilder:  assignmentsBuilder,
		interruptionBuilder: interruptionBuilder,
		failureBuilder:      failureBuilder,
	}

	return &out
}

// Execute executes the instructions
func (app *application) Execute(
	stack stacks.Stack,
	instructions instructions.Instructions,
) (stacks.Stack, interruptions.Interruption, error) {
	currentStack := stack
	list := instructions.List()
	for idx, oneInstruction := range list {
		newStack, interruption, err := app.instruction(uint(idx), stack, oneInstruction)
		if err != nil {
			return nil, nil, err
		}

		if interruption != nil {
			return newStack, interruption, nil
		}

		currentStack = newStack
	}

	return currentStack, nil, nil
}

func (app *application) instruction(
	idx uint,
	stack stacks.Stack,
	instruction instructions.Instruction,
) (stacks.Stack, interruptions.Interruption, error) {
	frame := stack.Head()
	if instruction.IsStop() {
		ins, err := app.interruptionBuilder.Create().
			WithStop(idx).
			Now()

		if err != nil {
			return nil, nil, err
		}

		return stack, ins, nil
	}

	if instruction.IsRaiseError() {
		code := instruction.RaiseError()
		failure, err := app.failureBuilder.Create().
			WithIndex(idx).
			WithCode(code).
			IsRaisedInLayer().
			Now()

		if err != nil {
			return nil, nil, err
		}

		ins, err := app.interruptionBuilder.Create().
			WithFailure(failure).
			Now()

		if err != nil {
			return nil, nil, err
		}

		return stack, ins, nil
	}

	if instruction.IsCondition() {
		condition := instruction.Condition()
		variable := condition.Variable()
		value, err := frame.FetchBool(variable)
		if err != nil {
			code := failures.CouldNotFetchConditionFromFrame
			failure, err := app.failureBuilder.Create().
				WithCode(code).
				Now()

			if err != nil {
				return nil, nil, err
			}

			ins, err := app.interruptionBuilder.Create().
				WithFailure(failure).
				Now()

			if err != nil {
				return nil, nil, err
			}

			return stack, ins, nil
		}

		if value {
			instructions := condition.Instructions()
			stack, interruption, err := app.Execute(stack, instructions)
			if err != nil {
				return nil, nil, err
			}

			if interruption != nil {
				if interruption.IsStop() {
					return stack, interruption, nil
				}

				failure := interruption.Failure()
				index := failure.Index() + idx
				code := failure.Code()
				builder := app.failureBuilder.Create().
					WithIndex(index).
					WithCode(code)

				if failure.IsRaisedInLayer() {
					builder.IsRaisedInLayer()
				}

				failure, err := builder.Now()
				if err != nil {
					return nil, nil, err
				}

				ins, err := app.interruptionBuilder.Create().
					WithFailure(failure).
					Now()

				if err != nil {
					return nil, nil, err
				}

				return stack, ins, nil
			}

			return stack, nil, nil
		}

		return stack, nil, nil
	}

	if instruction.IsAccount() {
		account := instruction.Account()
		pCode, err := app.execAccountApp.Execute(frame, account)
		if pCode != nil {
			message := ""
			if err != nil {
				message = err.Error()
			}

			ins, err := app.errorCodeToInterruption(idx, *pCode, message)
			return stack, ins, err
		}

		return stack, nil, nil
	}

	if instruction.IsDatabase() {
		database := instruction.Database()
		pCode, err := app.execDatabaseApp.Execute(frame, database)
		if pCode != nil {
			message := ""
			if err != nil {
				message = err.Error()
			}

			ins, err := app.errorCodeToInterruption(idx, *pCode, message)
			return stack, ins, err
		}

		return stack, nil, nil
	}

	assignment := instruction.Assignment()
	retAssignment, pCode, err := app.execAssignmentApp.Execute(frame, assignment)
	if pCode != nil {
		message := ""
		if err != nil {
			message = err.Error()
		}

		ins, err := app.errorCodeToInterruption(idx, *pCode, message)
		return stack, ins, err
	}

	assignmentsList := []stacks.Assignment{}
	if frame.HasAssignments() {
		assignmentsList = frame.Assignments().List()
	}

	assignmentsList = append(assignmentsList, retAssignment)
	assignments, err := app.assignmentsBuilder.Create().WithList(assignmentsList).Now()
	if err != nil {
		return nil, nil, err
	}

	newFrame, err := app.frameBuilder.Create().WithAssignments(assignments).Now()
	if err != nil {
		return nil, nil, err
	}

	framesList := []stacks.Frame{}
	if stack.HasBody() {
		framesList = stack.Body().List()
	}

	framesList = append(framesList, newFrame)
	frames, err := app.framesBuilder.Create().WithList(framesList).Now()
	if err != nil {
		return nil, nil, err
	}

	newStack, err := app.stackBuilder.Create().WithFrames(frames).Now()
	if err != nil {
		return nil, nil, err
	}

	return newStack, nil, nil
}

func (app *application) errorCodeToInterruption(idx uint, code uint, message string) (interruptions.Interruption, error) {
	builder := app.failureBuilder.Create().
		WithIndex(idx).
		WithCode(code)

	if message != "" {
		builder.WithMessage(message)
	}

	failure, err := builder.Now()
	if err != nil {
		return nil, err
	}

	return app.interruptionBuilder.Create().
		WithFailure(failure).
		Now()
}
