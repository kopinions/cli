package parser

import (
	"fmt"
	"github.com/cde/client/cmd"
	docopt "github.com/docopt/docopt-go"
)

func Keys(argv []string) error {
	usage := `
Valid commands for keys:

keys:list        list SSH keys for the logged in user
keys:add         add an SSH key
keys:remove      remove an SSH key

Use 'cde help [command]' to learn more.
`
	switch argv[0] {
	case "keys:list":
		return keyList(argv)
	case "keys:add":
		return addKey(argv)
	case "keys:remove":
		return removeKey(argv)
	case "keys":
		fmt.Print(usage)
	default:
		if printHelp(argv, usage) {
			return nil
		}
		PrintUsage()
		return nil
	}
	return nil
}

func keyList(argv []string) error {
	usage := `
List keys.

Usage: cde keys:list [<user>]

Arguments:
  <user>
    the logged user itself.
`
	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	user := safeGetValue(args, "<user>")
	return cmd.ListKeys(user)
}

func addKey(argv []string) error {
	usage := `
Add a key.

Usage: cde keys:add <user> <ssh>

Arguments:
  <user>
    the logged user itself.
  <ssh>
  	the ssh content
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	user := safeGetValue(args, "<user>")
	ssh := safeGetValue(args, "<ssh>")
	return cmd.AddKey(user, ssh)
}

func removeKey(argv []string) error {
	usage := `
Remove a key.

Usage: cde keys:remove <user> <key>

Arguments:
  <user>
    the logged user itself.
  <key>
  	the key id
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	user := safeGetValue(args, "<user>")
	keyId := safeGetValue(args, "<key>")
	return cmd.RemoveKey(user, keyId)
}