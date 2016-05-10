package parser

import (
	"github.com/docopt/docopt-go"
	"github.com/sjkyspa/stacks/client/cmd"
)

func Stacks(argv []string) error {
	usage := `
Valid commands for apps:

stacks:create        create a new stack
stacks:list          list accessible stacks
stacks:remove        remove an existing stack

Use 'cde help [command]' to learn more.
`
	switch argv[0] {
	case "stacks:create":
		return stackCreate(argv)
	case "stacks:list":
		return stackList()
	case "stacks:remove":
		return stackRemove(argv)
	default:
		if printHelp(argv, usage) {
			return nil
		}

		if argv[0] == "stacks" {
			argv[0] = "stacks:list"
			return stackList()
		}

		PrintUsage()
		return nil
	}

}

func stackCreate(argv []string) error {
	usage := `
Create a stack.

Usage: cde stacks:create <stackfile>

Arguments:
  <stackfile>
    the stack file.
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	return cmd.StackCreate(safeGetValue(args, "<stackfile>"))
}

func stackList() error {
	return cmd.StacksList()

}

func stackRemove(argv []string) error {
	usage := `
Remove an existing stack.

Usage: cde stacks:remove <stack>

Arguments:
  <stack>
    the stack name.
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	return cmd.StackRemove(safeGetValue(args, "<stack>"))
}
